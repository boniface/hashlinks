package main

import (
	"fmt"
	"github.com/mmcdole/gofeed"

)

func main() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("https://zambiareports.com/feed/")
	fmt.Println("This is the Feeds For ", feed.FeedType)
	fmt.Println("This is the Feeds For ", feed.Categories)
	fmt.Println("The Author  ", feed.Author)
	fmt.Println("The Image ", feed.Image)
	fmt.Println("the Items  ", feed.Items)
	fmt.Println("Published  ", feed.Published)
	for i := 0; i <= len(feed.Items)-1; i++ {
		// display title content
		fmt.Println("Published  ", feed.Items[i].Title)
		// display content author & the date published
		fmt.Println("Published  ", feed.Items[i].Author.Name+" at "+feed.Items[i].Published)
		// display original link
		fmt.Println("Published  ", feed.Items[i].Link)

	}
}