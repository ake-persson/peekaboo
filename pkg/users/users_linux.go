// +build linux

package users

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

	users := &services.ListUsersResponse{Users: []*resources.User{}}
	for _, l := range strings.Split(string(b), "\n") {
		a := strings.Split(l, ":")

		uid, err := strconv.ParseUint(a[2], 10, 64)
		if err != nil {
			return nil, err
		}

		gid, err := strconv.ParseUint(a[3], 10, 64)
		if err != nil {
			return nil, err
		}

		users.Users = append(users.Users, &resources.User{
			Username:  a[0],
			Uid:       uid,
			Gid:       gid,
			Name:      a[4],
			Directory: a[5],
			Shell:     a[6],
		})
	}

	return users, nil
}
