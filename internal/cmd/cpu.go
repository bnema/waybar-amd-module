package cmd

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/bnema/waybar-amd-module/internal/cpu"
	"github.com/bnema/waybar-amd-module/internal/nerdfonts"
)

func formatCPUWithSymbols(metrics *cpu.Metrics) (string, string) {
	var text string
	
	if nerdFontFlag {
		text = fmt.Sprintf("%s %.1f%% %s %d°C %s %.1fGHz %s %d", 
			nerdfonts.CPUUsage, metrics.Usage, 
			nerdfonts.CPUTemp, metrics.Temperature, 
			nerdfonts.CPUFreq, metrics.Frequency, 
			nerdfonts.CPUCores, metrics.Cores)
		
		tooltipLines := []string{
			fmt.Sprintf("%s Usage: %.1f%%", nerdfonts.CPUUsage, metrics.Usage),
			fmt.Sprintf("%s Temp: %d°C", nerdfonts.CPUTemp, metrics.Temperature),
			fmt.Sprintf("%s Freq: %.1fGHz", nerdfonts.CPUFreq, metrics.Frequency),
			fmt.Sprintf("%s Cores: %d", nerdfonts.CPUCores, metrics.Cores),
			fmt.Sprintf("%s Memory: %.1f%%", nerdfonts.CPUMemory, metrics.MemoryUsage),
			fmt.Sprintf("%s Load: %.2f", nerdfonts.CPULoad, metrics.LoadAvg),
			fmt.Sprintf("%s Governor: %s", nerdfonts.CPUGovernor, metrics.Governor),
			fmt.Sprintf("%s Boost: %t", nerdfonts.CPUBoost, metrics.BoostEnabled),
			fmt.Sprintf("%s Min/Max Freq: %.1f-%.1fGHz", nerdfonts.CPUMinMax, metrics.MinFreq, metrics.MaxFreq),
			fmt.Sprintf("%s IO Wait: %.1f%%", nerdfonts.CPUIOwait, metrics.IOWait),
		}
		return text, strings.Join(tooltipLines, "\n")
	} else {
		text = fmt.Sprintf("%.1f%% %d°C %.1fGHz %d cores", metrics.Usage, metrics.Temperature, metrics.Frequency, metrics.Cores)
		
		tooltipLines := []string{
			fmt.Sprintf("Usage: %.1f%%", metrics.Usage),
			fmt.Sprintf("Temp: %d°C", metrics.Temperature),
			fmt.Sprintf("Freq: %.1fGHz", metrics.Frequency),
			fmt.Sprintf("Cores: %d", metrics.Cores),
			fmt.Sprintf("Memory: %.1f%%", metrics.MemoryUsage),
			fmt.Sprintf("Load: %.2f", metrics.LoadAvg),
			fmt.Sprintf("Governor: %s", metrics.Governor),
			fmt.Sprintf("Boost: %t", metrics.BoostEnabled),
			fmt.Sprintf("Min/Max Freq: %.1f-%.1fGHz", metrics.MinFreq, metrics.MaxFreq),
			fmt.Sprintf("IO Wait: %.1f%%", metrics.IOWait),
		}
		return text, strings.Join(tooltipLines, "\n")
	}
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

		text, tooltip := formatCPUWithSymbols(metrics)
		
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
			_, tooltip := formatCPUWithSymbols(metrics)
			
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
			_, tooltip := formatCPUWithSymbols(metrics)
			
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
			_, tooltip := formatCPUWithSymbols(metrics)
			
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
			_, tooltip := formatCPUWithSymbols(metrics)
			
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