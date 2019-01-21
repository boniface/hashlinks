package services

import (
	"fmt"
	"hashlinks/domain"
	"testing"
)

func TestGetZones(t *testing.T) {
	links := GetZoneLinks("ZM")
	fmt.Println(" The Zones are ", links)
}

func TestGeLinksFromFeeds(t *testing.T) {
	feed := domain.Feed{"ZA", "NT", "1", "None", "http://feeds.news24.com/articles/news24/TopStories/rss", "RSS"}
	links := getLinks(feed)
	for _, link := range links {
		fmt.Println(" The Link is ", link.Linktitle)
	}
}
