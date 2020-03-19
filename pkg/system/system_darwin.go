// +build darwin

package system

import (
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/resources"
)

func GetSystem() (*resources.System, error) {
	s := &resources.System{
		Manufacturer: "Apple Inc.",
		Kernel:       "Darwin",
		Os:           "Mac OS X",
	}

	s.Hostname, _ = os.Hostname()

	o, err := exec.Command("system_profiler", "SPHardwareDataType").Output()
	if err != nil {
		return nil, err
	}

	for _, l := range strings.Split(string(o), "\n") {
		a := strings.SplitN(l, ":", 2)
		if len(a) < 2 {
			continue
		}

		k := strings.TrimSpace(a[0])
		v := strings.TrimSpace(a[1])
		switch k {
		case "Model Name":
			s.Product = v
		case "Model Identifier":
			s.ProductVersion = v
		case "Serial Number (system)":
			s.SerialNumber = v
		case "Boot ROM Version":
			s.BootRomVersion = v
		case "SMC Version (system)":
			s.SmcVersion = v
		}
	}

	o, err = exec.Command("sysctl", "-a").Output()
	if err != nil {
		return nil, err
	}

	for _, l := range strings.Split(string(o), "\n") {
		a := strings.SplitN(l, ":", 2)
		if len(a) < 2 {
			continue
		}

		k := strings.TrimSpace(a[0])
		v := strings.TrimSpace(a[1])
		switch k {
		case "hw.memsize":
			i, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				return nil, err
			}
			s.MemoryGb = int32(i / 1024 / 1024 / 1024)
		case "machdep.cpu.core_count":
			i, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				return nil, err
			}
			s.CpuCoresPerSocket = int32(i)
		case "hw.physicalcpu_max":
			i, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				return nil, err
			}
			s.CpuPhysicalCores = int32(i)
		case "hw.logicalcpu_max":
			i, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				return nil, err
			}
			s.CpuLogicalCores = int32(i)
		case "machdep.cpu.brand_string":
			s.CpuModel = v
		case "machdep.cpu.features":
			s.CpuFlags = v
		case "kern.osproductversion":
			s.OsVersion = v
		case "kern.osversion":
			s.OsBuild = v
		case "kern.version":
			s.KernelVersion = v
		case "kern.osrelease":
			s.KernelRelease = v
		}
	}

	s.CpuSockets = s.CpuPhysicalCores / s.CpuCoresPerSocket
	s.CpuThreadsPerCore = s.CpuLogicalCores / s.CpuSockets / s.CpuCoresPerSocket

	return s, nil
}
