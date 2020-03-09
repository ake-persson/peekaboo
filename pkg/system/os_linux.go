// +build linux

package system

import (
	"fmt"
	"strings"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/resources"
)

func getOs(s *resources.System) error {
	descr, err := readFile("/etc/redhat-release")
	if err != nil {
		return err
	}

	s.OsDescr = strings.TrimSpace(string(descr))

	a := strings.SplitN(s.OsDescr, " release ", 2)
	if len(a) != 2 {
		return fmt.Errorf("unknown format in [%s]: %s", "/etc/redhat-release", s.OsDescr)
	}

	s.Os = strings.Replace(a[0], " Linux", "", 1)

	a = strings.SplitN(a[1], " ", 2)
	if len(a) != 2 {
		return fmt.Errorf("unknown format in [%s]: %s", "/etc/redhat-release", s.OsDescr)
	}

	s.OsVersion = a[0]
	s.OsBuild = a[1][1 : len(a[1])-1]

	return nil
}
