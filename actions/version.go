package actions

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const latestURL = "https://api.github.com/repos/gobuffalo/buffalo/releases/latest"

var buffaloVersion = findLatestVersion()

func init() {
	go func() {
		for {
			time.Sleep(60 * time.Minute)
			buffaloVersion = findLatestVersion()
		}
	}()
}

func findLatestVersion() string {
	res, err := http.Get(latestURL)
	if err != nil {
		return "unknown"
	}
	b, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return "unknown"
	}

	gh := struct {
		Tag string `json:"tag_name"`
	}{}

	if err := json.Unmarshal(b, &gh); err != nil {
		return "unknown"
	}
	if len(gh.Tag) == 0 {
		return "unknown"
	}
	return gh.Tag
}
