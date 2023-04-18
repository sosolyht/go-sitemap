package sitemap

import (
	"encoding/xml"
	"os"
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
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod"`
	// Google ignores ChangeFrequency and Priority
	// https://developers.google.com/search/docs/crawling-indexing/sitemaps/build-sitemap
	ChangeFreq ChangeFrequency `xml:"changefreq"`
	Priority   float32         `xml:"priority"`
}

func NewURL() *Sitemap {
	return &Sitemap{
		Xmlns: XMLNS,
	}
}

func (s *Sitemap) AddURL(url URLs) error {
	url.LastMod = time.Now().Format("2006-01-02")
	s.URL = append(s.URL, url)

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

	return nil
}

func (s *Sitemap) FrequencyAlways() {
	var url URLs
	url.ChangeFreq = ALWAYS
}

func (s *Sitemap) FrequencyHourly() {
	var url URLs
	url.ChangeFreq = HOURLY
}

func (s *Sitemap) FrequencyDaily() {
	var url URLs
	url.ChangeFreq = DAILY
}

func (s *Sitemap) FrequencyWeekly() {
	var url URLs
	url.ChangeFreq = WEEKLY
}

func (s *Sitemap) FrequencyMonthly() {
	var url URLs
	url.ChangeFreq = MONTHLY
}

func (s *Sitemap) FrequencyYearly() {
	var url URLs
	url.ChangeFreq = YEARLY
}

func (s *Sitemap) FrequencyNever() {
	var url URLs
	url.ChangeFreq = NEVER
}
