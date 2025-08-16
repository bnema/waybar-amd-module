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
	
	return &Metrics{
		Power:       power,
		Temperature: temp,
		Frequency:   freq,
		Utilization: util,
	}, nil
}