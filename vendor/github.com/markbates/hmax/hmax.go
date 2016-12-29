package hmax

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"io/ioutil"
	"net/http"
)

func Sign(key, message []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write(message)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func Verify(signature string, key, message []byte) bool {
	s := Sign(key, message)
	return hmac.Equal([]byte(s), []byte(signature))
}

func SignRequest(req *http.Request, key []byte) error {
	b, err := readBody(req)
	if err != nil {
		return err
	}

	s := Sign(key, b)
	req.Header.Set("X-Signature", s)
	return nil
}

func VerifyRequest(req *http.Request, key []byte) (bool, error) {
	b, err := readBody(req)
	if err != nil {
		return false, err
	}
	return Verify(req.Header.Get("X-Signature"), key, b), nil
}

func readBody(req *http.Request) ([]byte, error) {
	var b []byte
	var err error
	if req.Body != nil {
		b, err = ioutil.ReadAll(req.Body)
	}

	// Restore the io.ReadCloser to its original state
	req.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	return b, err
}
