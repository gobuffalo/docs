package search

type Document struct {
	URL      string `json:"url"`
	Body     string `json:"body"`
	Language int    `json:"language"`
	Source   int    `json:"source"`
}
