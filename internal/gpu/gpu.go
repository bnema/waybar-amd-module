package gpu

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

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
	
	path := filepath.Join(gpuPath, filename)
	data, err := os.ReadFile(path)
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
	
	path := filepath.Join(devicePath, filename)
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}

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

func GetJunctionTemp() (int, error) {
	tempStr, err := readMetricFile("temp2_input")
	if err != nil {
		return 0, err
	}
	
	tempMillidegrees, err := strconv.ParseInt(tempStr, 10, 64)
	if err != nil {
		return 0, err
	}
	
	return int(tempMillidegrees / 1000), nil
}

func GetMemoryTemp() (int, error) {
	tempStr, err := readMetricFile("temp3_input")
	if err != nil {
		return 0, err
	}
	
	tempMillidegrees, err := strconv.ParseInt(tempStr, 10, 64)
	if err != nil {
		return 0, err
	}
	
	return int(tempMillidegrees / 1000), nil
}

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