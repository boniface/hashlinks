package services

import (
	"fmt"
	"testing"
)

func TestGetZones(t *testing.T) {
	links := GetZoneLinks("ZM")
	fmt.Println(" The Zones are ", links)
}
