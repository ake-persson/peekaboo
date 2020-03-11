// +build linux

package user

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/resources"
	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/services"
)

func ListGroups() (*services.ListGroupsResponse, error) {
	b, err := ioutil.ReadFile("/etc/group")
	if err != nil {
		return nil, err
	}

	resp := &services.ListGroupsResponse{Groups: []*resources.Group{}}
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
			Members:   splitOmitEmpty(a[3], ","),
		})
	}

	return resp, nil
}
