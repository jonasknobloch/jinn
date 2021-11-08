package corenlp

import (
	"strings"
)

type Annotators []Annotator

func (as Annotators) Join() string {
	var sb strings.Builder

	for k, a := range as {
		if k > 0 {
			sb.WriteString(",")
		}

		sb.WriteString(string(a))
	}

	return sb.String()
}
