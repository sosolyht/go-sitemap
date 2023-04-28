package main

import "github.com/sosolyht/go-sitemap/sitemap"

func main() {
	s := sitemap.NewSitemap()

	links := []string{
		"https://google.com",
		"https://naver.com",
	}
	for i := range links {
		s.AddURL(links[i])
	}
}
