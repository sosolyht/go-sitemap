[![Go Report Card](https://goreportcard.com/badge/github.com/sosolyht/go-sitemap)](https://goreportcard.com/report/github.com/sosolyht/go-sitemap)
# Go Sitemap Generator

A flexible library for creating sitemap and google video sitemaps in Go

## Installation

```bash
go get github.com/sosolyht/go-sitemap/sitemap@v1.0.0
```

## Usage
Here is an example of how to use the Go Sitemap Generator library to create both standard and video sitemaps.

```go
package main

import (
	"log"

	"github.com/sosolyht/go-sitemap/sitemap"
)

func main() {
	s, err := sitemap.NewSitemap().Path("sitemaps")
	if err != nil {
		log.Fatal(err)
	}

	links := []string{
		"https://google.com",
		"https://naver.com",
	}
	for _, link := range links {
		if err := s.AddURL(link); err != nil {
			log.Fatal(err)
		}
	}
	if err := s.Save(); err != nil {
		log.Fatal(err)
	}

	vs, err := sitemap.NewVideoSitemap().Path("sitemaps")
	if err != nil {
		log.Fatal(err)
	}

	videoURLs := []sitemap.VideoURL{
		// ... (the example video URLs)
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
```

Replace the videoURLs variable with your video URLs and their respective video information.
The sitemaps will be generated and saved as `sitemaps/sitemap.xml` and `sitemaps/sitemap_video.xml` respectively.