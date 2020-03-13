// +build darwin

package user

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/resources"
	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/services"
)

func ListUsers() (*services.ListUsersResponse, error) {
	out, err := exec.Command("dscacheutil", "-q", "user").Output()
	if err != nil {
		return nil, err
	}

	hostname, _ := os.Hostname()
	resp := &services.ListUsersResponse{
		Hostname: hostname,
		Users:    []*resources.User{},
	}
	u := &resources.User{}
	i := 0
	for _, l := range strings.Split(string(out), "\n") {
		kv := strings.SplitN(l, ": ", 2)
		switch kv[0] {
		case "name":
			u.Username = kv[1]
			i++
		case "password":
			i++
			continue
		case "uid":
			signed, err := strconv.ParseInt(kv[1], 10, 64)
			if err != nil {
				return nil, err
			}
			u.UidSigned = signed
			i++
		case "gid":
			signed, err := strconv.ParseInt(kv[1], 10, 64)
			if err != nil {
				return nil, err
			}
			u.GidSigned = signed
			i++
		case "dir":
			u.Directory = kv[1]
			i++
		case "shell":
			u.Shell = kv[1]
			i++
		case "gecos":
			u.Comment = kv[1]
			i++
		case "":
			if i >= 7 {
				resp.Users = append(resp.Users, u)
			}
			u = &resources.User{}
			i = 0
		default:
			return nil, fmt.Errorf("unknown key: %s", kv[0])
		}
	}

	return resp, nil
}
