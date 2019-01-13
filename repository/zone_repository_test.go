package repository

import (
	"fmt"
	"hashlinks/domain"
	"testing"
)

func TestGetZones(t *testing.T) {
	zone:=getAllZones()
	fmt.Println(" The Zones are ", zone)
}

func TestGetZone(t *testing.T) {
	zone:=getZone("SA")
	fmt.Println(" The Zones is ", zone)
}


func TestAZone(t *testing.T) {
	zone:= domain.Zone{"UK","United Kingdom","ACTIVE","NO"}
	result:=addZone(zone)
	fmt.Println(" The Zones is ", result)
}

