// Package cmd provides CLI commands for monitoring AMD GPU metrics
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
	}
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

func formatGPUAllMetrics(metrics *gpu.Metrics) string {
	if nerdFontFlag {
		return fmt.Sprintf("%s %.1fW %s %d°C %s %.1fGHz %s %d%% %s %.1f%% %s %d RPM %s %.2fV %s %d°C %s %d°C %s %.1fW",
			nerdfonts.GPUPower, metrics.Power,
			nerdfonts.GPUTemp, metrics.Temperature,
			nerdfonts.GPUFreq, metrics.Frequency,
			nerdfonts.GPUUtil, metrics.Utilization,
			nerdfonts.GPUMemory, metrics.MemoryUsage,
			nerdfonts.GPUFan, metrics.FanSpeed,
			nerdfonts.GPUVoltage, metrics.Voltage,
			nerdfonts.GPUTemp, metrics.JunctionTemp,
			nerdfonts.GPUTemp, metrics.MemoryTemp,
			nerdfonts.GPUPower, metrics.PowerCap)
	}
	return fmt.Sprintf("%.1fW %d°C %.1fGHz %d%% util %.1f%% memory %d RPM %.2fV %d°C junction %d°C memtemp %.1fW cap",
		metrics.Power, metrics.Temperature, metrics.Frequency, metrics.Utilization,
		metrics.MemoryUsage, metrics.FanSpeed, metrics.Voltage, metrics.JunctionTemp,
		metrics.MemoryTemp, metrics.PowerCap)
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

func formatMemory(memory float64) string {
	if nerdFontFlag {
		return fmt.Sprintf("%s %.1f%%", nerdfonts.GPUMemory, memory)
	}
	return fmt.Sprintf("%.1f%%", memory)
}

func formatFan(fan int) string {
	if nerdFontFlag {
		return fmt.Sprintf("%s %d RPM", nerdfonts.GPUFan, fan)
	}
	return fmt.Sprintf("%d RPM", fan)
}

func formatVoltage(voltage float64) string {
	if nerdFontFlag {
		return fmt.Sprintf("%s %.2fV", nerdfonts.GPUVoltage, voltage)
	}
	return fmt.Sprintf("%.2fV", voltage)
}

func formatJunctionTemp(temp int) string {
	if nerdFontFlag {
		return fmt.Sprintf("%s %d°C", nerdfonts.GPUTemp, temp)
	}
	return fmt.Sprintf("%d°C (junction)", temp)
}

func formatMemoryTemp(temp int) string {
	if nerdFontFlag {
		return fmt.Sprintf("%s %d°C", nerdfonts.GPUTemp, temp)
	}
	return fmt.Sprintf("%d°C (memory)", temp)
}

func formatPowerCap(cap float64) string {
	if nerdFontFlag {
		return fmt.Sprintf("%s %.1fW", nerdfonts.GPUPower, cap)
	}
	return fmt.Sprintf("%.1fW (cap)", cap)
}

var gpuCmd = &cobra.Command{
	Use:   "gpu",
	Short: "AMD GPU monitoring commands",
	Long:  "Monitor AMD GPU power, temperature, frequency, and utilization",
}

var gpuAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Get all GPU metrics",
	Run: func(_ *cobra.Command, _ []string) {
		metrics, err := gpu.GetAllMetrics()
		if err != nil {
			switch formatFlag {
			case jsonFormat:
				fmt.Println("{}")
			default:
				return
			}
			return
		}

		text, tooltip := formatWithSymbols(metrics)
		
		switch formatFlag {
		case jsonFormat:
			output := gpu.WaybarOutput{
				Text:    text,
				Tooltip: tooltip,
				Class:   "custom-gpu",
			}
			jsonData, _ := json.Marshal(output)
			fmt.Println(string(jsonData))
		default:
			fmt.Println(formatGPUAllMetrics(metrics))
		}
	},
}

var gpuPowerCmd = &cobra.Command{
	Use:   "power",
	Short: "Get GPU power consumption",
	Run: func(_ *cobra.Command, _ []string) {
		power, err := gpu.GetPower()
		if err != nil {
			switch formatFlag {
			case jsonFormat:
				fmt.Println("{}")
			default:
				return
			}
			return
		}

		switch formatFlag {
		case jsonFormat:
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
	Run: func(_ *cobra.Command, _ []string) {
		temp, err := gpu.GetTemperature()
		if err != nil {
			switch formatFlag {
			case jsonFormat:
				fmt.Println("{}")
			default:
				return
			}
			return
		}

		switch formatFlag {
		case jsonFormat:
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
	Run: func(_ *cobra.Command, _ []string) {
		freq, err := gpu.GetFrequency()
		if err != nil {
			switch formatFlag {
			case jsonFormat:
				fmt.Println("{}")
			default:
				return
			}
			return
		}

		switch formatFlag {
		case jsonFormat:
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
	Run: func(_ *cobra.Command, _ []string) {
		util, err := gpu.GetUtilization()
		if err != nil {
			switch formatFlag {
			case jsonFormat:
				fmt.Println("{}")
			default:
				return
			}
			return
		}

		switch formatFlag {
		case jsonFormat:
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

var gpuMemoryCmd = &cobra.Command{
	Use:   "memory",
	Short: "Get VRAM usage percentage",
	Run: func(_ *cobra.Command, _ []string) {
		memory, err := gpu.GetMemoryUsage()
		if err != nil {
			switch formatFlag {
			case jsonFormat:
				fmt.Println("{}")
			default:
				return
			}
			return
		}

		switch formatFlag {
		case jsonFormat:
			// Get all metrics for tooltip
			metrics, err := gpu.GetAllMetrics()
			if err != nil {
				fmt.Println("{}")
				return
			}
			
			text := formatMemory(memory)
			_, tooltip := formatWithSymbols(metrics)
			
			output := gpu.WaybarOutput{
				Text:    text,
				Tooltip: tooltip,
				Class:   "custom-gpu",
			}
			jsonData, _ := json.Marshal(output)
			fmt.Println(string(jsonData))
		default:
			fmt.Println(formatMemory(memory))
		}
	},
}

var gpuFanCmd = &cobra.Command{
	Use:   "fan",
	Short: "Get GPU fan speed in RPM",
	Run: func(_ *cobra.Command, _ []string) {
		fan, err := gpu.GetFanSpeed()
		if err != nil {
			switch formatFlag {
			case jsonFormat:
				fmt.Println("{}")
			default:
				return
			}
			return
		}

		switch formatFlag {
		case jsonFormat:
			// Get all metrics for tooltip
			metrics, err := gpu.GetAllMetrics()
			if err != nil {
				fmt.Println("{}")
				return
			}
			
			text := formatFan(fan)
			_, tooltip := formatWithSymbols(metrics)
			
			output := gpu.WaybarOutput{
				Text:    text,
				Tooltip: tooltip,
				Class:   "custom-gpu",
			}
			jsonData, _ := json.Marshal(output)
			fmt.Println(string(jsonData))
		default:
			fmt.Println(formatFan(fan))
		}
	},
}

var gpuVoltageCmd = &cobra.Command{
	Use:   "voltage",
	Short: "Get GPU voltage",
	Run: func(_ *cobra.Command, _ []string) {
		voltage, err := gpu.GetVoltage()
		if err != nil {
			switch formatFlag {
			case jsonFormat:
				fmt.Println("{}")
			default:
				return
			}
			return
		}

		switch formatFlag {
		case jsonFormat:
			// Get all metrics for tooltip
			metrics, err := gpu.GetAllMetrics()
			if err != nil {
				fmt.Println("{}")
				return
			}
			
			text := formatVoltage(voltage)
			_, tooltip := formatWithSymbols(metrics)
			
			output := gpu.WaybarOutput{
				Text:    text,
				Tooltip: tooltip,
				Class:   "custom-gpu",
			}
			jsonData, _ := json.Marshal(output)
			fmt.Println(string(jsonData))
		default:
			fmt.Println(formatVoltage(voltage))
		}
	},
}

var gpuJunctionCmd = &cobra.Command{
	Use:   "junction",
	Short: "Get GPU junction temperature",
	Run: func(_ *cobra.Command, _ []string) {
		junctionTemp, err := gpu.GetJunctionTemp()
		if err != nil {
			switch formatFlag {
			case jsonFormat:
				fmt.Println("{}")
			default:
				return
			}
			return
		}

		switch formatFlag {
		case jsonFormat:
			// Get all metrics for tooltip
			metrics, err := gpu.GetAllMetrics()
			if err != nil {
				fmt.Println("{}")
				return
			}
			
			text := formatJunctionTemp(junctionTemp)
			_, tooltip := formatWithSymbols(metrics)
			
			output := gpu.WaybarOutput{
				Text:    text,
				Tooltip: tooltip,
				Class:   "custom-gpu",
			}
			jsonData, _ := json.Marshal(output)
			fmt.Println(string(jsonData))
		default:
			fmt.Println(formatJunctionTemp(junctionTemp))
		}
	},
}

var gpuMemTempCmd = &cobra.Command{
	Use:   "memtemp",
	Short: "Get GPU memory temperature",
	Run: func(_ *cobra.Command, _ []string) {
		memTemp, err := gpu.GetMemoryTemp()
		if err != nil {
			switch formatFlag {
			case jsonFormat:
				fmt.Println("{}")
			default:
				return
			}
			return
		}

		switch formatFlag {
		case jsonFormat:
			// Get all metrics for tooltip
			metrics, err := gpu.GetAllMetrics()
			if err != nil {
				fmt.Println("{}")
				return
			}
			
			text := formatMemoryTemp(memTemp)
			_, tooltip := formatWithSymbols(metrics)
			
			output := gpu.WaybarOutput{
				Text:    text,
				Tooltip: tooltip,
				Class:   "custom-gpu",
			}
			jsonData, _ := json.Marshal(output)
			fmt.Println(string(jsonData))
		default:
			fmt.Println(formatMemoryTemp(memTemp))
		}
	},
}

var gpuPowerCapCmd = &cobra.Command{
	Use:   "powercap",
	Short: "Get GPU power cap limit",
	Run: func(_ *cobra.Command, _ []string) {
		powerCap, err := gpu.GetPowerCap()
		if err != nil {
			switch formatFlag {
			case jsonFormat:
				fmt.Println("{}")
			default:
				return
			}
			return
		}

		switch formatFlag {
		case jsonFormat:
			// Get all metrics for tooltip
			metrics, err := gpu.GetAllMetrics()
			if err != nil {
				fmt.Println("{}")
				return
			}
			
			text := formatPowerCap(powerCap)
			_, tooltip := formatWithSymbols(metrics)
			
			output := gpu.WaybarOutput{
				Text:    text,
				Tooltip: tooltip,
				Class:   "custom-gpu",
			}
			jsonData, _ := json.Marshal(output)
			fmt.Println(string(jsonData))
		default:
			fmt.Println(formatPowerCap(powerCap))
		}
	},
}

func init() {
	gpuCmd.AddCommand(gpuAllCmd)
	gpuCmd.AddCommand(gpuPowerCmd)
	gpuCmd.AddCommand(gpuTempCmd)
	gpuCmd.AddCommand(gpuFreqCmd)
	gpuCmd.AddCommand(gpuUtilCmd)
	gpuCmd.AddCommand(gpuMemoryCmd)
	gpuCmd.AddCommand(gpuFanCmd)
	gpuCmd.AddCommand(gpuVoltageCmd)
	gpuCmd.AddCommand(gpuJunctionCmd)
	gpuCmd.AddCommand(gpuMemTempCmd)
	gpuCmd.AddCommand(gpuPowerCapCmd)
}