package sitemap

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var defaultHTTPClient = &http.Client{Timeout: 15 * time.Second}

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

// Path sets the output directory (relative to the process working directory)
// and configures sitemap.xml as the output file.
func (s *sitemap) Path(dir string) (*sitemap, error) {
	outDir, err := ensureOutputDir(dir)
	if err != nil {
		return nil, err
	}
	s.path = filepath.Join(outDir, "sitemap.xml")
	return s, nil
}

// AddURL appends a URL to the sitemap. Pass an empty string to load URLs from
// sitemaps/links (one URL per line). Call Save to write the XML file.
//
// Google ignores ChangeFrequency and Priority:
// https://developers.google.com/search/docs/crawling-indexing/sitemaps/build-sitemap
func (s *sitemap) AddURL(url string) error {
	var urls []string
	var err error

	if url != "" {
		urls = []string{url}
	} else {
		urls, err = s.createSitemapFromLinksFile()
		if err != nil {
			return err
		}
	}

	for _, v := range urls {
		if strings.TrimSpace(v) == "" {
			continue
		}
		lastMod, merr := getLastModifiedOrNow(v)
		if merr != nil {
			return merr
		}
		s.URL = append(s.URL, URLs{
			Loc:     v,
			LastMod: lastMod,
		})
	}

	return nil
}

// Save writes the accumulated URLs to the sitemap XML file.
func (s *sitemap) Save() error {
	if s.path == "" {
		return errors.New("sitemap: call Path before Save")
	}
	if len(s.URL) == 0 {
		return errors.New("sitemap: no URLs to write")
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

	return nil
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
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			links = append(links, line)
		}
	}

	return links, nil
}

func getLastModifiedOrNow(url string) (string, error) {
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Last-Modified
	resp, err := defaultHTTPClient.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	_, _ = io.Copy(io.Discard, resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return time.Now().Format("2006-01-02"), nil
	}

	lastModified := resp.Header.Get("Last-Modified")
	if lastModified == "" {
		return time.Now().Format("2006-01-02"), nil
	}

	parseTime, err := time.Parse(time.RFC1123, lastModified)
	if err != nil {
		return "", fmt.Errorf("sitemap: parse Last-Modified for %q: %w", url, err)
	}

	return parseTime.Format("2006-01-02"), nil
}
