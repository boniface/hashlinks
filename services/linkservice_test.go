package services

import (
	"fmt"
	"hashlinks/domain"
	"testing"
)

func TestGetZones(t *testing.T) {
	links := GetZoneLinks("ZM")
	//for i, link:=range links{
	//	fmt.Println(i, "Site:  ", link.Linksitecode, "Link:  ", link.Linkurl)
	//}

	fmt.Println(" The Link Size is ", len(links))
}

func TestGeLinksFromFeeds(t *testing.T) {
	var links domain.Links
	feed := domain.Feed{"ZA", "NT", "1", "None", "https://diggers.news/feed/", "RSS"}
	feed1 := domain.Feed{"ZA", "NT", "1", "None", "https://www.daily-mail.co.zm/feed/", "RSS"}
	feed2 := domain.Feed{"ZA", "NT", "1", "None", "https://www.dailynation.info/feed/", "RSS"}
	feed3 := domain.Feed{"ZA", "NT", "1", "None", "https://www.lusakatimes.com/feed/", "RSS"}
	feed4 := domain.Feed{"ZA", "NT", "1", "None", "https://www.themastonline.com/feed/", "RSS"}
	feed5 := domain.Feed{"ZA", "NT", "1", "None", "http://www.muvitv.com/feed/", "RSS"}
	feed6 := domain.Feed{"ZA", "NT", "1", "None", "http://www.qfmzambia.com/feed/", "RSS"}
	feed7 := domain.Feed{"ZA", "NT", "1", "None", "https://www.tumfweko.com/feed/", "RSS"}
	feed8 := domain.Feed{"ZA", "NT", "1", "None", "http://www.times.co.zm/?feed=rss2", "RSS"}
	feed9 := domain.Feed{"ZA", "NT", "1", "None", "https://www.znbc.co.zm/feed/", "RSS"}
	feed10 := domain.Feed{"ZA", "NT", "1", "None", "https://www.zambiawatchdog.com/feed/", "RSS"}
	feed11 := domain.Feed{"ZA", "NT", "1", "None", "http://feeds.news24.com/articles/news24/TopStories/rss", "RSS"}
	feeds := []domain.Feed{feed, feed1, feed2, feed3, feed4, feed5, feed6, feed7, feed8, feed9, feed10, feed11}
	//for _, fed := range feeds {
	//	links = append(links)
	//}

	fmt.Println(" The Size of the Array is ", len(feeds), links)

}
