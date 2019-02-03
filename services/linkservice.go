package services

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/mmcdole/gofeed"
	"hashlinks/domain"
	"hashlinks/repository"
	"sync"
)

var totalLinks domain.Links
var lock sync.Mutex

func GetZoneLinks(zone string) domain.Links {
	var wg sync.WaitGroup
	feeds := repository.GetZoneFeeds(zone)
	wg.Add(len(feeds))
	for _, feed := range feeds {
		go getLinks(feed, &wg)
	}
	wg.Wait()
	return totalLinks
}

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from ", r)
	}
}

func updateLinks(links domain.Links) {
	lock.Lock()
	defer lock.Unlock()
	totalLinks = append(totalLinks, links...)

}

func getLinks(feedLink domain.Feed, wg *sync.WaitGroup) {
	defer recoverName()
	defer wg.Done()
	var feedLinks domain.Links
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(feedLink.Feedlink)

	if err != nil {
		fmt.Println("Feed: ", err)
	}
	for _, feed := range feed.Items {
		link := domain.Link{
			Zone:          feedLink.Zone,
			Datepublished: *feed.PublishedParsed,
			Linkhash:      hashLink(feed.Link),
			Linkurl:       feed.Link,
			Linksite:      feedLink.Feedlink,
			Linktitle:     feed.Title,
			Linktype:      feedLink.Feedtype,
			Linksitecode:  feedLink.Sitecode,
		}
		feedLinks = append(feedLinks, link)
	}
	updateLinks(feedLinks)
}

func hashLink(url string) string {
	message := md5.Sum([]byte(url))
	return hex.EncodeToString(message[:])

}
