package services

import (
	"hashlinks/domain"
	"hashlinks/repository"
	"sync"
)

func GetZones() domain.Zones {
	return repository.GetAllZones()
}

func ProcessLinks() {
	zones := GetZones()
	var wg sync.WaitGroup
	wg.Add(len(zones))
	for _, zone := range zones {
		links := GetZoneLinks(zone.Code)
		go saveLinks(links, &wg)
	}
	defer wg.Wait()
}

func saveLinks(links domain.Links, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, link := range links {
		repository.AddLink(link)

	}

}
