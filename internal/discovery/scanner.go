// Package discovery provides hardware path discovery algorithms
package discovery

import (
	"errors"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// Scan performs a full system scan to discover AMD hardware paths
func (c *PathCache) Scan() error {
	// Get system information
	c.System = SystemInfo{
		Kernel: getKernelVersion(),
	}

	// Scan for AMD CPU
	if cpuPaths, err := c.scanCPU(); err == nil {
		c.CPU = cpuPaths
		c.System.AMDCpu = true
	}

	// Scan for AMD GPU
	if gpuPaths, err := c.scanGPU(); err == nil {
		c.GPU = gpuPaths
		c.System.AMDGpuCount = 1
	}

	// Scan for power paths
	if powerPaths, err := c.scanPower(); err == nil {
		c.Power = powerPaths
	}

	// Ensure we found at least something
	if c.CPU == nil && c.GPU == nil {
		return errors.New("no AMD hardware detected")
	}

	return nil
}

// scanGPU discovers AMD GPU paths using multiple methods
func (c *PathCache) scanGPU() (*GPUPaths, error) {
	// Method 1: Scan DRM cards
	if gpu, err := c.scanDRMCards(); err == nil {
		return gpu, nil
	}

	// Method 2: Scan PCI drivers directly
	if gpu, err := c.scanPCIDrivers(); err == nil {
		return gpu, nil
	}

	return nil, errors.New("no AMD GPU found")
}

// scanDRMCards scans /sys/class/drm/card* for AMD GPUs
func (c *PathCache) scanDRMCards() (*GPUPaths, error) {
	cardDirs, err := filepath.Glob("/sys/class/drm/card*")
	if err != nil {
		return nil, err
	}

	for _, cardDir := range cardDirs {
		driverPath := filepath.Join(cardDir, "device", "driver")
		if target, err := os.Readlink(driverPath); err == nil {
			if strings.Contains(target, "amdgpu") {
				// Found AMD GPU, now find hwmon
				hwmonDirs, err := filepath.Glob(filepath.Join(cardDir, "device", "hwmon", "hwmon*"))
				if err == nil && len(hwmonDirs) > 0 {
					gpu := &GPUPaths{
						Card:   cardDir,
						HwMon:  hwmonDirs[0],
						Device: filepath.Join(cardDir, "device"),
					}

					// Validate essential files exist
					if c.validateGPUPaths(gpu) {
						return gpu, nil
					}
				}
			}
		}
	}

	return nil, errors.New("no valid AMD GPU found in DRM cards")
}

// scanPCIDrivers scans PCI bus for AMD GPU drivers
func (c *PathCache) scanPCIDrivers() (*GPUPaths, error) {
	pciDirs, err := filepath.Glob("/sys/bus/pci/drivers/amdgpu/*/hwmon/hwmon*")
	if err != nil || len(pciDirs) == 0 {
		return nil, errors.New("no AMD GPU found in PCI drivers")
	}

	hwmonPath := pciDirs[0]
	devicePath := filepath.Dir(filepath.Dir(hwmonPath))

	// Find corresponding card
	cardPath := ""
	if cardDirs, err := filepath.Glob("/sys/class/drm/card*"); err == nil {
		for _, card := range cardDirs {
			if cardDevice := filepath.Join(card, "device"); cardDevice == devicePath {
				cardPath = card
				break
			}
		}
	}

	gpu := &GPUPaths{
		Card:   cardPath,
		HwMon:  hwmonPath,
		Device: devicePath,
	}

	if c.validateGPUPaths(gpu) {
		return gpu, nil
	}

	return nil, errors.New("found AMD GPU but validation failed")
}

// validateGPUPaths checks if essential GPU metric files exist
func (c *PathCache) validateGPUPaths(gpu *GPUPaths) bool {
	essentialFiles := []string{
		"power1_input",
		"temp1_input",
		"freq1_input",
	}

	for _, file := range essentialFiles {
		path := filepath.Join(gpu.HwMon, file)
		if _, err := os.Stat(path); err != nil {
			return false
		}
	}

	// Check device files (optional - don't fail validation if missing)
	deviceFiles := []string{
		"gpu_busy_percent",
		"mem_info_vram_total",
	}

	for _, file := range deviceFiles {
		path := filepath.Join(gpu.Device, file)
		if _, err := os.Stat(path); err != nil {
			// These are optional, continue checking others
			continue
		}
	}

	return true
}

// scanCPU discovers AMD CPU paths
func (c *PathCache) scanCPU() (*CPUPaths, error) {
	// Check if this is an AMD CPU
	if !c.isAMDCpu() {
		return nil, errors.New("not an AMD CPU")
	}

	// Find temperature sensor
	hwmonPath, sensorType, err := c.findCPUTempSensor()
	if err != nil {
		return nil, err
	}

	// Find boost path
	boostPath := c.findBoostPath()

	cpu := &CPUPaths{
		HwMon:       hwmonPath,
		SensorType:  sensorType,
		CPUFreqBase: "/sys/devices/system/cpu",
		BoostPath:   boostPath,
		CoreCount:   runtime.NumCPU(),
	}

	return cpu, nil
}

// isAMDCpu checks if the CPU is from AMD
func (c *PathCache) isAMDCpu() bool {
	data, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		return false
	}

	return strings.Contains(string(data), "AuthenticAMD")
}

// findCPUTempSensor finds the appropriate temperature sensor for AMD CPUs
func (c *PathCache) findCPUTempSensor() (string, string, error) {
	hwmonDirs, err := filepath.Glob("/sys/class/hwmon/hwmon*")
	if err != nil {
		return "", "", err
	}

	// Look for k10temp (modern AMD) or fam15h_power (older AMD)
	preferredSensors := []string{"k10temp", "fam15h_power"}

	for _, sensor := range preferredSensors {
		for _, hwmonDir := range hwmonDirs {
			nameFile := filepath.Join(hwmonDir, "name")
			nameData, err := os.ReadFile(nameFile)
			if err != nil {
				continue
			}

			if strings.TrimSpace(string(nameData)) == sensor {
				// Verify temp file exists
				tempFile := filepath.Join(hwmonDir, "temp1_input")
				if _, err := os.Stat(tempFile); err == nil {
					return hwmonDir, sensor, nil
				}
			}
		}
	}

	return "", "", errors.New("no AMD CPU temperature sensor found")
}

// findBoostPath finds the CPU boost control path
func (c *PathCache) findBoostPath() string {
	// Check various possible boost paths
	boostPaths := []string{
		"/sys/devices/system/cpu/cpufreq/boost",
		"/sys/devices/system/cpu/cpu0/cpufreq/boost",
	}

	for _, path := range boostPaths {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	return ""
}

// scanPower discovers power-related paths
func (c *PathCache) scanPower() (*PowerPaths, error) {
	power := &PowerPaths{}

	// Find battery
	if batteryPath, err := c.findBattery(); err == nil {
		power.Battery = batteryPath
	}

	// Find RAPL (if available)
	if raplPath, err := c.findRAPL(); err == nil {
		power.RAPL = raplPath
	}

	// Return even if empty - power monitoring is optional
	return power, nil
}

// findBattery finds the primary battery path
func (c *PathCache) findBattery() (string, error) {
	powerSupplyDirs, err := filepath.Glob("/sys/class/power_supply/*")
	if err != nil {
		return "", err
	}

	for _, dir := range powerSupplyDirs {
		typePath := filepath.Join(dir, "type")
		if typeData, err := os.ReadFile(typePath); err == nil {
			if strings.TrimSpace(string(typeData)) == "Battery" {
				return dir, nil
			}
		}
	}

	return "", errors.New("no battery found")
}

// findRAPL finds RAPL power monitoring paths (Intel/AMD)
func (c *PathCache) findRAPL() (string, error) {
	raplPaths := []string{
		"/sys/class/powercap/intel-rapl/intel-rapl:0",
		"/sys/class/powercap/amd-rapl/amd-rapl:0",
	}

	for _, path := range raplPaths {
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	return "", errors.New("no RAPL found")
}

// getKernelVersion gets the kernel version
func getKernelVersion() string {
	if data, err := os.ReadFile("/proc/version"); err == nil {
		fields := strings.Fields(string(data))
		if len(fields) >= 3 {
			return fields[2]
		}
	}
	return "unknown"
}
