package pathx

import (
	"path"
	"path/filepath"
	"runtime"
)

func Here() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}

func FromHere(s ...string) string {
	args := []string{Here()}
	args = append(args, s...)
	return filepath.Join(args...)
}
