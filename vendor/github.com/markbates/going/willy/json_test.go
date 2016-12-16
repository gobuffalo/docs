package willy_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/gorilla/pat"
	"github.com/markbates/going/willy"
	"github.com/stretchr/testify/require"
)

type jBody struct {
	Method  string `json:"method"`
	Name    string `json:"name"`
	Message string `json:"message"`
}

func JSONApp() http.Handler {
	p := pat.New()
	p.Get("/get", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(201)
		json.NewEncoder(res).Encode(jBody{
			Method:  req.Method,
			Message: "Hello from Get!",
		})
	})
	p.Delete("/delete", func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(201)
		json.NewEncoder(res).Encode(jBody{
			Method:  req.Method,
			Message: "Goodbye",
		})
	})
	p.Post("/post", func(res http.ResponseWriter, req *http.Request) {
		jb := jBody{}
		json.NewDecoder(req.Body).Decode(&jb)
		jb.Method = req.Method
		json.NewEncoder(res).Encode(jb)
	})
	p.Put("/put", func(res http.ResponseWriter, req *http.Request) {
		jb := jBody{}
		json.NewDecoder(req.Body).Decode(&jb)
		jb.Method = req.Method
		json.NewEncoder(res).Encode(jb)
	})
	p.Post("/sessions/set", func(res http.ResponseWriter, req *http.Request) {
		sess, _ := Store.Get(req, "my-session")
		sess.Values["name"] = req.PostFormValue("name")
		sess.Save(req, res)
	})
	p.Get("/sessions/get", func(res http.ResponseWriter, req *http.Request) {
		sess, _ := Store.Get(req, "my-session")
		if sess.Values["name"] != nil {
			json.NewEncoder(res).Encode(jBody{
				Method: req.Method,
				Name:   sess.Values["name"].(string),
			})
		}
	})
	return p
}

func Test_JSON_Headers_Dont_Overwrite_App_Headers(t *testing.T) {
	r := require.New(t)
	w := willy.New(JSONApp())
	w.Headers["foo"] = "bar"

	req := w.JSON("/")
	req.Headers["foo"] = "baz"
	r.Equal("baz", req.Headers["foo"])
	r.Equal("bar", w.Headers["foo"])
}

func Test_JSON_Get(t *testing.T) {
	r := require.New(t)
	w := willy.New(JSONApp())

	req := w.JSON("/get")
	r.Equal("/get", req.URL)

	res := req.Get()
	r.Equal(201, res.Code)
	jb := &jBody{}
	res.Bind(jb)
	r.Equal("GET", jb.Method)
	r.Equal("Hello from Get!", jb.Message)
}

func Test_JSON_Delete(t *testing.T) {
	r := require.New(t)
	w := willy.New(JSONApp())

	req := w.JSON("/delete")
	r.Equal("/delete", req.URL)

	res := req.Delete()
	jb := &jBody{}
	res.Bind(jb)
	r.Equal("DELETE", jb.Method)
	r.Equal("Goodbye", jb.Message)
}

func Test_JSON_Post_Struct(t *testing.T) {
	r := require.New(t)
	w := willy.New(JSONApp())

	req := w.JSON("/post")
	res := req.Post(User{Name: "Mark"})

	jb := &jBody{}
	res.Bind(jb)
	r.Equal("POST", jb.Method)
	r.Equal("Mark", jb.Name)
}

func Test_JSON_Post_Struct_Pointer(t *testing.T) {
	r := require.New(t)
	w := willy.New(JSONApp())

	req := w.JSON("/post")
	res := req.Post(&User{Name: "Mark"})

	jb := &jBody{}
	res.Bind(jb)
	r.Equal("POST", jb.Method)
	r.Equal("Mark", jb.Name)
}

func Test_JSON_Put(t *testing.T) {
	r := require.New(t)
	w := willy.New(JSONApp())

	req := w.JSON("/put")
	res := req.Put(User{Name: "Mark"})

	jb := &jBody{}
	res.Bind(jb)
	r.Equal("PUT", jb.Method)
	r.Equal("Mark", jb.Name)
}

func Test_JSON_Put_Struct_Pointer(t *testing.T) {
	r := require.New(t)
	w := willy.New(JSONApp())

	req := w.JSON("/put")
	res := req.Put(&User{Name: "Mark"})

	jb := &jBody{}
	res.Bind(jb)
	r.Equal("PUT", jb.Method)
	r.Equal("Mark", jb.Name)
}
