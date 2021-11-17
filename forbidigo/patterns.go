package forbidigo

import (
	"fmt"
	"regexp"
)

var ptrnWithMsg = regexp.MustCompile(`(#\[(?P<msg>[^\]]+)\])?(?P<pattern>.+)`)

type pattern struct {
	pattern *regexp.Regexp
	msg     string
}

func parse(ptrn string) (*pattern, error) {
	p := &pattern{}
	matches := ptrnWithMsg.FindStringSubmatch(ptrn)
	for i, name := range ptrnWithMsg.SubexpNames() {
		if name == "msg" {
			p.msg = matches[i]
		} else if name == "pattern" {
			re, err := regexp.Compile(matches[i])
			if err != nil {
				return nil, fmt.Errorf("unable to compile pattern `%s`: %s", matches[i], err)
			}
			p.pattern = re
		}
	}

	return p, nil
}
