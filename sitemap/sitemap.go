package sitemap

import (
	"encoding/xml"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	XMLNS = "http://www.sitemaps.org/schemas/sitemap/0.9"
)

type ChangeFrequency string

const (
	ALWAYS  ChangeFrequency = "always"
	HOURLY  ChangeFrequency = "hourly"
	DAILY   ChangeFrequency = "daily"
	WEEKLY  ChangeFrequency = "weekly"
	MONTHLY ChangeFrequency = "monthly"
	YEARLY  ChangeFrequency = "yearly"
	NEVER   ChangeFrequency = "never"
)

type Sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	URL     []URLs   `xml:"url,omitempty"`
}

type URLs struct {
	Loc        string           `xml:"loc"`
	LastMod    string           `xml:"lastmod"`
	ChangeFreq *ChangeFrequency `xml:"changefreq,omitempty"`
	Priority   *float32         `xml:"priority,omitempty"`
}

func NewSitemap() *Sitemap {
	return &Sitemap{
		Xmlns: XMLNS,
	}
}

// AddURL
// Google ignores ChangeFrequency and Priority
// https://developers.google.com/search/docs/crawling-indexing/sitemaps/build-sitemap
func (s *Sitemap) AddURL(url string) (err error) {
	var urls []string
	if url != "" {
		urls = []string{
			url,
		}
	} else {
		urls, err = s.CreateSitemapFromLinksFile()
		if err != nil {
			return err
		}
	}

	for _, v := range urls {
		lastMod, merr := s.GetLastModifiedOrNow(v)
		if merr != nil {
			return merr
		}
		s.URL = append(s.URL, URLs{
			Loc:     url,
			LastMod: lastMod,
		})
	}

	xmlBytes, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	sitemapFile, err := os.Create("sitemaps/sitemap.xml")
	if err != nil {
		return err
	}

	defer sitemapFile.Close()

	if _, err = sitemapFile.Write([]byte(xml.Header)); err != nil {
		return err
	}

	if _, err = sitemapFile.Write(xmlBytes); err != nil {
		return err
	}

	return
}

func (s *Sitemap) CreateSitemapFromLinksFile() ([]string, error) {
	linkFile, err := os.Open("sitemaps/links")
	if err != nil {
		return nil, err
	}
	defer linkFile.Close()

	var links []string
	data, err := io.ReadAll(linkFile)
	if err != nil {
		return nil, err
	}

	splitLinks := strings.Split(string(data), "\n")
	for i := range splitLinks {
		links = append(links, splitLinks[i])
	}

	return splitLinks, err
}

func (s *Sitemap) GetLastModifiedOrNow(url string) (string, error) {
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Last-Modified
	data, err := http.Get(url)
	if err != nil {
		return "", err
	}
	lastModified := data.Header["Last-Modified"]

	defer data.Body.Close()

	var lastMod string
	if len(lastModified) == 0 {
		lastMod = time.Now().Format("2006-01-02")
	} else {
		parseTime, perr := time.Parse(time.RFC1123, lastModified[0])
		if perr != nil {
			return "", perr
		}

		lastMod = parseTime.Format("2006-01-02")
	}
	return lastMod, err
}

// CollectLinksFromURL
// TODO
//func (s *Sitemap) CollectLinksFromURL(url string) error {
//	http.Get(url)
//	return nil
//}
