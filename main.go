package main

import (
	"fmt"
	"github.com/sosolyht/go-sitemap/sitemap"
)

func main() {
	res, _ := sitemap.NewURL().AddURL(sitemap.URL{
		Loc:      "http://google.com",
		LastMod:  "2023-04-18",
		Freq:     "monthly",
		Priority: 0.5,
	})
	fmt.Println(string(res))
}
