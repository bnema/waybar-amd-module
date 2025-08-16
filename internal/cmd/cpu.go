package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/bnema/waybar-amd-module/internal/cpu"
	"github.com/bnema/waybar-amd-module/internal/nerdfonts"
)

func formatCPUWithSymbols(usage float64, temp int, freq float64, cores int) (string, string) {
	var text, tooltip string
	
	if nerdFontFlag {
		text = fmt.Sprintf("%s %.1f%% %s %d°C %s %.1fGHz %s %d", 
			nerdfonts.CPUUsage, usage, 
			nerdfonts.CPUTemp, temp, 
			nerdfonts.CPUFreq, freq, 
			nerdfonts.CPUCores, cores)
		tooltip = fmt.Sprintf("%s Usage: %.1f%%\n%s Temp: %d°C\n%s Freq: %.1fGHz\n%s Cores: %d", 
			nerdfonts.CPUUsage, usage, 
			nerdfonts.CPUTemp, temp, 
			nerdfonts.CPUFreq, freq, 
			nerdfonts.CPUCores, cores)
	} else {
		text = fmt.Sprintf("%.1f%% %d°C %.1fGHz %d cores", usage, temp, freq, cores)
		tooltip = fmt.Sprintf("Usage: %.1f%%\nTemp: %d°C\nFreq: %.1fGHz\nCores: %d", usage, temp, freq, cores)
	}
	
	return text, tooltip
}

func formatCPUUsage(usage float64) string {
	if nerdFontFlag {
		return fmt.Sprintf("%s %.1f%%", nerdfonts.CPUUsage, usage)
	}
	return fmt.Sprintf("%.1f%%", usage)
}

func formatCPUTemp(temp int) string {
	if nerdFontFlag {
		return fmt.Sprintf("%s %d°C", nerdfonts.CPUTemp, temp)
	}
	return fmt.Sprintf("%d°C", temp)
}

func formatCPUFreq(freq float64) string {
	if nerdFontFlag {
		return fmt.Sprintf("%s %.1fGHz", nerdfonts.CPUFreq, freq)
	}
	return fmt.Sprintf("%.1fGHz", freq)
}

func formatCPUCores(cores int) string {
	if nerdFontFlag {
		return fmt.Sprintf("%s %d", nerdfonts.CPUCores, cores)
	}
	return fmt.Sprintf("%d cores", cores)
}

var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "AMD CPU monitoring commands",
	Long:  "Monitor AMD CPU usage, temperature, frequency, and core count",
}

var cpuAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Get all CPU metrics",
	Run: func(cmd *cobra.Command, args []string) {
		metrics, err := cpu.GetAllMetrics()
		if err != nil {
			switch formatFlag {
			case "json":
				fmt.Println("{}")
			default:
				return
			}
			return
		}

		text, tooltip := formatCPUWithSymbols(metrics.Usage, metrics.Temperature, metrics.Frequency, metrics.Cores)
		
		switch formatFlag {
		case "json":
			output := cpu.WaybarOutput{
				Text:    text,
				Tooltip: tooltip,
				Class:   "custom-cpu",
			}
			jsonData, _ := json.Marshal(output)
			fmt.Println(string(jsonData))
		default:
			fmt.Println(text)
		}
	},
}

var cpuUsageCmd = &cobra.Command{
	Use:   "usage",
	Short: "Get CPU usage percentage",
	Run: func(cmd *cobra.Command, args []string) {
		usage, err := cpu.GetUsage()
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
			metrics, err := cpu.GetAllMetrics()
			if err != nil {
				fmt.Println("{}")
				return
			}
			
			text := formatCPUUsage(usage)
			_, tooltip := formatCPUWithSymbols(metrics.Usage, metrics.Temperature, metrics.Frequency, metrics.Cores)
			
			output := cpu.WaybarOutput{
				Text:    text,
				Tooltip: tooltip,
				Class:   "custom-cpu",
			}
			jsonData, _ := json.Marshal(output)
			fmt.Println(string(jsonData))
		default:
			fmt.Println(formatCPUUsage(usage))
		}
	},
}

var cpuTempCmd = &cobra.Command{
	Use:   "temp",
	Short: "Get CPU temperature",
	Run: func(cmd *cobra.Command, args []string) {
		temp, err := cpu.GetTemperature()
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
			metrics, err := cpu.GetAllMetrics()
			if err != nil {
				fmt.Println("{}")
				return
			}
			
			text := formatCPUTemp(temp)
			_, tooltip := formatCPUWithSymbols(metrics.Usage, metrics.Temperature, metrics.Frequency, metrics.Cores)
			
			output := cpu.WaybarOutput{
				Text:    text,
				Tooltip: tooltip,
				Class:   "custom-cpu",
			}
			jsonData, _ := json.Marshal(output)
			fmt.Println(string(jsonData))
		default:
			fmt.Println(formatCPUTemp(temp))
		}
	},
}

var cpuFreqCmd = &cobra.Command{
	Use:   "freq",
	Short: "Get CPU frequency",
	Run: func(cmd *cobra.Command, args []string) {
		freq, err := cpu.GetFrequency()
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
			metrics, err := cpu.GetAllMetrics()
			if err != nil {
				fmt.Println("{}")
				return
			}
			
			text := formatCPUFreq(freq)
			_, tooltip := formatCPUWithSymbols(metrics.Usage, metrics.Temperature, metrics.Frequency, metrics.Cores)
			
			output := cpu.WaybarOutput{
				Text:    text,
				Tooltip: tooltip,
				Class:   "custom-cpu",
			}
			jsonData, _ := json.Marshal(output)
			fmt.Println(string(jsonData))
		default:
			fmt.Println(formatCPUFreq(freq))
		}
	},
}

var cpuCoresCmd = &cobra.Command{
	Use:   "cores",
	Short: "Get CPU core count",
	Run: func(cmd *cobra.Command, args []string) {
		cores, err := cpu.GetCores()
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
			metrics, err := cpu.GetAllMetrics()
			if err != nil {
				fmt.Println("{}")
				return
			}
			
			text := formatCPUCores(cores)
			_, tooltip := formatCPUWithSymbols(metrics.Usage, metrics.Temperature, metrics.Frequency, metrics.Cores)
			
			output := cpu.WaybarOutput{
				Text:    text,
				Tooltip: tooltip,
				Class:   "custom-cpu",
			}
			jsonData, _ := json.Marshal(output)
			fmt.Println(string(jsonData))
		default:
			fmt.Println(formatCPUCores(cores))
		}
	},
}

func init() {
	cpuCmd.AddCommand(cpuAllCmd)
	cpuCmd.AddCommand(cpuUsageCmd)
	cpuCmd.AddCommand(cpuTempCmd)
	cpuCmd.AddCommand(cpuFreqCmd)
	cpuCmd.AddCommand(cpuCoresCmd)
}