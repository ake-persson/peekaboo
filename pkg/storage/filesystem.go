package storage

import (
	"fmt"
	"strings"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/resources"
)

func FilesystemsToStringTable(hostname string, fs []*resources.Filesystem) ([]string, [][]string) {
	h := []string{"hostname", "filesystem", "type", "size_kb", "used_kb", "free_kb", "used_pct", "inodes",
		"inodes_free", "inodes_used", "inodes_used_pct", "mounted_on", "mount_options", "is_local", "is_automounted"}
	t := make([][]string, 0)
	for _, f := range fs {
		r := make([]string, 15)
		r[0] = hostname
		r[1] = f.Filesystem
		r[2] = f.Type
		r[3] = fmt.Sprintf("%dK", f.SizeKb)
		r[4] = fmt.Sprintf("%dK", f.UsedKb)
		r[5] = fmt.Sprintf("%dK", f.FreeKb)
		r[6] = fmt.Sprintf("%.2f%%", f.UsedPct)
		r[7] = fmt.Sprint(f.Inodes)
		r[8] = fmt.Sprint(f.InodesFree)
		r[9] = fmt.Sprint(f.InodesUsed)
		r[10] = fmt.Sprintf("%.2f%%", f.InodesUsedPct)
		r[11] = f.MountedOn
		r[12] = strings.Join(f.MountOptions, ",")
		r[13] = fmt.Sprintf("%t", f.IsLocal)
		r[14] = fmt.Sprintf("%t", f.IsAutomounted)
		t = append(t, r)
	}
	return h, t
}
