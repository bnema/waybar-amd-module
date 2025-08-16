package cpu

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type Metrics struct {
	Usage       float64 `json:"usage"`
	Temperature int     `json:"temperature"`
	Frequency   float64 `json:"frequency"`
	Cores       int     `json:"cores"`
}

type WaybarOutput struct {
	Text    string `json:"text"`
	Tooltip string `json:"tooltip"`
	Class   string `json:"class"`
}

type cpuStat struct {
	user, nice, system, idle, iowait, irq, softirq uint64
}

func parseCPUStat(line string) (cpuStat, error) {
	fields := strings.Fields(line)
	if len(fields) < 8 {
		return cpuStat{}, errors.New("invalid cpu stat line")
	}
	
	var stat cpuStat
	var err error
	
	stat.user, err = strconv.ParseUint(fields[1], 10, 64)
	if err != nil {
		return cpuStat{}, err
	}
	stat.nice, err = strconv.ParseUint(fields[2], 10, 64)
	if err != nil {
		return cpuStat{}, err
	}
	stat.system, err = strconv.ParseUint(fields[3], 10, 64)
	if err != nil {
		return cpuStat{}, err
	}
	stat.idle, err = strconv.ParseUint(fields[4], 10, 64)
	if err != nil {
		return cpuStat{}, err
	}
	stat.iowait, err = strconv.ParseUint(fields[5], 10, 64)
	if err != nil {
		return cpuStat{}, err
	}
	stat.irq, err = strconv.ParseUint(fields[6], 10, 64)
	if err != nil {
		return cpuStat{}, err
	}
	stat.softirq, err = strconv.ParseUint(fields[7], 10, 64)
	if err != nil {
		return cpuStat{}, err
	}
	
	return stat, nil
}

func (s cpuStat) total() uint64 {
	return s.user + s.nice + s.system + s.idle + s.iowait + s.irq + s.softirq
}

func (s cpuStat) active() uint64 {
	return s.user + s.nice + s.system + s.irq + s.softirq
}

func GetUsage() (float64, error) {
	// Read initial CPU stats
	data, err := os.ReadFile("/proc/stat")
	if err != nil {
		return 0, err
	}
	
	lines := strings.Split(string(data), "\n")
	if len(lines) == 0 {
		return 0, errors.New("no CPU stat data")
	}
	
	stat1, err := parseCPUStat(lines[0])
	if err != nil {
		return 0, err
	}
	
	// Wait 100ms
	time.Sleep(100 * time.Millisecond)
	
	// Read CPU stats again
	data, err = os.ReadFile("/proc/stat")
	if err != nil {
		return 0, err
	}
	
	lines = strings.Split(string(data), "\n")
	if len(lines) == 0 {
		return 0, errors.New("no CPU stat data")
	}
	
	stat2, err := parseCPUStat(lines[0])
	if err != nil {
		return 0, err
	}
	
	// Calculate usage percentage
	totalDiff := stat2.total() - stat1.total()
	activeDiff := stat2.active() - stat1.active()
	
	if totalDiff == 0 {
		return 0, nil
	}
	
	usage := float64(activeDiff) / float64(totalDiff) * 100
	return usage, nil
}

func GetTemperature() (int, error) {
	// Find k10temp hwmon device
	hwmonDirs, err := filepath.Glob("/sys/class/hwmon/hwmon*")
	if err != nil {
		return 0, err
	}
	
	for _, hwmonDir := range hwmonDirs {
		nameFile := filepath.Join(hwmonDir, "name")
		nameData, err := os.ReadFile(nameFile)
		if err != nil {
			continue
		}
		
		if strings.TrimSpace(string(nameData)) == "k10temp" {
			tempFile := filepath.Join(hwmonDir, "temp1_input")
			tempData, err := os.ReadFile(tempFile)
			if err != nil {
				continue
			}
			
			tempMillidegrees, err := strconv.ParseInt(strings.TrimSpace(string(tempData)), 10, 64)
			if err != nil {
				continue
			}
			
			return int(tempMillidegrees / 1000), nil
		}
	}
	
	return 0, errors.New("k10temp not found")
}

func GetFrequency() (float64, error) {
	cpuDirs, err := filepath.Glob("/sys/devices/system/cpu/cpu*/cpufreq/scaling_cur_freq")
	if err != nil {
		return 0, err
	}
	
	if len(cpuDirs) == 0 {
		return 0, errors.New("no CPU frequency info available")
	}
	
	var totalFreq float64
	var count int
	
	for _, freqFile := range cpuDirs {
		data, err := os.ReadFile(freqFile)
		if err != nil {
			continue
		}
		
		freqKHz, err := strconv.ParseFloat(strings.TrimSpace(string(data)), 64)
		if err != nil {
			continue
		}
		
		totalFreq += freqKHz
		count++
	}
	
	if count == 0 {
		return 0, errors.New("no valid CPU frequency data")
	}
	
	// Convert kHz to GHz and return average
	avgFreqGHz := (totalFreq / float64(count)) / 1000000
	return avgFreqGHz, nil
}

func GetCores() (int, error) {
	return runtime.NumCPU(), nil
}

func GetAllMetrics() (*Metrics, error) {
	usage, err := GetUsage()
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
	
	cores, err := GetCores()
	if err != nil {
		return nil, err
	}
	
	return &Metrics{
		Usage:       usage,
		Temperature: temp,
		Frequency:   freq,
		Cores:       cores,
	}, nil
}