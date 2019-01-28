package services

import (
	"crypto/md5"
	"fmt"
	"github.com/mmcdole/gofeed"
	"hashlinks/domain"
	"hashlinks/repository"
)

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func GetZoneLinks(zone string) domain.Links {
	linksChannel := make(chan domain.Links)
	var links domain.Links
	feeds := repository.GetZoneFeeds(zone)
	for _, feed := range feeds {
		go getLinks(feed, linksChannel)

	}

	return links
}

func getLinks(feedLink domain.Feed, linksChannel chan domain.Links) {
	defer recoverName()
	var links domain.Links
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(feedLink.Feedlink)

	if err != nil {
		fmt.Println("Feed: ", err)
	}

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
	linksChannel <- links
}

func hashLink(url string) string {
	message := []byte(url)
	return fmt.Sprintf("a %s", md5.Sum(message))
}
