package system

import (
	"fmt"

	"github.com/ake-persson/peekaboo/pkg/pb/v1/resources"
)

var Headers = []string{
	"address",
	"manufacturer",
	"product",
	"product_version",
	"serial_number",
	"bios_vendor",
	"bios_date",
	"bios_version",
	"boot_rom_version",
	"smc_version",
	"memory_gb",
	"cpu_cores_per_socket",
	"cpu_physical_cores",
	"cpu_logical_cores",
	"cpu_sockets",
	"cpu_threads_per_core",
	"cpu_model",
	"cpu_flags",
	"os",
	"os_version",
	"os_build",
	"os_descr",
	"kernel",
	"kernel_version",
	"kernel_release",
	"description",
	"site",
	"rack",
	"rack_position",
	"rack_height",
	"vm",
	"vm_server",
}

func StringSlice(s *resources.System) []string {
	return []string{
		s.Address,
		s.Manufacturer,
		s.Product,
		s.ProductVersion,
		s.SerialNumber,
		s.BiosVendor,
		s.BiosDate,
		s.BiosVersion,
		s.BootRomVersion,
		s.SmcVersion,
		fmt.Sprintf("%dG", s.MemoryGb),
		fmt.Sprint(s.CpuCoresPerSocket),
		fmt.Sprint(s.CpuPhysicalCores),
		fmt.Sprint(s.CpuLogicalCores),
		fmt.Sprint(s.CpuSockets),
		fmt.Sprint(s.CpuThreadsPerCore),
		s.CpuModel,
		s.CpuFlags,
		s.Os,
		s.OsVersion,
		s.OsBuild,
		s.OsDescr,
		s.Kernel,
		s.KernelVersion,
		s.KernelRelease,
		s.Description,
		s.Site,
		s.Rack,
		fmt.Sprint(s.RackPosition),
		fmt.Sprintf("%dRU", s.RackHeight),
		fmt.Sprint(s.Vm),
		s.VmServer,
	}
}
