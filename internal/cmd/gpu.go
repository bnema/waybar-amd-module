package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/bnema/waybar-amd-module/internal/gpu"
	"github.com/bnema/waybar-amd-module/internal/nerdfonts"
)

func formatWithSymbols(metrics *gpu.Metrics) (string, string) {
	var text string
	
	if nerdFontFlag {
		text = fmt.Sprintf("%s %.1fW %s %d°C %s %.1fGHz %s %d%%", 
			nerdfonts.GPUPower, metrics.Power, 
			nerdfonts.GPUTemp, metrics.Temperature, 
			nerdfonts.GPUFreq, metrics.Frequency, 
			nerdfonts.GPUUtil, metrics.Utilization)
		
		tooltipLines := []string{
			fmt.Sprintf("%s Power: %.1fW", nerdfonts.GPUPower, metrics.Power),
			fmt.Sprintf("%s Temp: %d°C", nerdfonts.GPUTemp, metrics.Temperature),
			fmt.Sprintf("%s Freq: %.1fGHz", nerdfonts.GPUFreq, metrics.Frequency),
			fmt.Sprintf("%s Util: %d%%", nerdfonts.GPUUtil, metrics.Utilization),
			fmt.Sprintf("%s Memory: %.1f%%", nerdfonts.GPUMemory, metrics.MemoryUsage),
			fmt.Sprintf("%s Fan: %d RPM", nerdfonts.GPUFan, metrics.FanSpeed),
			fmt.Sprintf("%s Voltage: %.2fV", nerdfonts.GPUVoltage, metrics.Voltage),
			fmt.Sprintf("%s Junction: %d°C", nerdfonts.GPUTemp, metrics.JunctionTemp),
			fmt.Sprintf("%s Memory Temp: %d°C", nerdfonts.GPUTemp, metrics.MemoryTemp),
			fmt.Sprintf("%s Power Cap: %.1fW", nerdfonts.GPUPower, metrics.PowerCap),
		}
		return text, strings.Join(tooltipLines, "\n")
	} else {
		text = fmt.Sprintf("%.1fW %d°C %.1fGHz %d%%", metrics.Power, metrics.Temperature, metrics.Frequency, metrics.Utilization)
		
		tooltipLines := []string{
			fmt.Sprintf("Power: %.1fW", metrics.Power),
			fmt.Sprintf("Temp: %d°C", metrics.Temperature),
			fmt.Sprintf("Freq: %.1fGHz", metrics.Frequency),
			fmt.Sprintf("Util: %d%%", metrics.Utilization),
			fmt.Sprintf("Memory: %.1f%%", metrics.MemoryUsage),
			fmt.Sprintf("Fan: %d RPM", metrics.FanSpeed),
			fmt.Sprintf("Voltage: %.2fV", metrics.Voltage),
			fmt.Sprintf("Junction: %d°C", metrics.JunctionTemp),
			fmt.Sprintf("Memory Temp: %d°C", metrics.MemoryTemp),
			fmt.Sprintf("Power Cap: %.1fW", metrics.PowerCap),
		}
		return text, strings.Join(tooltipLines, "\n")
	}
}

func formatPower(power float64) string {
	if nerdFontFlag {
		return fmt.Sprintf("%s %.1fW", nerdfonts.GPUPower, power)
	}
	return fmt.Sprintf("%.1fW", power)
}

func formatTemp(temp int) string {
	if nerdFontFlag {
		return fmt.Sprintf("%s %d°C", nerdfonts.GPUTemp, temp)
	}
	return fmt.Sprintf("%d°C", temp)
}

func formatFreq(freq float64) string {
	if nerdFontFlag {
		return fmt.Sprintf("%s %.1fGHz", nerdfonts.GPUFreq, freq)
	}
	return fmt.Sprintf("%.1fGHz", freq)
}

func formatUtil(util int) string {
	if nerdFontFlag {
		return fmt.Sprintf("%s %d%%", nerdfonts.GPUUtil, util)
	}
	return fmt.Sprintf("%d%%", util)
}

var gpuCmd = &cobra.Command{
	Use:   "gpu",
	Short: "AMD GPU monitoring commands",
	Long:  "Monitor AMD GPU power, temperature, frequency, and utilization",
}

var gpuAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Get all GPU metrics",
	Run: func(cmd *cobra.Command, args []string) {
		metrics, err := gpu.GetAllMetrics()
		if err != nil {
			switch formatFlag {
			case "json":
				fmt.Println("{}")
			default:
				return
			}
			return
		}

		text, tooltip := formatWithSymbols(metrics)
		
		switch formatFlag {
		case "json":
			output := gpu.WaybarOutput{
				Text:    text,
				Tooltip: tooltip,
				Class:   "custom-gpu",
			}
			jsonData, _ := json.Marshal(output)
			fmt.Println(string(jsonData))
		default:
			fmt.Println(text)
		}
	},
}

var gpuPowerCmd = &cobra.Command{
	Use:   "power",
	Short: "Get GPU power consumption",
	Run: func(cmd *cobra.Command, args []string) {
		power, err := gpu.GetPower()
		if err != nil {
			switch formatFlag {
			case "json":
				fmt.Println("{}")
			default:
				return
			}
			return
		}

		switch formatFlag {
		case "json":
			// Get all metrics for tooltip
			metrics, err := gpu.GetAllMetrics()
			if err != nil {
				fmt.Println("{}")
				return
			}
			
			text := formatPower(power)
			_, tooltip := formatWithSymbols(metrics)
			
			output := gpu.WaybarOutput{
				Text:    text,
				Tooltip: tooltip,
				Class:   "custom-gpu",
			}
			jsonData, _ := json.Marshal(output)
			fmt.Println(string(jsonData))
		default:
			fmt.Println(formatPower(power))
		}
	},
}

var gpuTempCmd = &cobra.Command{
	Use:   "temp",
	Short: "Get GPU temperature",
	Run: func(cmd *cobra.Command, args []string) {
		temp, err := gpu.GetTemperature()
		if err != nil {
			switch formatFlag {
			case "json":
				fmt.Println("{}")
			default:
				return
			}
			return
		}

		switch formatFlag {
		case "json":
			// Get all metrics for tooltip
			metrics, err := gpu.GetAllMetrics()
			if err != nil {
				fmt.Println("{}")
				return
			}
			
			text := formatTemp(temp)
			_, tooltip := formatWithSymbols(metrics)
			
			output := gpu.WaybarOutput{
				Text:    text,
				Tooltip: tooltip,
				Class:   "custom-gpu",
			}
			jsonData, _ := json.Marshal(output)
			fmt.Println(string(jsonData))
		default:
			fmt.Println(formatTemp(temp))
		}
	},
}

var gpuFreqCmd = &cobra.Command{
	Use:   "freq",
	Short: "Get GPU frequency",
	Run: func(cmd *cobra.Command, args []string) {
		freq, err := gpu.GetFrequency()
		if err != nil {
			switch formatFlag {
			case "json":
				fmt.Println("{}")
			default:
				return
			}
			return
		}

		switch formatFlag {
		case "json":
			// Get all metrics for tooltip
			metrics, err := gpu.GetAllMetrics()
			if err != nil {
				fmt.Println("{}")
				return
			}
			
			text := formatFreq(freq)
			_, tooltip := formatWithSymbols(metrics)
			
			output := gpu.WaybarOutput{
				Text:    text,
				Tooltip: tooltip,
				Class:   "custom-gpu",
			}
			jsonData, _ := json.Marshal(output)
			fmt.Println(string(jsonData))
		default:
			fmt.Println(formatFreq(freq))
		}
	},
}

var gpuUtilCmd = &cobra.Command{
	Use:   "util",
	Short: "Get GPU utilization",
	Run: func(cmd *cobra.Command, args []string) {
		util, err := gpu.GetUtilization()
		if err != nil {
			switch formatFlag {
			case "json":
				fmt.Println("{}")
			default:
				return
			}
			return
		}

		switch formatFlag {
		case "json":
			// Get all metrics for tooltip
			metrics, err := gpu.GetAllMetrics()
			if err != nil {
				fmt.Println("{}")
				return
			}
			
			text := formatUtil(util)
			_, tooltip := formatWithSymbols(metrics)
			
			output := gpu.WaybarOutput{
				Text:    text,
				Tooltip: tooltip,
				Class:   "custom-gpu",
			}
			jsonData, _ := json.Marshal(output)
			fmt.Println(string(jsonData))
		default:
			fmt.Println(formatUtil(util))
		}
	},
}

func init() {
	gpuCmd.AddCommand(gpuAllCmd)
	gpuCmd.AddCommand(gpuPowerCmd)
	gpuCmd.AddCommand(gpuTempCmd)
	gpuCmd.AddCommand(gpuFreqCmd)
	gpuCmd.AddCommand(gpuUtilCmd)
}