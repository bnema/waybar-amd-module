# waybar-amd-module

AMD GPU and CPU metrics for Waybar - monitor power, temperature, frequency, and utilization.

## Installation

```bash
go install github.com/bnema/waybar-amd-module@latest
```

## Usage

### CPU Commands

```bash
# All CPU metrics
waybar-amd-module cpu all --nerd-font

# Individual metrics
waybar-amd-module cpu usage --nerd-font
waybar-amd-module cpu temp --nerd-font
waybar-amd-module cpu freq --nerd-font
waybar-amd-module cpu cores --nerd-font
```

### GPU Commands

```bash
# All GPU metrics
waybar-amd-module gpu all --nerd-font

# Individual metrics  
waybar-amd-module gpu power --nerd-font
waybar-amd-module gpu temp --nerd-font
waybar-amd-module gpu freq --nerd-font
waybar-amd-module gpu util --nerd-font
```

### Flags

- `--format json|text` - Output format (default: json)
- `--nerd-font` - Use nerd font symbols

## Example Output

### CPU All Metrics (JSON)
```json
{"text":"󰍛 15.6% 󰔏 42°C 󰓅 1.6GHz 󰻠 12","tooltip":"󰍛 Usage: 15.6%\n󰔏 Temp: 42°C\n󰓅 Freq: 1.6GHz\n󰻠 Cores: 12","class":"custom-cpu"}
```

### CPU Individual Metric (JSON)
```json
{"text":"󰍛 16.7%","tooltip":"󰍛 Usage: 18.3%\n󰔏 Temp: 42°C\n󰓅 Freq: 1.9GHz\n󰻠 Cores: 12","class":"custom-cpu"}
```

### GPU All Metrics (JSON)
```json
{"text":"󰢮 5.2W 󰔏 41°C 󰾆 0.4GHz 󰾅 0%","tooltip":"󰢮 Power: 5.2W\n󰔏 Temp: 41°C\n󰾆 Freq: 0.4GHz\n󰾅 Util: 0%","class":"custom-gpu"}
```

### GPU Individual Metric (JSON)
```json
{"text":"󰢮 5.1W","tooltip":"󰢮 Power: 5.1W\n󰔏 Temp: 41°C\n󰾆 Freq: 0.4GHz\n󰾅 Util: 0%","class":"custom-gpu"}
```

### Text Format
```
󰍛 15.1%
```

## Waybar Configuration

```json
{
  "custom/cpu": {
    "exec": "waybar-amd-module cpu usage --nerd-font",
    "return-type": "json",
    "interval": 2
  },
  "custom/gpu": {
    "exec": "waybar-amd-module gpu power --nerd-font", 
    "return-type": "json",
    "interval": 2
  }
}
```
