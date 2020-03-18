package system

import (
	"fmt"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/resources"
	"github.com/peekaboo-labs/peekaboo/pkg/text"
)

func ToTable(s *resources.System) *text.Table {
	t := text.Table{
		Headers: []string{"hostname", "address", "manufacturer", "product", "product_version", "serial_number",
			"bios_vendor", "bios_date", "bios_version", "boot_rom_version", "smc_version", "memory_gb",
			"cpu_cores_per_socket", "cpu_physical_cores", "cpu_logical_cores", "cpu_sockets", "cpu_threads_per_core",
			"cpu_model", "cpu_flags", "os", "os_version", "os_build", "os_descr", "kernel", "kernel_version",
			"kernel_release", "description", "site", "rack", "rack_position", "rack_size"},
		Rows: make([][]string, 1),
	}

	t.Rows[0] = make([]string, 31)
	t.Rows[0][0] = s.Hostname
	t.Rows[0][1] = s.Address
	t.Rows[0][2] = s.Manufacturer
	t.Rows[0][3] = s.Product
	t.Rows[0][4] = s.ProductVersion
	t.Rows[0][5] = s.SerialNumber
	t.Rows[0][6] = s.BiosVendor
	t.Rows[0][7] = s.BiosDate
	t.Rows[0][8] = s.BiosVersion
	t.Rows[0][9] = s.BootRomVersion
	t.Rows[0][10] = s.SmcVersion
	t.Rows[0][11] = fmt.Sprint(s.MemoryGb)
	t.Rows[0][12] = fmt.Sprint(s.CpuCoresPerSocket)
	t.Rows[0][13] = fmt.Sprint(s.CpuPhysicalCores)
	t.Rows[0][14] = fmt.Sprint(s.CpuLogicalCores)
	t.Rows[0][15] = fmt.Sprint(s.CpuSockets)
	t.Rows[0][16] = fmt.Sprint(s.CpuThreadsPerCore)
	t.Rows[0][17] = s.CpuModel
	t.Rows[0][18] = s.CpuFlags
	t.Rows[0][19] = s.Os
	t.Rows[0][20] = s.OsVersion
	t.Rows[0][21] = s.OsBuild
	t.Rows[0][22] = s.OsDescr
	t.Rows[0][23] = s.Kernel
	t.Rows[0][24] = s.KernelVersion
	t.Rows[0][25] = s.KernelRelease
	t.Rows[0][26] = s.Description
	t.Rows[0][27] = s.Site
	t.Rows[0][28] = s.Rack
	t.Rows[0][29] = fmt.Sprint(s.RackPosition)
	t.Rows[0][30] = fmt.Sprint(s.RackSize)

	return &t
}
