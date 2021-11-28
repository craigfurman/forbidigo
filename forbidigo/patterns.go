package forbidigo

import (
	"fmt"
	"regexp"
	"regexp/syntax"
	"strings"
)

type pattern struct {
	pattern *regexp.Regexp
	msg     string
}

func parse(ptrn string) (*pattern, error) {
	p := &pattern{}
	parsedPattern, err := syntax.Parse(ptrn, syntax.Perl)
	if err != nil {
		return nil, fmt.Errorf("unable to parse pattern: %s: %s", ptrn, err)
	}
	if len(parsedPattern.Sub) == 0 {
		p.pattern, err = regexp.Compile(parsedPattern.String())
		if err != nil {
			return nil, fmt.Errorf("unable to compile pattern: %s: %s", ptrn, err)
		}
		return p, nil
	}
	p.pattern, err = regexp.Compile(parsedPattern.Sub[0].String())
	if err != nil {
		return nil, fmt.Errorf("unable to compile pattern: %s: %s", ptrn, err)
	}
	if len(parsedPattern.Sub) < 2 {
		return p, nil
	}
	msgPattern := deepestSubmatch(parsedPattern).String()
	p.msg = strings.TrimSpace(strings.TrimPrefix(msgPattern, "#"))

	return p, nil
}

func deepestSubmatch(expr *syntax.Regexp) *syntax.Regexp {
	for {
		if len(expr.Sub) == 0 {
			return expr
		}
		expr = expr.Sub[len(expr.Sub)-1]
	}
}
