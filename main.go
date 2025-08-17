// Package main provides the entry point for the waybar-amd-module application
package main

import (
	"os"

	"github.com/bnema/waybar-amd-module/internal/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
