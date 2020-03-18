package system

import (
	"fmt"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/resources"
)

func SystemToStringTable(s *resources.System) ([]string, [][]string) {
	h := []string{"hostname", "address", "manufacturer", "product", "product_version", "serial_number",
		"bios_vendor", "bios_date", "bios_version", "boot_rom_version", "smc_version", "memory_gb",
		"cpu_cores_per_socket", "cpu_physical_cores", "cpu_logical_cores", "cpu_sockets", "cpu_threads_per_core",
		"cpu_model", "cpu_flags", "os", "os_version", "os_build", "os_descr", "kernel", "kernel_version",
		"kernel_release", "description", "site", "rack", "rack_position", "rack_size"}
	t := make([][]string, 0)
	r := make([]string, 31)
	r[0] = s.Hostname
	r[1] = s.Address
	r[2] = s.Manufacturer
	r[3] = s.Product
	r[4] = s.ProductVersion
	r[5] = s.SerialNumber
	r[6] = s.BiosVendor
	r[7] = s.BiosDate
	r[8] = s.BiosVersion
	r[9] = s.BootRomVersion
	r[10] = s.SmcVersion
	r[11] = fmt.Sprint(s.MemoryGb)
	r[12] = fmt.Sprint(s.CpuCoresPerSocket)
	r[13] = fmt.Sprint(s.CpuPhysicalCores)
	r[14] = fmt.Sprint(s.CpuLogicalCores)
	r[15] = fmt.Sprint(s.CpuSockets)
	r[16] = fmt.Sprint(s.CpuThreadsPerCore)
	r[17] = s.CpuModel
	r[18] = s.CpuFlags
	r[19] = s.Os
	r[20] = s.OsVersion
	r[21] = s.OsBuild
	r[22] = s.OsDescr
	r[23] = s.Kernel
	r[24] = s.KernelVersion
	r[25] = s.KernelRelease
	r[26] = s.Description
	r[27] = s.Site
	r[28] = s.Rack
	r[29] = fmt.Sprint(s.RackPosition)
	r[30] = fmt.Sprint(s.RackSize)
	t = append(t, r)
	return h, t
}
