// +build darwin

package group

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/resources"
	"github.com/peekaboo-labs/peekaboo/pkg/pb/v1/services"
	"github.com/peekaboo-labs/peekaboo/pkg/text"
)

func ListGroups() (*services.ListGroupsResponse, error) {
	out, err := exec.Command("dscacheutil", "-q", "group").Output()
	if err != nil {
		return nil, err
	}

	hostname, _ := os.Hostname()
	resp := &services.ListGroupsResponse{
		Hostname: hostname,
		Groups:   []*resources.Group{},
	}
	g := &resources.Group{}
	i := 0
	for _, l := range strings.Split(string(out), "\n") {
		kv := strings.SplitN(l, ": ", 2)
		switch kv[0] {
		case "name":
			g.Groupname = kv[1]
			i++
		case "password":
			i++
			continue
		case "gid":
			signed, err := strconv.ParseInt(kv[1], 10, 64)
			if err != nil {
				return nil, err
			}
			g.GidSigned = signed
			i++
		case "users":
			g.Members = text.Split(kv[1], " ")
			i++
		case "":
			if i >= 3 {
				resp.Groups = append(resp.Groups, g)
			}
			g = &resources.Group{}
			i = 0
		default:
			return nil, fmt.Errorf("unknown key: %s", kv[0])
		}
	}

	return resp, nil
}
