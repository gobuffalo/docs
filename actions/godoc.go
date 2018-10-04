package actions

import (
	"fmt"
	"go/build"
	"path/filepath"

	"github.com/pkg/errors"
	"golang.org/x/tools/godoc"
	"golang.org/x/tools/godoc/vfs"
)

func indexGodocs() error {

	fs := vfs.NewNameSpace()
	c := build.Default
	fmt.Println("### c.SrcDirs() ->", c.SrcDirs())

	root := vfs.OS(build.Default.GOROOT)

	fs.Bind("/", root, "/", vfs.BindReplace)
	fmt.Println("### build.Default.GOPATH ->", build.Default.GOPATH)
	for _, p := range filepath.SplitList(build.Default.GOPATH) {
		fmt.Println("### p ->", p)
		fs.Bind("/src", vfs.OS(p), "/src", vfs.BindAfter)
	}
	corpus := godoc.NewCorpus(fs)
	corpus.Verbose = true
	if err := corpus.Init(); err != nil {
		return errors.WithStack(err)
	}
	return nil
}
