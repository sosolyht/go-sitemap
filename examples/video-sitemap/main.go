package main

import "github.com/sosolyht/go-sitemap/sitemap"

func main() {
	vs := sitemap.NewVideoSitemap().Path("sitemaps")

	videoURLs := []sitemap.VideoURL{
		{
			Loc: "https://www.example.com/videos/video1.html",
			Videos: []sitemap.Video{
				{
					ThumbnailLoc:         "https://www.example.com/thumbnail/thumbnail1.png",
					Title:                "example1",
					Description:          "example1 desc",
					ContentLoc:           "https://www.example.com",
					PlayerLoc:            "https://www.example.com",
					Duration:             nil,
					Rating:               nil,
					ViewCount:            nil,
					PublicationDate:      nil,
					ExpirationDate:       nil,
					FamilyFriendly:       nil,
					Restriction:          nil,
					Price:                nil,
					RequiresSubscription: nil,
					Uploader:             nil,
					Live:                 nil,
				},
			},
		},
		{
			Loc: "https://www.example.com/videos/video2.html",
			Videos: []sitemap.Video{
				{
					ThumbnailLoc:         "https://www.example.com/thumbnail/thumbnail2.png",
					Title:                "example2",
					Description:          "example2 desc",
					ContentLoc:           "https://www.example.com",
					PlayerLoc:            "https://www.example.com",
					Duration:             nil,
					Rating:               nil,
					ViewCount:            nil,
					PublicationDate:      nil,
					ExpirationDate:       nil,
					FamilyFriendly:       nil,
					Restriction:          nil,
					Price:                nil,
					RequiresSubscription: nil,
					Uploader:             nil,
					Live:                 nil,
				},
			},
		},
		{
			Loc: "https://www.example.com/videos/video3.html",
			Videos: []sitemap.Video{
				{
					ThumbnailLoc:         "https://www.example.com/thumbnail/thumbnail3.png",
					Title:                "example3",
					Description:          "example3 desc",
					ContentLoc:           "https://www.example.com",
					PlayerLoc:            "https://www.example.com",
					Duration:             nil,
					Rating:               nil,
					ViewCount:            nil,
					PublicationDate:      nil,
					ExpirationDate:       nil,
					FamilyFriendly:       nil,
					Restriction:          nil,
					Price:                nil,
					RequiresSubscription: nil,
					Uploader:             nil,
					Live:                 nil,
				},
			},
		},
	}

	for i := range videoURLs {
		vs.AddVideoURL(videoURLs[i])
	}
}
