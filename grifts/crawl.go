package grifts

// import (
// 	"bytes"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strings"
// 	"sync"
//
// 	"golang.org/x/net/html"
// 	"golang.org/x/sync/errgroup"
//
// 	"github.com/blevesearch/bleve"
// 	"github.com/markbates/grift/grift"
// 	"github.com/pkg/errors"
// )
//
// var visited = map[string]bool{}
// var moot = &sync.Mutex{}
// var wg = &errgroup.Group{}
// var host = "http://localhost:3000"
// var index bleve.Index
// var cache = map[string]page{}
//
// type page struct {
// 	URL   string
// 	Title string
// 	Body  string
// }
//
// var _ = grift.Add("crawl", func(c *grift.Context) error {
// 	var err error
// 	mapping := bleve.NewIndexMapping()
// 	if _, err := os.Stat("gobuffalo.bleve"); err == nil {
// 		index, err = bleve.Open("gobuffalo.bleve")
// 	} else {
// 		index, err = bleve.New("gobuffalo.bleve", mapping)
// 	}
//
// 	if err != nil {
// 		return errors.WithStack(err)
// 	}
// 	wg.Go(func() error {
// 		return processURL(host)
// 	})
//
// 	err = wg.Wait()
// 	if err != nil {
// 		return errors.WithStack(err)
// 	}
//
// 	res, err := index.Search(bleve.NewSearchRequest(bleve.NewMatchQuery("sqlite3")))
// 	if err != nil {
// 		return errors.WithStack(err)
// 	}
// 	fmt.Printf("### res -> %+v\n", res)
// 	for _, v := range res.Hits {
// 		p := cache[v.ID]
// 		fmt.Printf("### p.Title -> %+v\n", p.Title)
// 	}
// 	return nil
// })
//
// func processURL(u string) error {
// 	moot.Lock()
// 	u = strings.Split(u, "#")[0]
// 	if visited[u] {
// 		moot.Unlock()
// 		return nil
// 	}
// 	visited[u] = true
// 	moot.Unlock()
// 	if u == "nofollow" {
// 		return nil
// 	}
// 	res, err := http.Get(u)
// 	if err != nil {
// 		return errors.WithStack(err)
// 	}
//
// 	page := page{
// 		URL: u,
// 	}
//
// 	bb := &bytes.Buffer{}
// 	defer func() {
// 		if bb.Len() > 0 {
// 			page.Body = bb.String()
// 			err := index.Index(page.URL, page)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			cache[page.URL] = page
// 		}
// 	}()
//
// 	z := html.NewTokenizer(res.Body)
// 	for {
// 		tt := z.Next()
// 		switch tt {
// 		case html.ErrorToken:
// 			return nil
// 		case html.StartTagToken:
// 			t := z.Token()
// 			if t.Data == "title" {
// 				tt = z.Next()
// 				title := string(bytes.TrimSpace(z.Text()))
// 				if title != "" {
// 					x := strings.Split(title, "-")
// 					title = strings.TrimSpace(x[len(x)-1])
// 					page.Title = title
// 				}
// 			}
// 			if t.Data == "a" {
// 				for _, a := range t.Attr {
// 					if a.Key == "href" {
// 						av := a.Val
// 						if strings.HasPrefix(av, "/") {
// 							av = host + av
// 						}
// 						if strings.HasPrefix(av, host) {
// 							wg.Go(func() error {
// 								return processURL(av)
// 							})
// 						}
// 					}
// 				}
// 			}
// 		case html.TextToken:
// 			b := bytes.TrimSpace(z.Raw())
// 			if len(b) > 0 {
// 				bb.Write(b)
// 			}
// 		}
// 	}
// }
