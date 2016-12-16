package hmax_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/markbates/hmax"
	"github.com/stretchr/testify/require"
)

var secret = []byte("password")
var message = []byte("secure message")
var signature = "nfVW5dkRrMtKxkn1gsCF0VeBi6/1ira0wmb3nW8YjK4="

func Test_Sign(t *testing.T) {
	r := require.New(t)

	s := hmax.Sign(secret, message)
	r.Equal(signature, s)
}

func Test_Verify(t *testing.T) {
	r := require.New(t)

	b := hmax.Verify(signature, secret, message)
	r.True(b)
}

func Test_SignRequest(t *testing.T) {
	r := require.New(t)

	rr := bytes.NewReader(message)
	req, err := http.NewRequest("GET", "/", rr)
	r.NoError(err)

	err = hmax.SignRequest(req, secret)
	r.NoError(err)

	xs := req.Header.Get("X-Signature")
	r.Equal(signature, xs)

	// ensure the body has been reset
	b, _ := ioutil.ReadAll(req.Body)
	r.Equal(message, b)
}

func Test_VerifyRequest(t *testing.T) {
	r := require.New(t)

	rr := bytes.NewReader(message)
	req, err := http.NewRequest("GET", "/", rr)
	r.NoError(err)
	req.Header.Set("X-Signature", signature)

	r.True(hmax.VerifyRequest(req, secret))
}
