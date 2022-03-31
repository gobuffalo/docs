package blog

const FeedURL = "https://api.rss2json.com/v1/api.json?rss_url=https://blog.gobuffalo.io/feed"

type Feed struct {
	Status string `json:"status"`
	Feed   struct {
		URL         string `json:"url"`
		Title       string `json:"title"`
		Link        string `json:"link"`
		Author      string `json:"author"`
		Description string `json:"description"`
		Image       string `json:"image"`
	} `json:"feed"`
	Items []Item `json:"items"`
}
