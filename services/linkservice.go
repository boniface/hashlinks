package services

import (
	"github.com/mmcdole/gofeed"
	"hashlinks/domain"
	"hashlinks/repository"
)

func GentZoneLinks(zone string) domain.Links {

	feeds := repository.GetZoneFeeds(zone)

	for _, feed := range feeds {

	}

}

func getLinks(feedLink domain.Feed) domain.Links {
	var links domain.Links
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(feedLink.Feedlink)
	for _, feed := range feed.Items {
		link := domain.Link{
			Zone:          feedLink.Zone,
			Datepublished: *feed.PublishedParsed,
			Linkhash:      "",
			Linkurl:       feed.Link,
			Linksite:      feedLink.Feedlink,
			Linktitle:     feed.Title,
			Linktype:      feedLink.Feedtype,
			Linksitecode:  feedLink.Sitecode,
		}
		links = append(links, link)

	}
	return links
}
