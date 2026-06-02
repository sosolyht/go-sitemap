package sitemap

import (
	"encoding/xml"
	"errors"
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
	Duration             *int         `xml:"video:duration,omitempty"`
	Rating               *float64     `xml:"video:rating,omitempty"`
	ViewCount            *int         `xml:"video:view_count,omitempty"`
	PublicationDate      *time.Time   `xml:"video:publication_date,omitempty"`
	ExpirationDate       *time.Time   `xml:"video:expiration_date,omitempty"`
	FamilyFriendly       *string      `xml:"video:family_friendly,omitempty"`
	Restriction          *Restriction `xml:"video:restriction,omitempty"`
	Price                *Price       `xml:"video:price,omitempty"`
	RequiresSubscription *string      `xml:"video:requires_subscription,omitempty"`
	Uploader             *Uploader    `xml:"video:uploader,omitempty"`
	Live                 *string      `xml:"video:live,omitempty"`
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
		Xmlns:      xmlns,
		XmlnsVideo: xmlnsVideo,
	}
}

// Path sets the output directory (relative to the process working directory)
// and configures sitemap_video.xml as the output file.
func (v *videoSitemap) Path(dir string) (*videoSitemap, error) {
	outDir, err := ensureOutputDir(dir)
	if err != nil {
		return nil, err
	}
	v.path = filepath.Join(outDir, "sitemap_video.xml")
	return v, nil
}

// AddVideoURL appends a video URL entry. Call Save to write the XML file.
func (v *videoSitemap) AddVideoURL(url VideoURL) error {
	v.URL = append(v.URL, url)
	return nil
}

// Save writes the accumulated video URLs to the video sitemap XML file.
func (v *videoSitemap) Save() error {
	if v.path == "" {
		return errors.New("sitemap: call Path before Save")
	}
	if len(v.URL) == 0 {
		return errors.New("sitemap: no video URLs to write")
	}

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

	return nil
}
