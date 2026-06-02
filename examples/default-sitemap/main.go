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
}
