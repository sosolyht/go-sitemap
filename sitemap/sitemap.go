package sitemap

import (
	"encoding/xml"
	"os"
)

const (
	XmlVersion = "<?xml version=\"1.0″ encoding=\"UTF-8″?>"
	URLSet     = "<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9″>"
)

type Sitemap struct {
	XMLVersion string `xml:"XMLVersion,omitempty"`
	URLSet     string `xml:"URLSet,omitempty"`
	URL        []URL  `xml:"url,omitempty"`
}

type URL struct {
	Loc      string  `xml:"loc"`
	LastMod  string  `xml:"lastmod"`
	Freq     string  `xml:"changefreq"`
	Priority float32 `xml:"priority"`
}

func NewURL() *Sitemap {
	return &Sitemap{
		XMLVersion: XmlVersion,
		URLSet:     URLSet,
	}
}

func (s *Sitemap) AddURL(url URL) ([]byte, error) {
	s.URL = append(s.URL, url)

	resp, err := xml.MarshalIndent(s.URL, "", " ")
	if err != nil {
		panic(err)
	}

	m, err := os.Create("sitemap.xml")
	if err != nil {
		panic(err)
	}

	defer m.Close()

	m.Write(resp)

	return resp, err
}
