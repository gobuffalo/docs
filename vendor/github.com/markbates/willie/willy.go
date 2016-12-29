package willie

import (
	"fmt"
	"net/http"
)

type encodable interface {
	Encode() string
}

type Willie struct {
	http.Handler
	Cookies    string
	Headers    map[string]string
	HmaxSecret string
}

func (w *Willie) Request(u string, args ...interface{}) *Request {
	hs := map[string]string{}
	for key, val := range w.Headers {
		hs[key] = val
	}
	return &Request{
		URL:     fmt.Sprintf(u, args...),
		Willie:  w,
		Headers: hs,
	}
}

func (w *Willie) JSON(u string, args ...interface{}) *JSON {
	hs := map[string]string{}
	for key, val := range w.Headers {
		hs[key] = val
	}
	hs["Content-Type"] = "application/json"
	return &JSON{
		URL:     fmt.Sprintf(u, args...),
		Willie:  w,
		Headers: hs,
	}
}

func New(h http.Handler) *Willie {
	return &Willie{
		Handler: h,
		Headers: map[string]string{},
	}
}
