// Package formatting provides common output formatting utilities for Waybar JSON output
package formatting

import (
	"encoding/json"
	"fmt"
)

// WaybarOutput represents the JSON structure expected by Waybar
type WaybarOutput struct {
	Text    string `json:"text"`
	Tooltip string `json:"tooltip"`
	Class   string `json:"class"`
}

// ValidateNoTooltipFlag checks if --no-tooltip is used with text format
func ValidateNoTooltipFlag(noTooltipFlag bool, formatFlag string) bool {
	switch {
	case noTooltipFlag && formatFlag != "json":
		fmt.Println("Error: --no-tooltip flag can only be used with --format=json")
		return false
	default:
		return true
	}
}

// FormatJSONOutput formats output for JSON mode, handling --no-tooltip flag
func FormatJSONOutput(text string, tooltip string, class string, noTooltipFlag bool) {
	switch {
	case noTooltipFlag:
		// Simple JSON output without tooltip
		output := map[string]any{
			"text":  text,
			"class": class,
		}
		jsonData, _ := json.Marshal(output)
		fmt.Println(string(jsonData))
	default:
		// Standard Waybar output with tooltip
		output := WaybarOutput{
			Text:    text,
			Tooltip: tooltip,
			Class:   class,
		}
		jsonData, _ := json.Marshal(output)
		fmt.Println(string(jsonData))
	}
}