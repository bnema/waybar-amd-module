// Package discovery provides path validation utilities
package discovery

import (
	"os"
	"path/filepath"
	"strings"
)

// Validate checks if all cached paths still exist and are accessible
func (c *PathCache) Validate() bool {
	if c.GPU != nil && !c.validateGPU() {
		return false
	}

	if c.CPU != nil && !c.validateCPU() {
		return false
	}

	if c.Power != nil && !c.validatePower() {
		return false
	}

	return true
}

// validateGPU checks if GPU paths are valid
func (c *PathCache) validateGPU() bool {
	if c.GPU == nil {
		return true
	}

	// Check hwmon path exists
	if c.GPU.HwMon != "" {
		if !pathExists(c.GPU.HwMon) {
			return false
		}

		// Check essential metric files
		essentialFiles := []string{
			"power1_input",
			"temp1_input",
			"freq1_input",
		}

		for _, file := range essentialFiles {
			path := filepath.Join(c.GPU.HwMon, file)
			if !pathExists(path) {
				return false
			}
		}
	}

	// Check device path exists
	if c.GPU.Device != "" && !pathExists(c.GPU.Device) {
		return false
	}

	// Check card path exists
	if c.GPU.Card != "" && !pathExists(c.GPU.Card) {
		return false
	}

	return true
}

// validateCPU checks if CPU paths are valid
func (c *PathCache) validateCPU() bool {
	if c.CPU == nil {
		return true
	}

	// Check hwmon path exists
	if c.CPU.HwMon != "" {
		if !pathExists(c.CPU.HwMon) {
			return false
		}

		// Check temperature file exists
		tempFile := filepath.Join(c.CPU.HwMon, "temp1_input")
		if !pathExists(tempFile) {
			return false
		}
	}

	// Check cpufreq base path exists
	if c.CPU.CPUFreqBase != "" && !pathExists(c.CPU.CPUFreqBase) {
		return false
	}

	// Check if at least one CPU has frequency scaling
	if c.CPU.CPUFreqBase != "" {
		freqFiles, err := filepath.Glob(filepath.Join(c.CPU.CPUFreqBase, "cpu*/cpufreq/scaling_cur_freq"))
		if err != nil || len(freqFiles) == 0 {
			return false
		}
	}

	// Boost path is optional, so don't fail validation if it doesn't exist

	return true
}

// validatePower checks if power paths are valid
func (c *PathCache) validatePower() bool {
	if c.Power == nil {
		return true
	}

	// Battery path is optional
	if c.Power.Battery != "" && !pathExists(c.Power.Battery) {
		return false
	}

	// RAPL path is optional
	if c.Power.RAPL != "" && !pathExists(c.Power.RAPL) {
		return false
	}

	return true
}

// pathExists checks if a path exists
func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// ValidateSystemPath validates that a path is within expected system directories
func ValidateSystemPath(path string) bool {
	if path == "" {
		return false
	}

	// Clean the path to prevent traversal attacks
	cleanPath := filepath.Clean(path)

	// Must be within /sys or /proc
	validPrefixes := []string{
		"/sys/",
		"/proc/",
	}

	for _, prefix := range validPrefixes {
		if strings.HasPrefix(cleanPath, prefix) {
			// Check for path traversal attempts
			if !containsPathTraversal(cleanPath) {
				return true
			}
		}
	}

	return false
}

// containsPathTraversal checks if a path contains traversal attempts
func containsPathTraversal(path string) bool {
	return filepath.Base(path) == ".." ||
		filepath.Dir(path) != path && containsPathTraversal(filepath.Dir(path))
}
