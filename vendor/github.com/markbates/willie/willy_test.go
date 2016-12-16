package willie_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gorilla/pat"
	"github.com/gorilla/sessions"
	"github.com/markbates/willie"
	"github.com/stretchr/testify/require"
)

var Store sessions.Store = sessions.NewCookieStore([]byte("something-very-secret"))

type User struct {
	Name string `form:"name"`
}

func App() http.Handler {
	p := pat.New()
	p.Get("/get", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(201)
		fmt.Fprintln(res, "METHOD:"+req.Method)
		fmt.Fprint(res, "Hello from Get!")
	})
	p.Delete("/delete", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(201)
		fmt.Fprintln(res, "METHOD:"+req.Method)
		fmt.Fprint(res, "Goodbye")
	})
	p.Post("/post", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "METHOD:"+req.Method)
		fmt.Fprint(res, "NAME:"+req.PostFormValue("name"))
	})
	p.Put("/put", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(res, "METHOD:"+req.Method)
		fmt.Fprint(res, "NAME:"+req.PostFormValue("name"))
	})
	p.Post("/sessions/set", func(res http.ResponseWriter, req *http.Request) {
		sess, _ := Store.Get(req, "my-session")
		sess.Values["name"] = req.PostFormValue("name")
		sess.Save(req, res)
	})
	p.Get("/sessions/get", func(res http.ResponseWriter, req *http.Request) {
		sess, _ := Store.Get(req, "my-session")
		if sess.Values["name"] != nil {
			fmt.Fprint(res, "NAME:"+sess.Values["name"].(string))
		}
	})
	return p
}

func Test_Sessions(t *testing.T) {
	r := require.New(t)
	w := willie.New(App())

	res := w.Request("/sessions/get").Get()
	r.NotContains(res.Body.String(), "mark")
	w.Request("/sessions/set").Post(User{Name: "mark"})
	res = w.Request("/sessions/get").Get()
	r.Contains(res.Body.String(), "mark")
}

func Test_Request_URL_Params(t *testing.T) {
	r := require.New(t)
	w := willie.New(App())

	req := w.Request("/foo?a=%s&b=%s", "A", "B")
	r.Equal("/foo?a=A&b=B", req.URL)
}

func Test_Request_Copies_Headers(t *testing.T) {
	r := require.New(t)
	w := willie.New(App())
	w.Headers["foo"] = "bar"

	req := w.Request("/")
	r.Equal("bar", req.Headers["foo"])
}
