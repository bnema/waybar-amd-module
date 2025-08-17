// Package discovery provides hardware path discovery and caching for AMD GPUs and CPUs
package discovery

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"time"
)

// GPUPaths contains all discovered GPU-related paths
type GPUPaths struct {
	Card   string `json:"card"`
	HwMon  string `json:"hwmon"`
	Device string `json:"device"`
}

// CPUPaths contains all discovered CPU-related paths
type CPUPaths struct {
	HwMon      string `json:"hwmon"`
	SensorType string `json:"sensor_type"`
	CPUFreqBase string `json:"cpufreq_base"`
	BoostPath  string `json:"boost_path"`
	CoreCount  int    `json:"core_count"`
}

// PowerPaths contains power-related paths
type PowerPaths struct {
	Battery string `json:"battery"`
	RAPL    string `json:"rapl"`
}

// SystemInfo contains system metadata
type SystemInfo struct {
	Kernel      string `json:"kernel"`
	AMDCpu      bool   `json:"amd_cpu"`
	AMDGpuCount int    `json:"amd_gpu_count"`
}

// PathCache represents the complete cached path information
type PathCache struct {
	Version   string     `json:"version"`
	Timestamp time.Time  `json:"timestamp"`
	System    SystemInfo `json:"system"`
	GPU       *GPUPaths  `json:"gpu"`
	CPU       *CPUPaths  `json:"cpu"`
	Power     *PowerPaths `json:"power"`
	
	cacheFile string
}

// NewPathCache creates a new PathCache instance
func NewPathCache() (*PathCache, error) {
	cacheDir, err := getCacheDir()
	if err != nil {
		return nil, errors.New("failed to get cache directory: " + err.Error())
	}

	cacheFile := filepath.Join(cacheDir, "paths.json")
	
	cache := &PathCache{
		Version:   "1.0",
		Timestamp: time.Now(),
		cacheFile: cacheFile,
	}

	// Try to load existing cache
	if err := cache.Load(); err != nil {
		// If cache doesn't exist or is invalid, perform scan
		if err := cache.Scan(); err != nil {
			return nil, errors.New("failed to scan hardware paths: " + err.Error())
		}
		// Save the newly discovered paths
		if err := cache.Save(); err != nil {
			return nil, errors.New("failed to save cache: " + err.Error())
		}
	}

	return cache, nil
}

// getCacheDir returns the XDG cache directory for the application
func getCacheDir() (string, error) {
	cacheHome := os.Getenv("XDG_CACHE_HOME")
	if cacheHome == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		cacheHome = filepath.Join(homeDir, ".cache")
	}

	cacheDir := filepath.Join(cacheHome, "waybar-amd-module")
	
	// Create directory if it doesn't exist
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return "", err
	}

	return cacheDir, nil
}

// Load reads the cache from the filesystem
func (c *PathCache) Load() error {
	data, err := os.ReadFile(c.cacheFile)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, c); err != nil {
		return err
	}

	// Validate that cached paths still exist
	if !c.Validate() {
		return errors.New("cached paths are no longer valid")
	}

	return nil
}

// Save writes the cache to the filesystem
func (c *PathCache) Save() error {
	c.Timestamp = time.Now()
	
	data, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(c.cacheFile, data, 0644)
}

// GetCacheFile returns the path to the cache file
func (c *PathCache) GetCacheFile() string {
	return c.cacheFile
}