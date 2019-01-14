package services

import (
	"fmt"
	"github.com/mmcdole/gofeed"
	"hashlinks/domain"
	"hashlinks/repository"
	"time"
)

func GentZoneLinks(zone string) domain.Links {

	feeds := repository.GetZoneFeeds(zone)

	for _, feed := range feeds {

	}

}

func getLinks(feedLink string) domain.Links {
	var links domain.Links
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(feedLink)

	for _, feed := range feed.Items {

		link := domain.Link{
			Zone:          feedLink,
			Datepublished: *feed.PublishedParsed,
			Linkhash:      "",
			Linkurl:       feed.Link,
			Linksite:      feed.Title,
			Linktitle:     feed.Title,
			Linktype:      feed.Content,
			Linksitecode:  "",
		}
		append(links)

	}
}
