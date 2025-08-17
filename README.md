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
waybar-amd-module cpu pstate-status # AMD pstate driver status
waybar-amd-module cpu energy-perf   # Energy performance preference
waybar-amd-module cpu pstate        # All AMD pstate information
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
- `--with-pstate` - Include AMD pstate information in CPU metrics tooltips


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
  },
  "custom/cpu-pstate": {
    "exec": "waybar-amd-module cpu pstate-status --nerd-font",
    "return-type": "json",
    "interval": 5
  },
  "custom/cpu-energy-perf": {
    "exec": "waybar-amd-module cpu energy-perf --nerd-font",
    "return-type": "json",
    "interval": 10
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
- Discovers AMD pstate driver paths (`/sys/devices/system/cpu/amd_pstate/`)

**Power Discovery:**
- Battery: `/sys/class/power_supply/BAT*/`
- RAPL: `/sys/class/powercap/intel-rapl/` or `/sys/class/powercap/amd-rapl/`

The cache is automatically regenerated if:
- Cache file doesn't exist
- Cached hardware paths become invalid
- Manual `waybar-amd-module scan` command is run
