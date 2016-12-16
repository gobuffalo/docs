package willy

import (
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"

	"github.com/ajg/form"
	"github.com/markbates/hmax"
)

type Request struct {
	URL      string
	Willy    *Willy
	Headers  map[string]string
	Username string
	Password string
}

func (r *Request) SetBasicAuth(username, password string) {
	r.Username = username
	r.Password = password
}

func toReader(body interface{}) io.Reader {
	modelType := reflect.TypeOf((*encodable)(nil)).Elem()
	s := reflect.ValueOf(body)
	t := s.Type()
	if !t.Implements(modelType) {
		body, _ = form.EncodeToValues(body)
	}
	return strings.NewReader(body.(encodable).Encode())
}

func (r *Request) Get() *Response {
	req, _ := http.NewRequest("GET", r.URL, nil)
	return r.perform(req)
}

func (r *Request) Delete() *Response {
	req, _ := http.NewRequest("DELETE", r.URL, nil)
	return r.perform(req)
}

func (r *Request) Post(body interface{}) *Response {
	req, _ := http.NewRequest("POST", r.URL, toReader(body))
	r.Headers["Content-Type"] = "application/x-www-form-urlencoded"
	return r.perform(req)
}

func (r *Request) Put(body interface{}) *Response {
	req, _ := http.NewRequest("PUT", r.URL, toReader(body))
	r.Headers["Content-Type"] = "application/x-www-form-urlencoded"
	return r.perform(req)
}

func (r *Request) perform(req *http.Request) *Response {
	if r.Willy.HmaxSecret != "" {
		hmax.SignRequest(req, []byte(r.Willy.HmaxSecret))
	}
	if r.Username != "" || r.Password != "" {
		req.SetBasicAuth(r.Username, r.Password)
	}
	res := &Response{httptest.NewRecorder()}
	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Cookie", r.Willy.Cookies)
	r.Willy.ServeHTTP(res, req)
	c := res.HeaderMap["Set-Cookie"]
	r.Willy.Cookies = strings.Join(c, ";")
	return res
}
