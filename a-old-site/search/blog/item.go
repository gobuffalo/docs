package blog

import "sync"

var lastPosts [3]Item
var mu = &sync.RWMutex{}

type Item struct {
	Title       string `json:"title"`
	PubDate     string `json:"pubDate"`
	Link        string `json:"link"`
	GUID        string `json:"guid"`
	Author      string `json:"author"`
	Thumbnail   string `json:"thumbnail"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Enclosure   struct {
	} `json:"enclosure"`
	Categories []string `json:"categories"`
}

func LastPosts() [3]Item {
	mu.RLock()
	defer mu.RUnlock()
	return lastPosts
}
