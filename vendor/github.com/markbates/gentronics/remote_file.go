package gentronics

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type RemoteFile struct {
	RemotePath string
	*File
}

func (f *RemoteFile) Run(rootPath string, data Data) error {
	if !f.Should(data) {
		return nil
	}

	res, err := http.Get(f.RemotePath)
	if err != nil {
		return errors.WithStack(err)
	}
	code := res.StatusCode
	if code < 200 || code >= 300 {
		return errors.WithStack(fmt.Errorf("Error fetching %s (%d)", f.RemotePath, code))
	}

	bb := bytes.Buffer{}
	_, err = io.Copy(&bb, res.Body)
	if err != nil {
		return errors.WithStack(err)
	}
	return f.save(rootPath, f.Path, bb.String())
}
