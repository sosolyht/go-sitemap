package sitemap

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type videoSitemap struct {
	XMLName    xml.Name `xml:"urlset"`
	Xmlns      string   `xml:"xmlns,attr"`
	XmlnsVideo string   `xml:"xmlns:video,attr"`
	URL        []VideoURL
	path       string
}

type VideoURL struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
	Videos  []Video  `xml:"video"`
}

type Video struct {
	ThumbnailLoc         string       `xml:"video:thumbnail_loc"`
	Title                string       `xml:"video:title"`
	Description          string       `xml:"video:description"`
	ContentLoc           string       `xml:"video:content_loc"`
	PlayerLoc            string       `xml:"video:player_loc"`
	Duration             *int         `xml:"video:duration,omitempty"`              // Optional
	Rating               *float64     `xml:"video:rating,omitempty"`                // Optional
	ViewCount            *int         `xml:"video:view_count,omitempty"`            // Optional
	PublicationDate      *time.Time   `xml:"video:publication_date,omitempty"`      // Optional
	ExpirationDate       *time.Time   `xml:"video:expiration_date,omitempty"`       // Optional
	FamilyFriendly       *string      `xml:"video:family_friendly,omitempty"`       // Optional
	Restriction          *Restriction `xml:"video:restriction,omitempty"`           // Optional
	Price                *Price       `xml:"video:price,omitempty"`                 // Optional
	RequiresSubscription *string      `xml:"video:requires_subscription,omitempty"` // Optional
	Uploader             *Uploader    `xml:"video:uploader,omitempty"`              // Optional
	Live                 *string      `xml:"video:live,omitempty"`                  // Optional
}

type Restriction struct {
	Relationship string `xml:"relationship,attr"`
	Country      string `xml:",chardata"`
}

type Price struct {
	Currency string  `xml:"currency,attr"`
	Amount   float64 `xml:",chardata"`
}

type Uploader struct {
	Info string `xml:"info,attr"`
	Name string `xml:",chardata"`
}

func NewVideoSitemap() *videoSitemap {
	return &videoSitemap{
		Xmlns:      "http://www.sitemaps.org/schemas/sitemap/0.9",
		XmlnsVideo: "http://www.google.com/schemas/sitemap-video/1.1",
	}
}

func (v *videoSitemap) AddVideoURL(url VideoURL) (err error) {
	v.URL = append(v.URL, url)
	xmlBytes, err := xml.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	sitemapFile, err := os.Create(v.path)
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

func (v *videoSitemap) Path(path string) *videoSitemap {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	projectRoot := fmt.Sprintf("%s/%s", filepath.Dir(currentDir), filepath.Base(currentDir))
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

	v.path = filepath.Join(sitemapsDir, "sitemap_video.xml")
	return v
}
