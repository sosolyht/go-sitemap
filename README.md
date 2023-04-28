[![Go Report Card](https://goreportcard.com/badge/github.com/sosolyht/go-sitemap)](https://goreportcard.com/report/github.com/sosolyht/go-sitemap)
# Go Sitemap Generator

A flexible library for creating sitemap and google video sitemaps in Go

## Installation

```bash
go get -u github.com/sosolyht/go-sitemap/sitemap
```

## Usage
Here is an example of how to use the Go Sitemap Generator library to create both standard and video sitemaps.

```go
package main

import "github.com/sosolyht/go-sitemap/sitemap"

func main() {
// Create a standard sitemap
s := sitemap.NewSitemap()

	links := []string{
		"https://google.com",
		"https://naver.com",
	}
	for _, link := range links {
		s.AddURL(link)
	}

	// Create a video sitemap
	vs := sitemap.NewVideoSitemap()

	videoURLs := []sitemap.VideoURL{
		// ... (the example video URLs)
	}

	for _, videoURL := range videoURLs {
		vs.AddVideoURL(videoURL)
	}
}
```

Replace the videoURLs variable with your video URLs and their respective video information.
The sitemaps will be generated and saved as `sitemaps/sitemap.xml` and `sitemaps/sitemap_video.xml` respectively.