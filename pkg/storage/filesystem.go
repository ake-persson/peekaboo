package storage

import (
	"strings"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/resources"
	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/services"
)

func ToStringTable(resp *services.FileSystemResponse) [][]string {
	t := make([][]string, 0)
	for _, f := range resp.Filesystems {
		r := make([]string, 15)
		r[0] = resp.Hostname
		r[1] = resp.Filesystem
		r[2] = resp.Type
		r[3] = fmt.Sprintf("%d KB", f.SizeKb)
		r[4] = fmt.Sprintf("%d KB", f.UsedKb)
		r[5] = fmt.Sprintf("%d KB", f.FreeKb)
		r[6] = fmt.Sprintf("%0.2f %", f.UsedPct)
		r[7] = fmt.Sprint(f.Inodes)
		r[8] = fmt.Sprint(f.InodesFree)
		r[9] = fmt.Sprint(f.InodesUsed)
		r[10] = fmt.Sprintf("%02.f %", f.InodesUsedPct)
		r[11] = f.MountedOn
		r[12] = strings.Join(f.Options, ", ")
		r[13] = fmt.Sprintf("%t", f.IsLocal)
		r[14] = fmt.Sprintf("%t", f.IsAutomounted)
		t = append(t, r)
	}
	return t
}
