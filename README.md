# waybar-amd-module

AMD GPU and CPU metrics for Waybar - monitor power, temperature, frequency, utilization, and more.

## Installation

```bash
go install github.com/bnema/waybar-amd-module@latest
```

## Usage

### CPU Commands

```bash
# All CPU metrics
waybar-amd-module cpu all

# Individual metrics
waybar-amd-module cpu usage
waybar-amd-module cpu temp
waybar-amd-module cpu freq
waybar-amd-module cpu cores
```

### GPU Commands

```bash
# All GPU metrics
waybar-amd-module gpu all

# Individual metrics  
waybar-amd-module gpu power
waybar-amd-module gpu temp
waybar-amd-module gpu freq
waybar-amd-module gpu util
```

### Flags

- `--format json|text` - Output format (default: json)
- `--nerd-font` - Use nerd font symbols for enhanced display

## Example Output

### CPU All Metrics (JSON)
```json
{"text":"15.6% 42°C 1.6GHz 12","tooltip":"Usage: 15.6%\nTemp: 42°C\nFreq: 1.6GHz\nCores: 12\nMemory: 45.2%\nLoad: 0.82\nGovernor: performance\nBoost: true\nMin/Max Freq: 0.4-4.2GHz\nIO Wait: 2.1%","class":"custom-cpu"}
```

### CPU Individual Metric (JSON)
```json
{"text":"16.7%","tooltip":"Usage: 16.7%\nTemp: 42°C\nFreq: 1.9GHz\nCores: 12\nMemory: 45.2%\nLoad: 0.82\nGovernor: performance\nBoost: true\nMin/Max Freq: 0.4-4.2GHz\nIO Wait: 2.1%","class":"custom-cpu"}
```

### GPU All Metrics (JSON)
```json
{"text":"5.2W 41°C 0.4GHz 0%","tooltip":"Power: 5.2W\nTemp: 41°C\nFreq: 0.4GHz\nUtil: 0%\nMemory: 46.9%\nFan: 0 RPM\nVoltage: 0.91V\nJunction: 0°C\nMemory Temp: 0°C\nPower Cap: 0.0W","class":"custom-gpu"}
```

### GPU Individual Metric (JSON)
```json
{"text":"5.1W","tooltip":"Power: 5.1W\nTemp: 41°C\nFreq: 0.4GHz\nUtil: 0%\nMemory: 46.9%\nFan: 0 RPM\nVoltage: 0.91V\nJunction: 0°C\nMemory Temp: 0°C\nPower Cap: 0.0W","class":"custom-gpu"}
```

### Text Format
```
15.1%
```

## Waybar Configuration

```json
{
  "custom/cpu": {
    "exec": "waybar-amd-module cpu usage",
    "return-type": "json",
    "interval": 2
  },
  "custom/gpu": {
    "exec": "waybar-amd-module gpu power", 
    "return-type": "json",
    "interval": 2
  }
}
```

Add `--nerd-font` flag to commands for icon display if you have nerd fonts installed.
