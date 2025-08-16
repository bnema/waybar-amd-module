// Package nerdfonts provides icon constants for GPU and CPU metrics display
package nerdfonts

// Icon represents a nerd font icon string
type Icon string

const (
	// GPU Icons
	GPUPower   Icon = "󰾲" // Power consumption (lightning bolt)
	GPUTemp    Icon = "🌡" // Temperature (thermometer)
	GPUFreq    Icon = "󰓅" // Frequency/clock speed (sine wave)
	GPUUtil    Icon = "󰈸" // Usage/activity (gauge/speedometer)
	GPUMemory  Icon = "󰍛" // Memory (memory chip)
	GPUFan     Icon = "󰈐" // Fan speed (fan blade)
	GPUVoltage Icon = "⚡" // Voltage (lightning bolt)

	// CPU Icons
	CPUUsage    Icon = "󰘚" // CPU utilization (processor chip)
	CPUTemp     Icon = "🌡" // Temperature (thermometer)
	CPUFreq     Icon = "󰓅" // Frequency/clock speed (sine wave)
	CPUCores    Icon = "󰻠" // Cores (multi-core processor)
	CPUMemory   Icon = "󰍛" // Memory (memory chip)
	CPULoad     Icon = "󰊚" // Load (weight/pressure)
	CPUGovernor Icon = "⚙" // Settings/power management (gear)
	CPUBoost    Icon = "󱓞" // Boost/turbo mode (rocket)
	CPUMinMax   Icon = "↕" // Range indicator (up-down arrow)
	CPUIOwait   Icon = "⏸" // Time/wait (pause symbol)
	CPUPower    Icon = "󰾲" // Power consumption (lightning bolt)
)

func (i Icon) String() string {
	return string(i)
}
