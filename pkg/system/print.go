package system

import (
	"fmt"

	"github.com/mickep76/color"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/resources"
)

func PrintSystem(s *resources.System) {
	f := fmt.Sprintf("\t%s%%-24s%s : %s%%v%s", color.Cyan, color.Reset, color.Yellow, color.Reset)
	fmt.Printf(f, "Hostname", s.Hostname)

	f = "\n" + f
	fmt.Printf(f, "Manufacturer", s.Manufacturer)
	fmt.Printf(f, "Product", s.Product)
	fmt.Printf(f, "Product Version", s.ProductVersion)
	fmt.Printf(f, "Serial Number", s.SerialNumber)

	switch s.Kernel {
	case "Linux":
		fmt.Printf(f, "BIOS Vendor", s.BiosVendor)
		fmt.Printf(f, "BIOS Date", s.BiosDate)
		fmt.Printf(f, "BIOS Version", s.BiosVersion)
	case "Darwin":
		fmt.Printf(f, "Boot ROM Version", s.BootRomVersion)
		fmt.Printf(f, "SMC Version", s.SmcVersion)
	}

	fmt.Printf("%s %sGB%s", fmt.Sprintf(f, "Memory", s.MemoryGb), color.Cyan, color.Reset)
	fmt.Printf(f, "CPU Model", s.CpuModel)
	fmt.Printf(f, "CPU Flags", s.CpuFlags)
	fmt.Printf(f, "CPU Cores Per Socket", s.CpuCoresPerSocket)
	fmt.Printf(f, "CPU Physical Cores", s.CpuPhysicalCores)
	fmt.Printf(f, "CPU Logical Cores", s.CpuLogicalCores)
	fmt.Printf(f, "CPU Sockets", s.CpuSockets)
	fmt.Printf(f, "CPU Threads Per Core", s.CpuThreadsPerCore)
	fmt.Printf(f, "Operating System", s.Os)
	fmt.Printf(f, "Operating System Version", s.OsVersion)
	fmt.Printf(f, "Operating System Build", s.OsBuild)
	fmt.Printf(f, "Kernel", s.Kernel)
	fmt.Printf(f, "Kernel Version", s.KernelVersion)
	fmt.Printf(f+"\n", "Kernel Release", s.KernelRelease)
}
