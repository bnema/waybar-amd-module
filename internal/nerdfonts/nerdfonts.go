package nerdfonts

type Icon string

const (
	// GPU Icons
	GPUPower Icon = "󰢮"
	GPUTemp  Icon = "󰔏"
	GPUFreq  Icon = "󰾆"
	GPUUtil  Icon = "󰾅"
	
	// CPU Icons
	CPUUsage Icon = "󰍛"
	CPUTemp  Icon = "󰔏"
	CPUFreq  Icon = "󰓅"
	CPUCores Icon = "󰻠"
)

func (i Icon) String() string {
	return string(i)
}