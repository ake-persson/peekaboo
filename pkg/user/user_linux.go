// +build linux

package user

import (
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/resources"
	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/services"
)

func ListUsers() (*services.ListUsersResponse, error) {
	b, err := ioutil.ReadFile("/etc/passwd")
	if err != nil {
		return nil, err
	}

	resp := &services.ListUsersResponse{Users: []*resources.User{}}
	for _, l := range strings.Split(string(b), "\n") {
		a := strings.Split(l, ":")

		if len(a) < 6 {
			continue
		}

		uid, err := strconv.ParseUint(a[2], 10, 64)
		if err != nil {
			return nil, err
		}

		gid, err := strconv.ParseUint(a[3], 10, 64)
		if err != nil {
			return nil, err
		}

		resp.Users = append(resp.Users, &resources.User{
			Username:  a[0],
			Uid:       uid,
			Gid:       gid,
			Comment:   a[4],
			Directory: a[5],
			Shell:     a[6],
		})
	}

	return resp, nil
}
