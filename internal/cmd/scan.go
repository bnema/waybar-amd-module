// Package cmd provides CLI commands for monitoring AMD hardware metrics
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "Scan and update hardware path cache",
	Long:  "Force a complete rescan of AMD hardware paths and update the cache file",
	RunE: func(cmd *cobra.Command, args []string) error {
		if pathCache == nil {
			return errors.New("path cache not initialized")
		}

		fmt.Println("Scanning AMD hardware...")
		if err := pathCache.ForceRescan(); err != nil {
			return errors.New("failed to scan hardware: " + err.Error())
		}

		fmt.Println("Hardware scan completed successfully!")
		fmt.Printf("Cache updated: %s\n", pathCache.GetCacheFile())
		return nil
	},
}