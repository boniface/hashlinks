package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://zambiareports.com/feed/")
	fmt.Println("This is the Feeds For ", feed.FeedType)

	for _, feed := range feed.Items {
		fmt.Println("The Link   ", feed.Link)
		//fmt.Println("Image URL  ", feed.Image.URL)
		fmt.Println("Date Publishged   ", feed.Published)
		fmt.Println("Published  ", feed.PublishedParsed)
	}

}
