package parserdto

import (
	"html"
	"regexp"
	"strings"
)

func TrimLastSlash(u string, count int) string {
	c := 0

	ind := strings.LastIndexFunc(u, func(r rune) bool {
		if r != rune('/') {
			return false
		}

		c++

		return c == count
	})

	return u[:ind]
}

func OneMatch(rgx *regexp.Regexp, raw string) (string, bool) {
	res := rgx.FindAllStringSubmatch(raw, -1)
	if len(res) < 1 || len(res[0]) != 2 {
		return "", false
	}

	return strings.TrimSpace(html.UnescapeString(res[0][1])), true
}
