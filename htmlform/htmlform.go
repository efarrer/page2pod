package htmlform

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Request struct {
	Title   string
	Podcast string
	Content string
	Credentials
}

type Credentials struct {
	Username string
	Password string
}

func first(in []string) string {
	if len(in) > 0 {
		return in[0]
	}
	return ""
}

func Parse(body string) (*Request, error) {
	// Go's http.Request has a build in parser for form submissions so (ab)use that
	req := http.Request{
		Method: "POST",
		Header: map[string][]string{"Content-Type": []string{"application/x-www-form-urlencoded"}},
		Body:   io.NopCloser(strings.NewReader(body)),
	}

	err := req.ParseForm()
	if err != nil {
		return nil, fmt.Errorf("Unable to parse form data: %w", err)
	}

	request := &Request{
		Title:   first(req.PostForm["title"]),
		Podcast: first(req.PostForm["podcast"]),
		Content: first(req.PostForm["content"]),
		Credentials: Credentials{
			Username: first(req.PostForm["username"]),
			Password: first(req.PostForm["password"]),
		},
	}

	return request, nil
}
