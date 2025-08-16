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
	MemoryUsage float64 `json:"memory_usage"`
	LoadAvg     float64 `json:"load_avg"`
	Governor    string  `json:"governor"`
	BoostEnabled bool   `json:"boost_enabled"`
	MinFreq     float64 `json:"min_freq"`
	MaxFreq     float64 `json:"max_freq"`
	IOWait      float64 `json:"io_wait"`
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

func GetMemoryUsage() (float64, error) {
	data, err := os.ReadFile("/proc/meminfo")
	if err != nil {
		return 0, err
	}
	
	lines := strings.Split(string(data), "\n")
	var memTotal, memAvailable uint64
	
	for _, line := range lines {
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		
		switch fields[0] {
		case "MemTotal:":
			memTotal, _ = strconv.ParseUint(fields[1], 10, 64)
		case "MemAvailable:":
			memAvailable, _ = strconv.ParseUint(fields[1], 10, 64)
		}
		
		if memTotal > 0 && memAvailable > 0 {
			break
		}
	}
	
	if memTotal == 0 {
		return 0, errors.New("could not parse memory info")
	}
	
	memUsed := memTotal - memAvailable
	return float64(memUsed) / float64(memTotal) * 100, nil
}

func GetLoadAverage() (float64, error) {
	data, err := os.ReadFile("/proc/loadavg")
	if err != nil {
		return 0, err
	}
	
	fields := strings.Fields(string(data))
	if len(fields) < 1 {
		return 0, errors.New("invalid loadavg format")
	}
	
	loadAvg, err := strconv.ParseFloat(fields[0], 64)
	if err != nil {
		return 0, err
	}
	
	return loadAvg, nil
}

func GetGovernor() (string, error) {
	data, err := os.ReadFile("/sys/devices/system/cpu/cpu0/cpufreq/scaling_governor")
	if err != nil {
		return "", err
	}
	
	return strings.TrimSpace(string(data)), nil
}

func GetBoostEnabled() (bool, error) {
	data, err := os.ReadFile("/sys/devices/system/cpu/cpu0/cpufreq/boost")
	if err != nil {
		return false, err
	}
	
	boost := strings.TrimSpace(string(data))
	return boost == "1", nil
}

func GetMinMaxFreq() (float64, float64, error) {
	minData, err := os.ReadFile("/sys/devices/system/cpu/cpu0/cpufreq/cpuinfo_min_freq")
	if err != nil {
		return 0, 0, err
	}
	
	maxData, err := os.ReadFile("/sys/devices/system/cpu/cpu0/cpufreq/cpuinfo_max_freq")
	if err != nil {
		return 0, 0, err
	}
	
	minFreqKHz, err := strconv.ParseFloat(strings.TrimSpace(string(minData)), 64)
	if err != nil {
		return 0, 0, err
	}
	
	maxFreqKHz, err := strconv.ParseFloat(strings.TrimSpace(string(maxData)), 64)
	if err != nil {
		return 0, 0, err
	}
	
	return minFreqKHz / 1000000, maxFreqKHz / 1000000, nil
}

func GetIOWait() (float64, error) {
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
	
	time.Sleep(100 * time.Millisecond)
	
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
	
	totalDiff := stat2.total() - stat1.total()
	iowaitDiff := stat2.iowait - stat1.iowait
	
	if totalDiff == 0 {
		return 0, nil
	}
	
	return float64(iowaitDiff) / float64(totalDiff) * 100, nil
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
	
	memUsage, err := GetMemoryUsage()
	if err != nil {
		memUsage = 0 // Don't fail on memory error, just set to 0
	}
	
	loadAvg, err := GetLoadAverage()
	if err != nil {
		loadAvg = 0 // Don't fail on load average error, just set to 0
	}
	
	governor, err := GetGovernor()
	if err != nil {
		governor = "unknown" // Don't fail on governor error, just set to unknown
	}
	
	boostEnabled, err := GetBoostEnabled()
	if err != nil {
		boostEnabled = false // Don't fail on boost error, just set to false
	}
	
	minFreq, maxFreq, err := GetMinMaxFreq()
	if err != nil {
		minFreq, maxFreq = 0, 0 // Don't fail on frequency range error, just set to 0
	}
	
	ioWait, err := GetIOWait()
	if err != nil {
		ioWait = 0 // Don't fail on IO wait error, just set to 0
	}
	
	return &Metrics{
		Usage:        usage,
		Temperature:  temp,
		Frequency:    freq,
		Cores:        cores,
		MemoryUsage:  memUsage,
		LoadAvg:      loadAvg,
		Governor:     governor,
		BoostEnabled: boostEnabled,
		MinFreq:      minFreq,
		MaxFreq:      maxFreq,
		IOWait:       ioWait,
	}, nil
}