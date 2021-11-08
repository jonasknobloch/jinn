package corenlp

import (
	"encoding/json"
)

type Properties struct {
	Annotators      Annotators
	OutputFormat    Format
	Serializer      Serializer
	InputFormat     Format
	InputSerializer Serializer
	Options         Options
}

func (p Properties) MarshalJSON() ([]byte, error) {
	pm := make(map[string]string)

	add := func(k, v string) {
		if v == "" {
			return
		}

		pm[k] = v
	}

	add("annotators", p.Annotators.Join())
	add("outputFormat", string(p.OutputFormat))
	add("serializer", string(p.Serializer))
	add("inputFormat", string(p.InputFormat))
	add("inputSerializer", string(p.InputSerializer))

	for k, v := range p.Options {
		add(k, v)
	}

	return json.Marshal(pm)
}
