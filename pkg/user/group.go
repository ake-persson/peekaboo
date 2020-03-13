package user

import (
	"fmt"
	"strings"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/resources"
)

func GroupsToStringTable(hostname string, groups []*resources.Group) ([]string, [][]string) {
	h := []string{"hostname", "groupname", "gid", "gid_signed", "members"}
	t := make([][]string, 0)
	for _, g := range groups {
		r := make([]string, 5)
		r[0] = hostname
		r[1] = g.Groupname
		r[2] = fmt.Sprint(g.Gid)
		r[3] = fmt.Sprint(g.GidSigned)
		r[4] = strings.Join(g.Members, ",")
		t = append(t, r)
	}
	return h, t
}
