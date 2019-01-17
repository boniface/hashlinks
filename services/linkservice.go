package services

import (
	"crypto/md5"
	"fmt"
	"github.com/mmcdole/gofeed"
	"hashlinks/domain"
	"hashlinks/repository"
)

func GetZoneLinks(zone string) domain.Links {
	var links domain.Links
	feeds := repository.GetZoneFeeds(zone)
	for _, feed := range feeds {
		links = append(links, getLinks(feed)...)
	}
	return links
}

func getLinks(feedLink domain.Feed) domain.Links {
	var links domain.Links
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(feedLink.Feedlink)
	for _, feed := range feed.Items {
		link := domain.Link{
			Zone:          feedLink.Zone,
			Datepublished: *feed.PublishedParsed,
			Linkhash:      hashLink(feedLink.Feedlink),
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

func hashLink(url string) string {
	message := []byte(url)
	return fmt.Sprintf("a %s", md5.Sum(message))
}
