// +build linux

package system

import (
	"strconv"
	"strings"

	"github.com/ake-persson/peekaboo/pkg/pb/v1/resources"
)

func getCPU(s *resources.System) error {
	o, err := readFile("/proc/cpuinfo")
	if err != nil {
		return err
	}

	cpuID := -1
	cpuIDs := make(map[int]bool)
	s.CpuCoresPerSocket = 0
	s.CpuLogicalCores = 0
	for _, line := range strings.Split(string(o), "\n") {
		vals := strings.Split(line, ":")
		if len(vals) < 1 {
			continue
		}

		v := strings.Trim(strings.Join(vals[1:], " "), " ")
		if s.CpuModel == "" && strings.HasPrefix(line, "model name") {
			s.CpuModel = v
		} else if s.CpuFlags == "" && strings.HasPrefix(line, "flags") {
			s.CpuFlags = v
		} else if s.CpuCoresPerSocket == 0 && strings.HasPrefix(line, "cpu cores") {
			i, err := strconv.ParseUint(v, 10, 32)
			if err != nil {
				return err
			}
			s.CpuCoresPerSocket = int32(i)
		} else if strings.HasPrefix(line, "processor") {
			s.CpuLogicalCores++
		} else if strings.HasPrefix(line, "physical id") {
			cpuID, err = strconv.Atoi(v)
			if err != nil {
				return err
			}
			cpuIDs[cpuID] = true
		}
	}

	s.CpuSockets = int32(len(cpuIDs))
	s.CpuPhysicalCores = s.CpuSockets * s.CpuCoresPerSocket
	s.CpuThreadsPerCore = s.CpuLogicalCores / s.CpuSockets / s.CpuCoresPerSocket

	return nil
}
