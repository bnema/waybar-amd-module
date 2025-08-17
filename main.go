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
	// Initialize hardware discovery
	cache, err := discovery.Initialize()
	if err != nil {
		log.Fatalf("Failed to initialize hardware discovery: %v", err)
	}

	// Set the cache for commands
	cmd.SetPathCache(cache)

	// Initialize GPU and CPU packages with discovered paths
	if cache.GPU != nil {
		if err := gpu.Initialize(cache); err != nil {
			log.Printf("Warning: GPU initialization failed: %v", err)
		}
	}

	if cache.CPU != nil {
		if err := cpu.Initialize(cache); err != nil {
			log.Printf("Warning: CPU initialization failed: %v", err)
		}
	}

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
