// Package gpu provides functions to monitor AMD GPU metrics including power, temperature, frequency and utilization
package gpu

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// Metrics contains comprehensive GPU monitoring data
type Metrics struct {
	Power       float64 `json:"power"`
	Temperature int     `json:"temperature"`
	Frequency   float64 `json:"frequency"`
	Utilization int     `json:"utilization"`
	MemoryUsage float64 `json:"memory_usage"`
	FanSpeed    int     `json:"fan_speed"`
	Voltage     float64 `json:"voltage"`
	JunctionTemp int    `json:"junction_temp"`
	MemoryTemp   int    `json:"memory_temp"`
	PowerCap     float64 `json:"power_cap"`
}

// WaybarOutput represents the JSON structure expected by Waybar
type WaybarOutput struct {
	Text    string `json:"text"`
	Tooltip string `json:"tooltip"`
	Class   string `json:"class"`
}

var gpuPath string

func init() {
	path, err := discoverAMDGPU()
	if err != nil {
		gpuPath = ""
	} else {
		gpuPath = path
	}
}

func discoverAMDGPU() (string, error) {
	cardDirs, err := filepath.Glob("/sys/class/drm/card*")
	if err != nil {
		return "", err
	}

	for _, cardDir := range cardDirs {
		driverPath := filepath.Join(cardDir, "device", "driver")
		if target, err := os.Readlink(driverPath); err == nil {
			if strings.Contains(target, "amdgpu") {
				hwmonDirs, err := filepath.Glob(filepath.Join(cardDir, "device", "hwmon", "hwmon*"))
				if err == nil && len(hwmonDirs) > 0 {
					return hwmonDirs[0], nil
				}
			}
		}
	}

	pciDirs, err := filepath.Glob("/sys/bus/pci/drivers/amdgpu/*/hwmon/hwmon*")
	if err == nil && len(pciDirs) > 0 {
		return pciDirs[0], nil
	}

	return "", errors.New("no AMD GPU found")
}

func readMetricFile(filename string) (string, error) {
	if gpuPath == "" {
		return "", errors.New("GPU path not available")
	}
	
	path := filepath.Clean(filepath.Join(gpuPath, filename))
	// Validate that the path is within expected system directory and doesn't contain path traversal
	if !strings.HasPrefix(path, "/sys/") || strings.Contains(path, "..") {
		return "", errors.New("invalid system path")
	}
	data, err := os.ReadFile(path) // #nosec G304 - path is validated above
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

func readDeviceFile(filename string) (string, error) {
	if gpuPath == "" {
		return "", errors.New("GPU path not available")
	}
	
	devicePath := filepath.Dir(filepath.Dir(gpuPath))
	
	path := filepath.Clean(filepath.Join(devicePath, filename))
	// Validate that the path is within expected system directory and doesn't contain path traversal
	if !strings.HasPrefix(path, "/sys/") || strings.Contains(path, "..") {
		return "", errors.New("invalid system path")
	}
	data, err := os.ReadFile(path) // #nosec G304 - path is validated above
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

// GetPower returns GPU power consumption in watts
func GetPower() (float64, error) {
	powerStr, err := readMetricFile("power1_input")
	if err != nil {
		return 0, err
	}
	
	powerMicrowatts, err := strconv.ParseInt(powerStr, 10, 64)
	if err != nil {
		return 0, err
	}
	
	return float64(powerMicrowatts) / 1000000.0, nil
}

// GetTemperature returns GPU temperature in Celsius
func GetTemperature() (int, error) {
	tempStr, err := readMetricFile("temp1_input")
	if err != nil {
		return 0, err
	}
	
	tempMillidegrees, err := strconv.ParseInt(tempStr, 10, 64)
	if err != nil {
		return 0, err
	}
	
	return int(tempMillidegrees / 1000), nil
}

// GetFrequency returns GPU frequency in GHz
func GetFrequency() (float64, error) {
	freqStr, err := readMetricFile("freq1_input")
	if err != nil {
		return 0, err
	}
	
	freqHz, err := strconv.ParseInt(freqStr, 10, 64)
	if err != nil {
		return 0, err
	}
	
	return float64(freqHz) / 1000000000.0, nil
}

// GetUtilization returns GPU utilization percentage
func GetUtilization() (int, error) {
	utilStr, err := readDeviceFile("gpu_busy_percent")
	if err != nil {
		return 0, err
	}
	
	utilization, err := strconv.ParseInt(utilStr, 10, 64)
	if err != nil {
		return 0, err
	}
	
	return int(utilization), nil
}

// GetMemoryUsage returns VRAM usage percentage
func GetMemoryUsage() (float64, error) {
	usedStr, err := readDeviceFile("mem_info_vram_used")
	if err != nil {
		return 0, err
	}
	
	totalStr, err := readDeviceFile("mem_info_vram_total")
	if err != nil {
		return 0, err
	}
	
	used, err := strconv.ParseInt(usedStr, 10, 64)
	if err != nil {
		return 0, err
	}
	
	total, err := strconv.ParseInt(totalStr, 10, 64)
	if err != nil {
		return 0, err
	}
	
	if total == 0 {
		return 0, errors.New("total VRAM is 0")
	}
	
	return float64(used) / float64(total) * 100, nil
}

// GetFanSpeed returns GPU fan speed in RPM
func GetFanSpeed() (int, error) {
	fanStr, err := readMetricFile("fan1_input")
	if err != nil {
		return 0, err
	}
	
	fanRPM, err := strconv.ParseInt(fanStr, 10, 64)
	if err != nil {
		return 0, err
	}
	
	return int(fanRPM), nil
}

// GetVoltage returns GPU voltage in volts
func GetVoltage() (float64, error) {
	voltageStr, err := readMetricFile("in0_input")
	if err != nil {
		return 0, err
	}
	
	voltageMillivolts, err := strconv.ParseInt(voltageStr, 10, 64)
	if err != nil {
		return 0, err
	}
	
	return float64(voltageMillivolts) / 1000.0, nil
}

// GetJunctionTemp returns GPU junction temperature in Celsius
// Falls back to main temperature if temp2_input is not available
func GetJunctionTemp() (int, error) {
	tempStr, err := readMetricFile("temp2_input")
	if err != nil {
		// Fall back to main temperature sensor if junction temp not available
		return GetTemperature()
	}
	
	tempMillidegrees, err := strconv.ParseInt(tempStr, 10, 64)
	if err != nil {
		return 0, err
	}
	
	return int(tempMillidegrees / 1000), nil
}

// GetMemoryTemp returns GPU memory temperature in Celsius
// Falls back to main temperature if temp3_input is not available
func GetMemoryTemp() (int, error) {
	tempStr, err := readMetricFile("temp3_input")
	if err != nil {
		// Fall back to main temperature sensor if memory temp not available
		return GetTemperature()
	}
	
	tempMillidegrees, err := strconv.ParseInt(tempStr, 10, 64)
	if err != nil {
		return 0, err
	}
	
	return int(tempMillidegrees / 1000), nil
}

// GetPowerCap returns GPU power cap limit in watts
func GetPowerCap() (float64, error) {
	capStr, err := readMetricFile("power1_cap")
	if err != nil {
		return 0, err
	}
	
	capMicrowatts, err := strconv.ParseInt(capStr, 10, 64)
	if err != nil {
		return 0, err
	}
	
	return float64(capMicrowatts) / 1000000.0, nil
}

// GetAllMetrics collects all GPU metrics and returns them in a single structure
func GetAllMetrics() (*Metrics, error) {
	power, err := GetPower()
	if err != nil {
		return nil, err
	}
	
	temp, err := GetTemperature()
	if err != nil {
		return nil, err
	}
	
	freq, err := GetFrequency()
	if err != nil {
		return nil, err
	}
	
	util, err := GetUtilization()
	if err != nil {
		return nil, err
	}
	
	memUsage, err := GetMemoryUsage()
	if err != nil {
		memUsage = 0 // Don't fail on memory error, just set to 0
	}
	
	fanSpeed, err := GetFanSpeed()
	if err != nil {
		fanSpeed = 0 // Don't fail on fan speed error, just set to 0
	}
	
	voltage, err := GetVoltage()
	if err != nil {
		voltage = 0 // Don't fail on voltage error, just set to 0
	}
	
	junctionTemp, err := GetJunctionTemp()
	if err != nil {
		junctionTemp = 0 // Don't fail on junction temp error, just set to 0
	}
	
	memoryTemp, err := GetMemoryTemp()
	if err != nil {
		memoryTemp = 0 // Don't fail on memory temp error, just set to 0
	}
	
	powerCap, err := GetPowerCap()
	if err != nil {
		powerCap = 0 // Don't fail on power cap error, just set to 0
	}
	
	return &Metrics{
		Power:        power,
		Temperature:  temp,
		Frequency:    freq,
		Utilization:  util,
		MemoryUsage:  memUsage,
		FanSpeed:     fanSpeed,
		Voltage:      voltage,
		JunctionTemp: junctionTemp,
		MemoryTemp:   memoryTemp,
		PowerCap:     powerCap,
	}, nil
}