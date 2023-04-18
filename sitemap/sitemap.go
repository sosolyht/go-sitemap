package sitemap

import (
	"encoding/xml"
	"log"
	"os"
	"time"
)

const (
	XmlVersion   = `<?xml version="1.0" encoding="UTF-8"?>`
	SitemapIndex = `<sitemapindex xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9">`
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
	XMLVersion   string `xml:"xmlVersion,omitempty"`
	SitemapIndex string `xml:"sitemapindex,omitempty"`
	URL          []URL  `xml:"url,omitempty"`
}

type URL struct {
	Loc        string          `xml:"loc"`
	LastMod    string          `xml:"lastmod"`
	ChangeFreq ChangeFrequency `xml:"changefreq"`
	Priority   float32         `xml:"priority"`
}

func NewURL() *Sitemap {
	return &Sitemap{
		XMLVersion:   XmlVersion,
		SitemapIndex: SitemapIndex,
	}
}

func (s *Sitemap) AddURL(url URL) error {
	url.LastMod = time.Now().Format("2006-01-02")
	s.URL = append(s.URL, url)

	resp, err := xml.MarshalIndent(s.URL, "", "   ")
	if err != nil {
		return err
	}

	sm, err := os.Create("sitemaps/sitemap.xml")
	if err != nil {
		return err
	}

	defer func() {
		if err = sm.Close(); err != nil {
			log.Printf("error closing sitemap file: %v", err)
		}
	}()

	if _, err = sm.Write(resp); err != nil {
		return err
	}

	return nil
}

func (s *Sitemap) FrequencyAlways() {
	var url URL
	url.ChangeFreq = ALWAYS
}

func (s *Sitemap) FrequencyHourly() {
	var url URL
	url.ChangeFreq = HOURLY
}

func (s *Sitemap) FrequencyDaily() {
	var url URL
	url.ChangeFreq = DAILY
}

func (s *Sitemap) FrequencyWeekly() {
	var url URL
	url.ChangeFreq = WEEKLY
}

func (s *Sitemap) FrequencyMonthly() {
	var url URL
	url.ChangeFreq = MONTHLY
}

func (s *Sitemap) FrequencyYearly() {
	var url URL
	url.ChangeFreq = YEARLY
}

func (s *Sitemap) FrequencyNever() {
	var url URL
	url.ChangeFreq = NEVER
}
