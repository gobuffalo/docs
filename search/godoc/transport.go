package godoc

import (
	"net/http"

	"github.com/gobuffalo/envy"
	"github.com/pkg/errors"
)

type githubTrans struct{}

func (githubTrans) RoundTrip(req *http.Request) (*http.Response, error) {
	q := req.URL.Query()
	tok, err := envy.MustGet(tk)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	q.Add("access_token", tok)
	req.URL.RawQuery = q.Encode()
	return http.DefaultTransport.RoundTrip(req)
}
