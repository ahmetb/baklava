package util

import (
	"fmt"
	"regexp"
)

func ExtractGroups(re *regexp.Regexp, in []byte) ([]string, error) {
	var out []string
	matches := re.FindAllSubmatch(in, -1)
	if len(matches) == 0 {
		return nil, fmt.Errorf("no matches for pattern in response")
	}
	for _, g := range matches {
		if len(g) != 2 {
			return nil, fmt.Errorf("oddly got non-2 matches (%d), in a group", len(g))
		}
		out = append(out, string(g[1]))
	}
	return out, nil
}
