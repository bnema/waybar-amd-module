// Package nerdfonts provides icon constants for GPU and CPU metrics display
package nerdfonts

// Icon represents a nerd font icon string
type Icon string

const (
	// GPU Icons
	GPUPower   Icon = "󰾲" // Power consumption
	GPUTemp    Icon = "🌡" // Temperature
	GPUFreq    Icon = "󰓅" // Frequency/clock speed
	GPUUtil    Icon = "󰈸" // Usage/activity
	GPUMemory  Icon = "󰍛" // Memory
	GPUFan     Icon = "󰈐" // Fan speed
	GPUVoltage Icon = "⚡" // Voltage

	// CPU Icons
	CPUUsage    Icon = "󰘚" // CPU utilization
	CPUTemp     Icon = "🌡" // Temperature
	CPUFreq     Icon = "󰓅" // Frequency/clock speed
	CPUCores    Icon = "󰻠" // Cores
	CPUMemory   Icon = "󰍛" // Memory
	CPULoad     Icon = "󰊚" // Load
	CPUGovernor Icon = "⚙" // Settings/power management
	CPUBoost    Icon = "󱓞" // Boost/turbo mode
	CPUMinMax   Icon = "↕" // Range indicator
	CPUIOwait   Icon = "⏸" // Time/wait
	CPUPower    Icon = "󰾲" // Power consumption

	// AMD Pstate Icons
	CPUPstateStatus        Icon = "" // Pstate driver status
	CPUPstatePrefcore      Icon = "" // Prefcore setting
	CPUEnergyPerfPref      Icon = "⚡" // Energy performance preference
	CPUHighestPerf         Icon = "󰓅" // Highest performance
	CPULowestNonlinearFreq Icon = "" // Lowest nonlinear frequency
)

func (i Icon) String() string {
	return string(i)
}
