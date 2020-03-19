// +build linux

package group

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/resources"
	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/services"
	"github.com/peekaboo-labs/peekaboo/pkg/text"
)

func ListGroups() (*services.ListGroupsResponse, error) {
	b, err := ioutil.ReadFile("/etc/group")
	if err != nil {
		return nil, err
	}

	hostname, _ := os.Hostname()
	resp := &services.ListGroupsResponse{
		Hostname: hostname,
		Groups:   []*resources.Group{},
	}
	for _, l := range strings.Split(string(b), "\n") {
		a := strings.Split(l, ":")
		if len(a) < 4 {
			continue
		}

		gid, err := strconv.ParseUint(a[2], 10, 64)
		if err != nil {
			return nil, err
		}

		resp.Groups = append(resp.Groups, &resources.Group{
			Groupname: a[0],
			Gid:       gid,
			Members:   text.Split(a[3], ","),
		})
	}

	return resp, nil
}
