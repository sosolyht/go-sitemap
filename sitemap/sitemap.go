package sitemap

import (
	"encoding/xml"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	URL     []URLs   `xml:"url,omitempty"`
	path    string
}

type URLs struct {
	Loc        string           `xml:"loc"`
	LastMod    string           `xml:"lastmod"`
	ChangeFreq *ChangeFrequency `xml:"changefreq,omitempty"`
	Priority   *float32         `xml:"priority,omitempty"`
}

func NewSitemap() *sitemap {
	return &sitemap{
		Xmlns: xmlns,
	}
}

// AddURL
// Google ignores ChangeFrequency and Priority
// https://developers.google.com/search/docs/crawling-indexing/sitemaps/build-sitemap
func (s *sitemap) AddURL(url string) (err error) {
	var urls []string
	if url != "" {
		urls = []string{
			url,
		}
	} else {
		urls, err = s.createSitemapFromLinksFile()
		if err != nil {
			return err
		}
	}

	for _, v := range urls {
		lastMod, merr := s.getLastModifiedOrNow(v)
		if merr != nil {
			return merr
		}
		s.URL = append(s.URL, URLs{
			Loc:     v,
			LastMod: lastMod,
		})
	}

	xmlBytes, err := xml.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	sitemapFile, err := os.Create(s.path)
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

func (s *sitemap) Path(path string) *sitemap {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	projectRoot := filepath.Join(currentDir, "..", "..")
	sitemapsDir := filepath.Join(projectRoot, path)

	_, err = os.Stat(sitemapsDir)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(sitemapsDir, 0755)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	output := filepath.Join(sitemapsDir, "sitemap.xml")
	s.path = output
	return s
}

func (s *sitemap) createSitemapFromLinksFile() ([]string, error) {
	linkFile, err := os.Open("sitemaps/links")
	if err != nil {
		return nil, err
	}
	defer linkFile.Close()

	data, err := io.ReadAll(linkFile)
	if err != nil {
		return nil, err
	}

	var links []string
	splitLinks := strings.Split(string(data), "\n")
	for i := range splitLinks {
		links = append(links, splitLinks[i])
	}

	return links, err
}

func (s *sitemap) getLastModifiedOrNow(url string) (string, error) {
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Last-Modified
	data, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer data.Body.Close()

	lastModified := data.Header["Last-Modified"]

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
