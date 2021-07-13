package group

import (
	"fmt"
	"strings"

	"github.com/ake-persson/peekaboo/pkg/pb/v1/resources"
)

var Headers = []string{
	"groupname",
	"gid",
	"gid_signed",
	"members",
}

func StringSlice(g *resources.Group) []string {
	return []string{
		g.Groupname,
		fmt.Sprint(g.Gid),
		fmt.Sprint(g.GidSigned),
		strings.Join(g.Members, ","),
	}
}
