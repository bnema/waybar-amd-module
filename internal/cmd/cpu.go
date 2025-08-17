// Package cmd provides CLI commands for monitoring AMD CPU metrics
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/bnema/waybar-amd-module/internal/cpu"
	"github.com/bnema/waybar-amd-module/internal/formatting"
	"github.com/bnema/waybar-amd-module/internal/nerdfonts"
)

const (
	jsonFormat = "json"
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
			func() string {
				if metrics.Power > 0 {
					return fmt.Sprintf("%s System Power: +%.1fW charging", nerdfonts.CPUPower, metrics.Power)
				}
				if metrics.Power < 0 {
					return fmt.Sprintf("%s System Power: %.1fW discharging", nerdfonts.CPUPower, -metrics.Power)
				}
				return fmt.Sprintf("%s System Power: %.1fW", nerdfonts.CPUPower, metrics.Power)
			}(),
		}
		return text, strings.Join(tooltipLines, "\n")
	}
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
		func() string {
			if metrics.Power > 0 {
				return fmt.Sprintf("System Power: +%.1fW charging", metrics.Power)
			}
			if metrics.Power < 0 {
				return fmt.Sprintf("System Power: %.1fW discharging", -metrics.Power)
			}
			return fmt.Sprintf("System Power: %.1fW", metrics.Power)
		}(),
	}
	return text, strings.Join(tooltipLines, "\n")
}

func formatCPUAllMetrics(metrics *cpu.Metrics) string {
	if nerdFontFlag {
		return fmt.Sprintf("%s %.1f%% %s %d°C %s %.1fGHz %s %d %s %.1f%% %s %.2f %s %s %s %t %s %.1f-%.1fGHz %s %.1f%% %s %.1fW",
			nerdfonts.CPUUsage, metrics.Usage,
			nerdfonts.CPUTemp, metrics.Temperature,
			nerdfonts.CPUFreq, metrics.Frequency,
			nerdfonts.CPUCores, metrics.Cores,
			nerdfonts.CPUMemory, metrics.MemoryUsage,
			nerdfonts.CPULoad, metrics.LoadAvg,
			nerdfonts.CPUGovernor, metrics.Governor,
			nerdfonts.CPUBoost, metrics.BoostEnabled,
			nerdfonts.CPUMinMax, metrics.MinFreq, metrics.MaxFreq,
			nerdfonts.CPUIOwait, metrics.IOWait,
			nerdfonts.CPUPower, metrics.Power)
	}
	return fmt.Sprintf("%.1f%% %d°C %.1fGHz %d cores %.1f%% memory %.2f load %s %t boost %.1f-%.1fGHz %.1f%% iowait %.1fW system",
		metrics.Usage, metrics.Temperature, metrics.Frequency, metrics.Cores,
		metrics.MemoryUsage, metrics.LoadAvg, metrics.Governor, metrics.BoostEnabled,
		metrics.MinFreq, metrics.MaxFreq, metrics.IOWait, metrics.Power)
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

func formatCPUMemory(memory float64) string {
	if nerdFontFlag {
		return fmt.Sprintf("%s %.1f%%", nerdfonts.CPUMemory, memory)
	}
	return fmt.Sprintf("%.1f%%", memory)
}

func formatCPULoad(load float64) string {
	if nerdFontFlag {
		return fmt.Sprintf("%s %.2f", nerdfonts.CPULoad, load)
	}
	return fmt.Sprintf("%.2f", load)
}

func formatCPUGovernor(governor string) string {
	if nerdFontFlag {
		return fmt.Sprintf("%s %s", nerdfonts.CPUGovernor, governor)
	}
	return governor
}

func formatCPUBoost(boost bool) string {
	status := "disabled"
	if boost {
		status = "enabled"
	}
	if nerdFontFlag {
		return fmt.Sprintf("%s %s", nerdfonts.CPUBoost, status)
	}
	return status
}

//nolint:unused // utility function for future frequency range formatting
func formatCPUMinMaxFreq(minFreq, maxFreq float64) string {
	if nerdFontFlag {
		return fmt.Sprintf("%s %.1f-%.1fGHz", nerdfonts.CPUMinMax, minFreq, maxFreq)
	}
	return fmt.Sprintf("%.1f-%.1fGHz", minFreq, maxFreq)
}

func formatCPUIOWait(iowait float64) string {
	if nerdFontFlag {
		return fmt.Sprintf("%s %.1f%%", nerdfonts.CPUIOwait, iowait)
	}
	return fmt.Sprintf("%.1f%%", iowait)
}

func formatCPUPower(power float64) string {
	var sign, action string
	if power > 0 {
		sign = "+"
		action = " charging"
	} else if power < 0 {
		sign = ""
		action = " discharging"
		power = -power // Make positive for display
	} else {
		sign = ""
		action = ""
	}
	
	if nerdFontFlag {
		return fmt.Sprintf("%s %s%.1fW%s", nerdfonts.CPUPower, sign, power, action)
	}
	return fmt.Sprintf("%s%.1fW%s", sign, power, action)
}

var cpuCmd = &cobra.Command{
	Use:   "cpu",
	Short: "AMD CPU monitoring commands",
	Long:  "Monitor AMD CPU usage, temperature, frequency, and core count",
}

var cpuAllCmd = &cobra.Command{
	Use:   "all",
	Short: "Get all CPU metrics",
	Run: func(_ *cobra.Command, _ []string) {
		if !formatting.ValidateNoTooltipFlag(noTooltipFlag, formatFlag) {
			return
		}

		metrics, err := cpu.GetAllMetrics()
		if err != nil {
			switch formatFlag {
			case jsonFormat:
				fmt.Println("{}")
			default:
				return
			}
			return
		}

		text, tooltip := formatCPUWithSymbols(metrics)
		
		switch formatFlag {
		case jsonFormat:
			formatting.FormatJSONOutput(text, tooltip, "custom-cpu", noTooltipFlag)
		default:
			fmt.Println(formatCPUAllMetrics(metrics))
		}
	},
}

var cpuUsageCmd = &cobra.Command{
	Use:   "usage",
	Short: "Get CPU usage percentage",
	Run: func(_ *cobra.Command, _ []string) {
		if !formatting.ValidateNoTooltipFlag(noTooltipFlag, formatFlag) {
			return
		}

		usage, err := cpu.GetUsage()
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
			metrics, err := cpu.GetAllMetrics()
			if err != nil {
				fmt.Println("{}")
				return
			}
			
			text := formatCPUUsage(usage)
			_, tooltip := formatCPUWithSymbols(metrics)
			
			formatting.FormatJSONOutput(text, tooltip, "custom-cpu", noTooltipFlag)
		default:
			fmt.Println(formatCPUUsage(usage))
		}
	},
}

var cpuTempCmd = &cobra.Command{
	Use:   "temp",
	Short: "Get CPU temperature",
	Run: func(_ *cobra.Command, _ []string) {
		if !formatting.ValidateNoTooltipFlag(noTooltipFlag, formatFlag) {
			return
		}
		temp, err := cpu.GetTemperature()
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
			metrics, err := cpu.GetAllMetrics()
			if err != nil {
				fmt.Println("{}")
				return
			}
			
			text := formatCPUTemp(temp)
			_, tooltip := formatCPUWithSymbols(metrics)
			
			formatting.FormatJSONOutput(text, tooltip, "custom-cpu", noTooltipFlag)
		default:
			fmt.Println(formatCPUTemp(temp))
		}
	},
}

var cpuFreqCmd = &cobra.Command{
	Use:   "freq",
	Short: "Get CPU frequency",
	Run: func(_ *cobra.Command, _ []string) {
		if !formatting.ValidateNoTooltipFlag(noTooltipFlag, formatFlag) {
			return
		}
		freq, err := cpu.GetFrequency()
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
			metrics, err := cpu.GetAllMetrics()
			if err != nil {
				fmt.Println("{}")
				return
			}
			
			text := formatCPUFreq(freq)
			_, tooltip := formatCPUWithSymbols(metrics)
			
			formatting.FormatJSONOutput(text, tooltip, "custom-cpu", noTooltipFlag)
		default:
			fmt.Println(formatCPUFreq(freq))
		}
	},
}

var cpuCoresCmd = &cobra.Command{
	Use:   "cores",
	Short: "Get CPU core count",
	Run: func(_ *cobra.Command, _ []string) {
		if !formatting.ValidateNoTooltipFlag(noTooltipFlag, formatFlag) {
			return
		}

		cores, err := cpu.GetCores()
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
			metrics, err := cpu.GetAllMetrics()
			if err != nil {
				fmt.Println("{}")
				return
			}
			
			text := formatCPUCores(cores)
			_, tooltip := formatCPUWithSymbols(metrics)
			
			formatting.FormatJSONOutput(text, tooltip, "custom-cpu", noTooltipFlag)
		default:
			fmt.Println(formatCPUCores(cores))
		}
	},
}

var cpuMemoryCmd = &cobra.Command{
	Use:   "memory",
	Short: "Get system memory usage percentage",
	Run: func(_ *cobra.Command, _ []string) {
		if !formatting.ValidateNoTooltipFlag(noTooltipFlag, formatFlag) {
			return
		}

		memory, err := cpu.GetMemoryUsage()
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
			metrics, err := cpu.GetAllMetrics()
			if err != nil {
				fmt.Println("{}")
				return
			}
			
			text := formatCPUMemory(memory)
			_, tooltip := formatCPUWithSymbols(metrics)
			
			formatting.FormatJSONOutput(text, tooltip, "custom-cpu", noTooltipFlag)
		default:
			fmt.Println(formatCPUMemory(memory))
		}
	},
}

var cpuLoadCmd = &cobra.Command{
	Use:   "load",
	Short: "Get 1-minute load average",
	Run: func(_ *cobra.Command, _ []string) {
		if !formatting.ValidateNoTooltipFlag(noTooltipFlag, formatFlag) {
			return
		}

		load, err := cpu.GetLoadAverage()
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
			metrics, err := cpu.GetAllMetrics()
			if err != nil {
				fmt.Println("{}")
				return
			}
			
			text := formatCPULoad(load)
			_, tooltip := formatCPUWithSymbols(metrics)
			
			formatting.FormatJSONOutput(text, tooltip, "custom-cpu", noTooltipFlag)
		default:
			fmt.Println(formatCPULoad(load))
		}
	},
}

var cpuGovernorCmd = &cobra.Command{
	Use:   "governor",
	Short: "Get CPU frequency scaling governor",
	Run: func(_ *cobra.Command, _ []string) {
		if !formatting.ValidateNoTooltipFlag(noTooltipFlag, formatFlag) {
			return
		}

		governor, err := cpu.GetGovernor()
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
			metrics, err := cpu.GetAllMetrics()
			if err != nil {
				fmt.Println("{}")
				return
			}
			
			text := formatCPUGovernor(governor)
			_, tooltip := formatCPUWithSymbols(metrics)
			
			formatting.FormatJSONOutput(text, tooltip, "custom-cpu", noTooltipFlag)
		default:
			fmt.Println(formatCPUGovernor(governor))
		}
	},
}

var cpuBoostCmd = &cobra.Command{
	Use:   "boost",
	Short: "Get CPU boost enabled status",
	Run: func(_ *cobra.Command, _ []string) {
		if !formatting.ValidateNoTooltipFlag(noTooltipFlag, formatFlag) {
			return
		}

		boost, err := cpu.GetBoostEnabled()
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
			metrics, err := cpu.GetAllMetrics()
			if err != nil {
				fmt.Println("{}")
				return
			}
			
			text := formatCPUBoost(boost)
			_, tooltip := formatCPUWithSymbols(metrics)
			
			formatting.FormatJSONOutput(text, tooltip, "custom-cpu", noTooltipFlag)
		default:
			fmt.Println(formatCPUBoost(boost))
		}
	},
}

var cpuMinFreqCmd = &cobra.Command{
	Use:   "minfreq",
	Short: "Get minimum CPU frequency",
	Run: func(_ *cobra.Command, _ []string) {
		if !formatting.ValidateNoTooltipFlag(noTooltipFlag, formatFlag) {
			return
		}

		minFreq, _, err := cpu.GetMinMaxFreq()
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
			metrics, err := cpu.GetAllMetrics()
			if err != nil {
				fmt.Println("{}")
				return
			}
			
			text := formatCPUFreq(minFreq)
			_, tooltip := formatCPUWithSymbols(metrics)
			
			formatting.FormatJSONOutput(text, tooltip, "custom-cpu", noTooltipFlag)
		default:
			fmt.Println(formatCPUFreq(minFreq))
		}
	},
}

var cpuMaxFreqCmd = &cobra.Command{
	Use:   "maxfreq",
	Short: "Get maximum CPU frequency",
	Run: func(_ *cobra.Command, _ []string) {
		if !formatting.ValidateNoTooltipFlag(noTooltipFlag, formatFlag) {
			return
		}

		_, maxFreq, err := cpu.GetMinMaxFreq()
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
			metrics, err := cpu.GetAllMetrics()
			if err != nil {
				fmt.Println("{}")
				return
			}
			
			text := formatCPUFreq(maxFreq)
			_, tooltip := formatCPUWithSymbols(metrics)
			
			formatting.FormatJSONOutput(text, tooltip, "custom-cpu", noTooltipFlag)
		default:
			fmt.Println(formatCPUFreq(maxFreq))
		}
	},
}

var cpuIOWaitCmd = &cobra.Command{
	Use:   "iowait",
	Short: "Get I/O wait percentage",
	Run: func(_ *cobra.Command, _ []string) {
		if !formatting.ValidateNoTooltipFlag(noTooltipFlag, formatFlag) {
			return
		}

		iowait, err := cpu.GetIOWait()
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
			metrics, err := cpu.GetAllMetrics()
			if err != nil {
				fmt.Println("{}")
				return
			}
			
			text := formatCPUIOWait(iowait)
			_, tooltip := formatCPUWithSymbols(metrics)
			
			formatting.FormatJSONOutput(text, tooltip, "custom-cpu", noTooltipFlag)
		default:
			fmt.Println(formatCPUIOWait(iowait))
		}
	},
}

var cpuPowerCmd = &cobra.Command{
	Use:   "power",
	Short: "Get overall system power consumption",
	Run: func(_ *cobra.Command, _ []string) {
		if !formatting.ValidateNoTooltipFlag(noTooltipFlag, formatFlag) {
			return
		}

		power, err := cpu.GetPower()
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
			metrics, err := cpu.GetAllMetrics()
			if err != nil {
				fmt.Println("{}")
				return
			}
			
			text := formatCPUPower(power)
			_, tooltip := formatCPUWithSymbols(metrics)
			
			formatting.FormatJSONOutput(text, tooltip, "custom-cpu", noTooltipFlag)
		default:
			fmt.Println(formatCPUPower(power))
		}
	},
}

func init() {
	cpuCmd.AddCommand(cpuAllCmd)
	cpuCmd.AddCommand(cpuUsageCmd)
	cpuCmd.AddCommand(cpuTempCmd)
	cpuCmd.AddCommand(cpuFreqCmd)
	cpuCmd.AddCommand(cpuCoresCmd)
	cpuCmd.AddCommand(cpuMemoryCmd)
	cpuCmd.AddCommand(cpuLoadCmd)
	cpuCmd.AddCommand(cpuGovernorCmd)
	cpuCmd.AddCommand(cpuBoostCmd)
	cpuCmd.AddCommand(cpuMinFreqCmd)
	cpuCmd.AddCommand(cpuMaxFreqCmd)
	cpuCmd.AddCommand(cpuIOWaitCmd)
	cpuCmd.AddCommand(cpuPowerCmd)
}