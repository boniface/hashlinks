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
	fmt.Println(" Get FEED")
	feed, _ := fp.ParseURL(feedLink.Feedlink)
	fmt.Println(" Got the Feed")
	for i, feed := range feed.Items {
		fmt.Println(" get Feed ", i)
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

		fmt.Println(" Got the link", link.Linksite)

	}
	return links
}

func hashLink(url string) string {
	message := []byte(url)
	return fmt.Sprintf("a %s", md5.Sum(message))
}
