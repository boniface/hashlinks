package services

import (
	"hashlinks/domain"
	"hashlinks/repository"
)

func GetZones() domain.Zones {
	return repository.GetAllZones()
}
