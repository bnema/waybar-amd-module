// Package cmd provides CLI commands for monitoring AMD hardware metrics
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/bnema/waybar-amd-module/internal/discovery"
)

var (
	formatFlag    string
	nerdFontFlag  bool
	
	pathCache *discovery.PathCache
)

var rootCmd = &cobra.Command{
	Use:   "waybar-amd-module",
	Short: "AMD GPU and CPU metrics for Waybar",
	Long:  "A tool to monitor AMD GPU and CPU power, temperature, frequency, and utilization for Waybar",
}

func init() {
	rootCmd.PersistentFlags().StringVar(&formatFlag, "format", "json", "Output format (json/text)")
	rootCmd.PersistentFlags().BoolVar(&nerdFontFlag, "nerd-font", false, "Use nerd font symbols in output")
	
	rootCmd.AddCommand(gpuCmd)
	rootCmd.AddCommand(cpuCmd)
	rootCmd.AddCommand(scanCmd)
}

// SetPathCache sets the path cache for use by commands
func SetPathCache(cache *discovery.PathCache) {
	pathCache = cache
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}