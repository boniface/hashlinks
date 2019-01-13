package repository

import (
	"fmt"
	"hashlinks/domain"
	"testing"
)

func TestGetFeed(t *testing.T) {
	feed:=getSiteFeeds("ZM","LT")
	fmt.Println(" The FEEDS is  ", feed)
}

func TestGetZoneFeed(t *testing.T) {
	feed:=getZoneFeeds("ZM")
	fmt.Println(" The FEEDS is  ", feed)
}

func TestGetFeedById(t *testing.T) {
	feed:=getFeedById("ZM","LT","1")
	fmt.Println(" The FEED is  ", feed)
}

func TestAddFeed(t *testing.T) {
	feed:=domain.Feed{"ZA","NT","1","None","http://feeds.news24.com/articles/news24/TopStories/rss","RSS"}
	results:=addFeed(feed)
	fmt.Println(" The FEEDS is  ", results)
}


