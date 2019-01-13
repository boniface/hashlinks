package repository

import (
	"fmt"
	"hashlinks/domain"
	"testing"
	"time"
)

func TestGetAllLinks(t *testing.T) {
	links:=getAllLinks()
	fmt.Println(" The Links are ", len(links))
}

func TestGetLinksPerZone(t *testing.T) {
	links:=getZoneLinks("ZM")
	fmt.Println(" The Links is ", len(links))
}

func TestGetLatestZone(t *testing.T) {
	links:=getLatestLinks("ZM")
	fmt.Println(" The Links is ", len(links))
}


func TestAddLink(t *testing.T) {
	link:= domain.Link{"ZA",
		time.Now(),
		"123",
		"https://www.bbc.com/sport/football/46766860",
		"www.bbc.com",
		"Home Team Wins ",
		"RSS",
		"BCC"}
	result:=addLink(link)
	fmt.Println(" The Zones is ", result)
}

