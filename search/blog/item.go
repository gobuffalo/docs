package blog

var LastPosts [3]Item

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
