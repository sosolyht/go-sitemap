package main

import (
	"fmt"
	"github.com/sosolyht/go-sitemap/sitemap"
)

func main() {
	err := sitemap.NewURL().AddURL(sitemap.URL{
		Loc:        "https://google.com",
		ChangeFreq: sitemap.MONTHLY,
		Priority:   0.5,
	})
	fmt.Println(err)
}
