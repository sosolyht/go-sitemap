package main

import (
	"log"

	"github.com/sosolyht/go-sitemap/sitemap"
)

func main() {
	vs, err := sitemap.NewVideoSitemap().Path("sitemaps")
	if err != nil {
		log.Fatal(err)
	}

	videoURLs := []sitemap.VideoURL{
		{
			Loc: "https://www.example.com/videos/video1.html",
			Videos: []sitemap.Video{
				{
					ThumbnailLoc: "https://www.example.com/thumbnail/thumbnail1.png",
					Title:        "example1",
					Description:  "example1 desc",
					ContentLoc:   "https://www.example.com",
					PlayerLoc:    "https://www.example.com",
				},
			},
		},
		{
			Loc: "https://www.example.com/videos/video2.html",
			Videos: []sitemap.Video{
				{
					ThumbnailLoc: "https://www.example.com/thumbnail/thumbnail2.png",
					Title:        "example2",
					Description:  "example2 desc",
					ContentLoc:   "https://www.example.com",
					PlayerLoc:    "https://www.example.com",
				},
			},
		},
		{
			Loc: "https://www.example.com/videos/video3.html",
			Videos: []sitemap.Video{
				{
					ThumbnailLoc: "https://www.example.com/thumbnail/thumbnail3.png",
					Title:        "example3",
					Description:  "example3 desc",
					ContentLoc:   "https://www.example.com",
					PlayerLoc:    "https://www.example.com",
				},
			},
		},
	}

	for _, videoURL := range videoURLs {
		if err := vs.AddVideoURL(videoURL); err != nil {
			log.Fatal(err)
		}
	}

	if err := vs.Save(); err != nil {
		log.Fatal(err)
	}
}
