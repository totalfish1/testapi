package webscraper

import (
	"github.com/gocolly/colly"
	"github.com/PuerkitoBio/goquery"
)

urls := [...]{
	"https://stackoverflow.com/jobs",
	"https://www.devex.com/jobs/search",
	"https://www.jrdevjobs.com/",
	"https://remoteok.io/remote-dev-jobs"
}

func GetJobListings(params string) []JobListing {
	c = colly.Coll
}