package sitemap

const (
	xmlns      = "http://www.sitemaps.org/schemas/sitemap/0.9"
	xmlnsVideo = "http://www.google.com/schemas/sitemap-video/1.1"
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
