package sitemap

type URL interface {
	WithLoc(loc string) URL
	WithChangeFreq(freq ChangeFrequency) URL
	Do() *URLs
}
