package user

import (
	"fmt"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/resources"
	"github.com/peekaboo-labs/peekaboo/pkg/text"
)

func ToTable(hostname string, users []*resources.User) *text.Table {
	t := text.Table{
		Headers: []string{"hostname", "username", "comment", "uid", "gid",
			"uid_signed", "gid_signed", "directory", "shell"},
		Rows: make([][]string, 0),
	}

	for _, u := range users {
		r := make([]string, 9)
		r[0] = hostname
		r[1] = u.Username
		r[2] = u.Comment
		r[3] = fmt.Sprint(u.Uid)
		r[4] = fmt.Sprint(u.Gid)
		r[5] = fmt.Sprint(u.UidSigned)
		r[6] = fmt.Sprint(u.GidSigned)
		r[7] = fmt.Sprint(u.Directory)
		r[8] = fmt.Sprint(u.Shell)
		t.Rows = append(t.Rows, r)
	}

	return &t
}
