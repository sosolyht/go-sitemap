package main

import (
	"github.com/sosolyht/go-sitemap/sitemap"
	"log"
)

func main() {
	err := sitemap.NewSitemap().AddURL(nil)
	if err != nil {
		log.Fatal(err)
	}
}
