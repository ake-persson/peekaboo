// +build darwin

package users

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/resources"
)

func GetUsers() (*resources.UserList, error) {
	out, err := exec.Command("dscacheutil", "-q", "user").Output()
	if err != nil {
		return nil, err
	}

	users := &resources.UserList{Users: []*resources.User{}}
	u := &resources.User{}
	for _, l := range strings.Split(string(out), "\n") {
		kv := strings.SplitN(l, ": ", 2)
		switch kv[0] {
		case "name":
			u.Username = kv[1]
		case "password":
			continue
		case "uid":
			i, err := strconv.ParseInt(kv[1], 10, 64)
			if err != nil {
				return nil, err
			}
			u.UidSigned = i
		case "gid":
			i, err := strconv.ParseInt(kv[1], 10, 64)
			if err != nil {
				return nil, err
			}
			u.GidSigned = i
		case "dir":
			u.Directory = kv[1]
		case "shell":
			u.Shell = kv[1]
		case "gecos":
			u.Description = kv[1]
		case "":
			users.Users = append(users.Users, u)
			u = &resources.User{}
		default:
			return nil, fmt.Errorf("unknown key: %s", kv[0])
		}
	}

	return users, nil
}
