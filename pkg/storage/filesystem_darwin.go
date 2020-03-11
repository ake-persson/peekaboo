// +build darwin

package storage

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/resources"
	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/services"
)

func ListFilesystems() (*services.ListFilesystemsResponse, error) {
	out, err := exec.Command("df", "-k").Output()
	if err != nil {
		return nil, err
	}

	resp := &services.ListFilesystemsResponse{Filesystems: []*resources.Filesystem{}}
	for i, l := range strings.Split(string(out), "\n") {
		if i < 1 || strings.HasPrefix(l, "map") {
			continue
		}

		a := strings.Fields(l)
		if len(a) < 6 {
			continue
		}

		f := &resources.Filesystem{
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

		f.UsedPct = float32(f.UsedKb) / float32(f.SizeKb) * 100

		if f.InodesUsed, err = strconv.ParseUint(a[5], 10, 64); err != nil {
			return nil, err
		}

		if f.InodesFree, err = strconv.ParseUint(a[6], 10, 64); err != nil {
			return nil, err
		}

		f.Inodes = f.InodesUsed + f.InodesFree
		f.InodesUsedPct = float32(f.InodesUsed) / float32(f.Inodes) * 100

		resp.Filesystems = append(resp.Filesystems, f)
	}

	return resp, nil
}
