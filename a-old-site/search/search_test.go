package search

import "log"

func init() {
	docs := []Document{
		{URL: "godoc.org/foo", Body: "this is a godoc foo; a foo by any other name is still a foo", Language: L_EN, Source: S_GODOC},
		{URL: "/en/foo", Body: "this is an english foo", Language: L_EN, Source: S_SITE},
		{URL: "/fr/foo", Body: "this is a french foo", Language: L_FR, Source: S_SITE},
	}
	for _, d := range docs {
		if err := Index(d); err != nil {
			log.Fatal(err)
		}
	}
}
