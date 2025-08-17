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

### Hardware Discovery

```bash
# Scan for AMD hardware and update cache
waybar-amd-module scan
```

### Flags

- `--format json|text` - Output format (default: json)
- `--nerd-font` - Use nerd font symbols for enhanced display
- `--no-tooltip` - Remove tooltip field from JSON output (only works with `--format=json`)

## Example Output

### CPU All Metrics (JSON)
```json
{"text":"15.6% 42°C 1.6GHz 12 cores 45.2% memory 0.82 load performance true boost 0.4-4.2GHz 2.1% iowait +67.3W system","tooltip":"Usage: 15.6%\nTemp: 42°C\nFreq: 1.6GHz\nCores: 12\nMemory: 45.2%\nLoad: 0.82\nGovernor: performance\nBoost: true\nMin/Max Freq: 0.4-4.2GHz\nIO Wait: 2.1%\nSystem Power: +67.3W charging","class":"custom-cpu"}
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
{"text":"5.2W 41°C 0.4GHz 0% util 46.9% memory 0 RPM 0.91V 0°C junction 0°C memtemp 0.0W cap","tooltip":"Power: 5.2W\nTemp: 41°C\nFreq: 0.4GHz\nUtil: 0%\nMemory: 46.9%\nFan: 0 RPM\nVoltage: 0.91V\nJunction: 0°C\nMemory Temp: 0°C\nPower Cap: 0.0W","class":"custom-gpu"}
```

### GPU Individual Metric (JSON)
```json
{"text":"5.1W","tooltip":"Power: 5.1W\nTemp: 41°C\nFreq: 0.4GHz\nUtil: 0%\nMemory: 46.9%\nFan: 0 RPM\nVoltage: 0.91V\nJunction: 0°C\nMemory Temp: 0°C\nPower Cap: 0.0W","class":"custom-gpu"}
```

### JSON Output without Tooltip (`--no-tooltip`)
```json
{"class":"custom-cpu","text":"15.6% 42°C 1.6GHz 12 cores 45.2% memory 0.82 load performance true boost 0.4-4.2GHz 2.1% iowait +67.3W system"}
```

```json
{"class":"custom-gpu","text":"5.2W 41°C 0.4GHz 0% util 46.9% memory 0 RPM 0.91V 0°C junction 0°C memtemp 0.0W cap"}
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
  },
  "custom/cpu-minimal": {
    "exec": "waybar-amd-module cpu usage --no-tooltip",
    "return-type": "json",
    "interval": 2
  }
}
```

### Available Options

- Add `--nerd-font` flag for icon display if you have nerd fonts installed
- Add `--no-tooltip` flag to remove tooltips and reduce JSON size
- Use `--format=text` for simple text output without JSON wrapper

## Hardware Discovery & Caching

The module automatically discovers AMD hardware paths on first run and caches them for fast subsequent startups:

- **Cache Location**: `~/.cache/waybar-amd-module/paths.json` (XDG-compliant)
- **Auto-Discovery**: Scans for AMD GPUs and CPUs using multiple detection methods
- **Smart Validation**: Checks if cached paths still exist, rescans automatically if invalid
- **Universal Support**: Works on any Linux system with AMD hardware

### Cache Management

```bash
# Force rescan (useful after hardware changes)
waybar-amd-module scan

# View cached paths
cat ~/.cache/waybar-amd-module/paths.json
```

### Detection Methods

**GPU Discovery:**
- Primary: `/sys/class/drm/card*` with amdgpu driver detection
- Fallback: `/sys/bus/pci/drivers/amdgpu/*/hwmon/`
- Validates essential metric files exist

**CPU Discovery:**
- Detects AMD CPUs via `/proc/cpuinfo` (AuthenticAMD)
- Finds k10temp sensor in `/sys/class/hwmon/`
- Locates cpufreq paths and boost controls

**Power Discovery:**
- Battery: `/sys/class/power_supply/BAT*/`
- RAPL: `/sys/class/powercap/intel-rapl/` or `/sys/class/powercap/amd-rapl/`

The cache is automatically regenerated if:
- Cache file doesn't exist
- Cached hardware paths become invalid
- Manual `waybar-amd-module scan` command is run
