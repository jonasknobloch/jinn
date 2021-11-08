package tree

import (
	"strconv"
	"strings"
)

type Position []int

func (p Position) String() string {
	var sb strings.Builder

	for _, s := range p {
		sb.WriteString(strconv.Itoa(s))
	}

	return sb.String()
}

func PositionFromString(s string) ([]int, error) {
	p := make([]int, 0)

	for _, c := range strings.Split(s, "") {
		step, err := strconv.Atoi(c)

		if err != nil {
			return nil, err
		}

		p = append(p, step)
	}

	return p, nil
}
