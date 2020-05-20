package user

import (
	"fmt"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/resources"
)

var Headers = []string{
	"username",
	"comment",
	"uid",
	"gid",
	"uid_signed",
	"gid_signed",
	"directory",
	"shell",
}

func ToTable(hostname string, users []*resources.User) *text.Table {
	return []string{
		u.Username,
		u.Comment,
		fmt.Sprint(u.Uid),
		fmt.Sprint(u.Gid),
		fmt.Sprint(u.UidSigned),
		fmt.Sprint(u.GidSigned),
		fmt.Sprint(u.Directory),
		fmt.Sprint(u.Shell),
	}
}
