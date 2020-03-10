// +build linux

package system

import (
	"os/exec"
	"strings"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/resources"
)

func getOs(s *resources.System) error {
	out, err := exec.Command("lsb_release", "-a").Output()
	if err != nil {
		return err
	}

	for _, l := range strings.Split(string(out), "\n") {
		kv := strings.SplitN(l, ":\t", 2)

		if len(kv) < 2 {
			continue
		}

		switch kv[0] {
		case "Distributor ID":
			s.Os = kv[1]
		case "Description":
			s.OsDescr = kv[1]
		case "Release":
			s.OsVersion = kv[1]
		case "Codename":
			s.OsBuild = kv[1]
		}
	}

	return nil
}
