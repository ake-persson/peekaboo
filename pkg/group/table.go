package group

import (
	"fmt"
	"strings"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/resources"
	"github.com/peekaboo-labs/peekaboo/pkg/text"
)

func ToTable(hostname string, groups []*resources.Group) *text.Table {
	t := text.Table{
		Headers: []string{"hostname", "groupname", "gid", "gid_signed", "members"},
		Rows:    make([][]string, 0),
	}

	for _, g := range groups {
		r := make([]string, 5)
		r[0] = hostname
		r[1] = g.Groupname
		r[2] = fmt.Sprint(g.Gid)
		r[3] = fmt.Sprint(g.GidSigned)
		r[4] = strings.Join(g.Members, ",")
		t.Rows = append(t.Rows, r)
	}

	return &t
}
