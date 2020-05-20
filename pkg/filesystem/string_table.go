package filesystem

import (
	"fmt"
	"strings"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/resources"
)

var Headers = []string{
	"filesystem",
	"type",
	"size_kb",
	"used_kb",
	"free_kb",
	"used_pct",
	"inodes",
	"inodes_free",
	"inodes_used",
	"inodes_used_pct",
	"mounted_on",
	"mount_options",
	"is_local",
	"is_automounted",
}

func StringSlice(f *resources.Filesystem) []string {
	return []String{
		f.Filesystem,
		f.Type,
		fmt.Sprintf("%dK", f.SizeKb),
		fmt.Sprintf("%dK", f.UsedKb),
		fmt.Sprintf("%dK", f.FreeKb),
		fmt.Sprintf("%.2f%%", f.UsedPct),
		fmt.Sprint(f.Inodes),
		fmt.Sprint(f.InodesFree),
		fmt.Sprint(f.InodesUsed),
		fmt.Sprintf("%.2f%%", f.InodesUsedPct),
		f.MountedOn,
		strings.Join(f.MountOptions, ","),
		fmt.Sprintf("%t", f.IsLocal),
		fmt.Sprintf("%t", f.IsAutomounted),
	}
}
