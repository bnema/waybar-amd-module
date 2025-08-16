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
waybar-amd-module cpu usage      # CPU usage percentage
waybar-amd-module cpu temp       # CPU temperature
waybar-amd-module cpu freq       # CPU frequency
waybar-amd-module cpu cores      # CPU core count
waybar-amd-module cpu memory     # System memory usage
waybar-amd-module cpu load       # 1-minute load average
waybar-amd-module cpu governor   # CPU frequency governor
waybar-amd-module cpu boost      # CPU boost status
waybar-amd-module cpu minfreq    # Minimum CPU frequency
waybar-amd-module cpu maxfreq    # Maximum CPU frequency
waybar-amd-module cpu iowait     # I/O wait percentage
waybar-amd-module cpu power      # System power consumption (charging/discharging)
```

### GPU Commands

```bash
# All GPU metrics
waybar-amd-module gpu all

# Individual metrics  
waybar-amd-module gpu power      # GPU power consumption
waybar-amd-module gpu temp       # GPU temperature
waybar-amd-module gpu freq       # GPU frequency
waybar-amd-module gpu util       # GPU utilization
waybar-amd-module gpu memory     # VRAM usage percentage
waybar-amd-module gpu fan        # GPU fan speed in RPM
waybar-amd-module gpu voltage    # GPU voltage
waybar-amd-module gpu junction   # GPU junction temperature
waybar-amd-module gpu memtemp    # GPU memory temperature
waybar-amd-module gpu powercap   # GPU power cap limit
```

### Flags

- `--format json|text` - Output format (default: json)
- `--nerd-font` - Use nerd font symbols for enhanced display

## Example Output

### CPU All Metrics (JSON)
```json
{"text":"15.6% 42°C 1.6GHz 12","tooltip":"Usage: 15.6%\nTemp: 42°C\nFreq: 1.6GHz\nCores: 12\nMemory: 45.2%\nLoad: 0.82\nGovernor: performance\nBoost: true\nMin/Max Freq: 0.4-4.2GHz\nIO Wait: 2.1%\nSystem Power: +67.3W charging","class":"custom-cpu"}
```

### CPU Individual Metric (JSON)
```json
{"text":"16.7%","tooltip":"Usage: 16.7%\nTemp: 42°C\nFreq: 1.9GHz\nCores: 12\nMemory: 45.2%\nLoad: 0.82\nGovernor: performance\nBoost: true\nMin/Max Freq: 0.4-4.2GHz\nIO Wait: 2.1%\nSystem Power: +67.3W charging","class":"custom-cpu"}
```

### CPU Power Consumption (JSON)
```json
{"text":"+67.3W charging","tooltip":"Usage: 13.3%\nTemp: 48°C\nFreq: 1.7GHz\nCores: 12\nMemory: 41.8%\nLoad: 0.58\nGovernor: powersave\nBoost: false\nMin/Max Freq: 0.4-3.3GHz\nIO Wait: 0.0%\nSystem Power: +67.3W charging","class":"custom-cpu"}
```

### GPU All Metrics (JSON)
```json
{"text":"5.2W 41°C 0.4GHz 0%","tooltip":"Power: 5.2W\nTemp: 41°C\nFreq: 0.4GHz\nUtil: 0%\nMemory: 46.9%\nFan: 0 RPM\nVoltage: 0.91V\nJunction: 0°C\nMemory Temp: 0°C\nPower Cap: 0.0W","class":"custom-gpu"}
```

### GPU Individual Metric (JSON)
```json
{"text":"5.1W","tooltip":"Power: 5.1W\nTemp: 41°C\nFreq: 0.4GHz\nUtil: 0%\nMemory: 46.9%\nFan: 0 RPM\nVoltage: 0.91V\nJunction: 0°C\nMemory Temp: 0°C\nPower Cap: 0.0W","class":"custom-gpu"}
```

### Text Format Examples
```bash
# CPU usage
15.1%

# CPU all metrics
14.0% 47°C 1.6GHz 12 cores 41.7% memory 0.55 load powersave false boost 0.4-3.3GHz 0.0% iowait 65.7W system

# System power consumption
+67.3W charging
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
  },
  "custom/battery-power": {
    "exec": "waybar-amd-module cpu power",
    "return-type": "json",
    "interval": 2
  },
  "custom/cpu-all": {
    "exec": "waybar-amd-module cpu all",
    "return-type": "json",
    "interval": 2
  }
}
```

Add `--nerd-font` flag to commands for icon display if you have nerd fonts installed.
