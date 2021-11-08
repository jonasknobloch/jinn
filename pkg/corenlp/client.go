package corenlp

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	endpoint   *url.URL
	properties string
	httpClient http.Client
}

func NewClient(endpoint *url.URL, properties Properties) (*Client, error) {
	p, err := json.Marshal(properties)

	if err != nil {
		return nil, err
	}

	return &Client{
		endpoint:   endpoint,
		properties: string(p),
		httpClient: http.Client{},
	}, nil
}

// Run executes the annotation pipeline for the given input text and returns the raw results.
func (c *Client) Run(text string) (io.ReadCloser, error) {
	q := c.endpoint.Query()
	q.Add("properties", c.properties)
	c.endpoint.RawQuery = q.Encode()

	req, err := http.NewRequest("POST", c.endpoint.String(), strings.NewReader(text))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	return res.Body, nil
}

// Annotate executes the annotation pipeline for the given input texts and return the results as a Document. Only the
// ParserAnnotator and DependencyParseAnnotator with their respective required annotators are fully supported for now.
func (c *Client) Annotate(text string) (*Document, error) {
	rc, err := c.Run(text)

	if err != nil {
		return nil, err
	}

	defer rc.Close()

	doc := &Document{}

	if err := json.NewDecoder(rc).Decode(doc); err != nil {
		return nil, err
	}

	return doc, nil
}
