package user

import "strings"

func splitOmitEmpty(s string, del string) []string {
	out := []string{}
	for _, v := range strings.Split(s, del) {
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}
