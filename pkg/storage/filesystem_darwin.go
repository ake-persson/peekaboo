// +build darwin

package storage

import (
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/resources"
	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/services"
)

var re = regexp.MustCompile("(.*) on .* \\((.*)\\)")

func ListFilesystems() (*services.ListFilesystemsResponse, error) {
	hostname, _ := os.Hostname()

	out, err := exec.Command("mount").Output()
	if err != nil {
		return nil, err
	}

	options := map[string][]string{}
	for _, l := range strings.Split(string(out), "\n") {
		a := re.FindStringSubmatch(l)
		if len(a) < 3 {
			continue
		}

		options[a[1]] = strings.Split(a[2], ", ")
	}

	out, err = exec.Command("df", "-k", "-i").Output()
	if err != nil {
		return nil, err
	}

	resp := &services.ListFilesystemsResponse{Filesystems: []*resources.Filesystem{}}
	for i, l := range strings.Split(string(out), "\n") {
		if i < 1 {
			continue
		}

		a := strings.Fields(l)
		if len(a) < 9 {
			continue
		}

		if a[0] == "map" {
			a = a[1:]
		}

		f := &resources.Filesystem{
			Hostname:   hostname,
			Filesystem: a[0],
			MountedOn:  a[8],
		}

		var err error
		if f.SizeKb, err = strconv.ParseUint(a[1], 10, 64); err != nil {
			return nil, err
		}

		if f.UsedKb, err = strconv.ParseUint(a[2], 10, 64); err != nil {
			return nil, err
		}

		if f.FreeKb, err = strconv.ParseUint(a[3], 10, 64); err != nil {
			return nil, err
		}

		if f.SizeKb != 0 {
			f.UsedPct = float32(f.UsedKb) / float32(f.SizeKb) * 100
		}

		if f.InodesUsed, err = strconv.ParseUint(a[5], 10, 64); err != nil {
			return nil, err
		}

		if f.InodesFree, err = strconv.ParseUint(a[6], 10, 64); err != nil {
			return nil, err
		}

		f.Inodes = f.InodesUsed + f.InodesFree

		if f.Inodes != 0 {
			f.InodesUsedPct = float32(f.InodesUsed) / float32(f.Inodes) * 100
		}

		if v, ok := options[f.Filesystem]; ok {
			if len(v) > 2 {
				f.Type = v[0]
				switch v[1] {
				case "local":
					f.IsLocal = true
				case "autofs":
					f.IsAutomounted = true
				}
				f.MountOptions = v[2:]
			}
		}

		resp.Filesystems = append(resp.Filesystems, f)
	}

	return resp, nil
}
