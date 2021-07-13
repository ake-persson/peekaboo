// +build linux

package system

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/ake-persson/peekaboo/pkg/pb/v1/resources"
)

var (
	Site    string
	Rack    string
	RackPos int32
	RackHgt int32
	VM      bool
	VMSrvr  string
)

func SetInfo(site string, rack string, rackPos int, rackHgt int, vm bool, vmSrvr string) {
	Site = site
	Rack = rack
	RackPos = int32(rackPos)
	RackHgt = int32(rackHgt)
	VM = vm
	VMSrvr = vmSrvr
}

func readFile(fn string) (string, error) {
	b, err := ioutil.ReadFile(fn)
	return strings.TrimSpace(strings.TrimRight(string(b), "\n")), err
}

func GetSystem() (*resources.System, error) {
	s := &resources.System{
		Kernel:       "Linux",
		Site:         Site,
		Rack:         Rack,
		RackPosition: RackPos,
		RackHeight:   RackHgt,
		Vm:           VM,
		VmServer:     VMSrvr,
	}

	s.Hostname, _ = os.Hostname()

	var err error
	s.Manufacturer, err = readFile("/sys/devices/virtual/dmi/id/sys_vendor")
	if err != nil {
		return nil, err
	}

	s.Product, err = readFile("/sys/devices/virtual/dmi/id/product_name")
	if err != nil {
		return nil, err
	}

	s.ProductVersion, err = readFile("/sys/devices/virtual/dmi/id/product_version")
	if err != nil {
		return nil, err
	}

	s.SerialNumber, err = readFile("/sys/devices/virtual/dmi/id/product_serial")
	if err != nil {
		return nil, err
	}

	s.BiosVendor, err = readFile("/sys/devices/virtual/dmi/id/bios_vendor")
	if err != nil {
		return nil, err
	}

	s.BiosDate, err = readFile("/sys/devices/virtual/dmi/id/bios_date")
	if err != nil {
		return nil, err
	}

	s.BiosVersion, err = readFile("/sys/devices/virtual/dmi/id/bios_version")
	if err != nil {
		return nil, err
	}

	s.KernelVersion, err = readFile("/proc/sys/kernel/version")
	if err != nil {
		return nil, err
	}

	s.KernelRelease, err = readFile("/proc/sys/kernel/osrelease")
	if err != nil {
		return nil, err
	}

	o, err := readFile("/proc/meminfo")
	if err != nil {
		return nil, err
	}

	for _, l := range strings.Split(o, "\n") {
		a := strings.SplitN(l, ":", 2)
		if len(a) < 2 {
			continue
		}

		k := strings.TrimSpace(a[0])
		v := strings.TrimSpace(a[1])
		v2 := strings.TrimSpace(strings.Split(v, " ")[0])

		switch k {
		case "MemTotal":
			i, err := strconv.ParseUint(v2, 10, 64)
			if err != nil {
				return nil, err
			}
			s.MemoryGb = int32(i / 1024 / 1024)
		}
	}

	if err := getCPU(s); err != nil {
		return nil, err
	}

	if err := getOs(s); err != nil {
		return nil, err
	}

	return s, nil
}
