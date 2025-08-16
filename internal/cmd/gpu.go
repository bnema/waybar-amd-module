package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/bnema/waybar-amd-module/internal/gpu"
	"github.com/bnema/waybar-amd-module/internal/nerdfonts"
)

func formatWithSymbols(power float64, temp int, freq float64, util int) (string, string) {
	var text, tooltip string
	
	if nerdFontFlag {
		text = fmt.Sprintf("%s %.1fW %s %d°C %s %.1fGHz %s %d%%", 
			nerdfonts.GPUPower, power, 
			nerdfonts.GPUTemp, temp, 
			nerdfonts.GPUFreq, freq, 
			nerdfonts.GPUUtil, util)
		tooltip = fmt.Sprintf("%s Power: %.1fW\n%s Temp: %d°C\n%s Freq: %.1fGHz\n%s Util: %d%%", 
			nerdfonts.GPUPower, power, 
			nerdfonts.GPUTemp, temp, 
			nerdfonts.GPUFreq, freq, 
			nerdfonts.GPUUtil, util)
	} else {
		text = fmt.Sprintf("%.1fW %d°C %.1fGHz %d%%", power, temp, freq, util)
		tooltip = fmt.Sprintf("Power: %.1fW\nTemp: %d°C\nFreq: %.1fGHz\nUtil: %d%%", power, temp, freq, util)
	}
	
	return text, tooltip
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

		text, tooltip := formatWithSymbols(metrics.Power, metrics.Temperature, metrics.Frequency, metrics.Utilization)
		
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
			_, tooltip := formatWithSymbols(metrics.Power, metrics.Temperature, metrics.Frequency, metrics.Utilization)
			
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
			_, tooltip := formatWithSymbols(metrics.Power, metrics.Temperature, metrics.Frequency, metrics.Utilization)
			
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
			_, tooltip := formatWithSymbols(metrics.Power, metrics.Temperature, metrics.Frequency, metrics.Utilization)
			
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
			_, tooltip := formatWithSymbols(metrics.Power, metrics.Temperature, metrics.Frequency, metrics.Utilization)
			
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