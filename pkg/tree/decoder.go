package tree

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type decoder struct{}

func NewDecoder() *decoder {
	return &decoder{}
}

func (d *decoder) Decode(s string) (*Tree, error) {
	re := regexp.MustCompile(`\(\s*([^\s()]+)?|\)|([^\s()]+)`)

	stack := make([]*Tree, 0)
	stack = append(stack, &Tree{})

	s = strings.NewReplacer("\n", " ", "\t", " ", "\r", " ").Replace(s)

	for _, m := range re.FindAllStringIndex(s, -1) {
		t := s[m[0]:m[1]]
		switch t[:1] {
		case "(":
			if len(stack) == 1 && len(stack[0].Children) > 0 {
				return nil, d.parseError(t, "end-of-string", m)
			}

			label := strings.TrimLeft(t[1:], " ")
			tree := &Tree{label, make([]*Tree, 0)}
			stack = append(stack, tree)
		case ")":
			if len(stack) == 1 {
				if len(stack[0].Children) == 0 {
					return nil, d.parseError(t, "(", m)
				}

				return nil, d.parseError(t, "end-of-string", m)
			}

			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, last)
		default:
			if len(stack) == 1 {
				return nil, d.parseError(t, "(", m)
			}

			tree := &Tree{t, nil}
			stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, tree)
		}
	}

	if len(stack) > 1 {
		return nil, d.parseError("end-of-string", ")", []int{len(s)})
	}

	if len(stack[0].Children) == 0 {
		return nil, d.parseError("end-of-string", "(", []int{len(s)})
	}

	if stack[0].Label != "" || len(stack[0].Children) != 1 {
		return nil, errors.New("malformed end result")
	}

	return stack[0].Children[0], nil
}

func (d *decoder) parseError(t, e string, pos []int) error {
	return fmt.Errorf("expected \"%s\" but got \"%s\" at postion %v", e, t, pos)
}
