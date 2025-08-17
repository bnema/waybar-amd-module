// Package main provides the entry point for the waybar-amd-module application
package main

import (
	"log"
	"os"

	"github.com/bnema/waybar-amd-module/internal/cmd"
	"github.com/bnema/waybar-amd-module/internal/cpu"
	"github.com/bnema/waybar-amd-module/internal/discovery"
	"github.com/bnema/waybar-amd-module/internal/gpu"
)

func main() {
	// Initialize hardware path discovery
	cache, err := discovery.Initialize()
	if err != nil {
		log.Fatalf("Failed to initialize hardware discovery: %v", err)
	}

	// Initialize GPU package if GPU is available
	if cache.GPU != nil {
		if err := gpu.Initialize(cache); err != nil {
			log.Printf("Warning: Failed to initialize GPU: %v", err)
		}
	}

	// Initialize CPU package if CPU is available
	if cache.CPU != nil {
		if err := cpu.Initialize(cache); err != nil {
			log.Printf("Warning: Failed to initialize CPU: %v", err)
		}
	}

	// Set the cache in the command package for access to rescan functionality
	cmd.SetPathCache(cache)

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}