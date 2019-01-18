package services

import (
	"crypto/md5"
	"fmt"
	"github.com/mmcdole/gofeed"
	"hashlinks/domain"
	"hashlinks/repository"
	"time"
)

func GetZoneLinks(zone string) domain.Links {
	var links domain.Links
	feeds := repository.GetZoneFeeds(zone)
	for _, feed := range feeds {
		sets := getLinks(feed)
		//links = append(links, sets...)
		fmt.Println(" the Size is ", len(sets))
	}
	return links
}

func getLinks(feedLink domain.Feed) domain.Links {
	fmt.Println(" has this been HIT")
	var links domain.Links
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(feedLink.Feedlink)
	fmt.Println(" Got the Feed")
	for _, feed := range feed.Items {
		link := domain.Link{
			Zone:          feedLink.Zone,
			Datepublished: time.Now(),
			Linkhash:      hashLink(feedLink.Feedlink),
			Linkurl:       feed.Link,
			Linksite:      feedLink.Feedlink,
			Linktitle:     feed.Title,
			Linktype:      feedLink.Feedtype,
			Linksitecode:  feedLink.Sitecode,
		}

		fmt.Println(" Got the link", link.Linksite)

	}
	return links
}

func hashLink(url string) string {
	message := []byte(url)
	return fmt.Sprintf("a %s", md5.Sum(message))
}
