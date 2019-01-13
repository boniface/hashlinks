package repository

import (
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
	"github.com/scylladb/gocqlx/table"
	"hashlinks/domain"
	"log"
)

var zoneMetadata = table.Metadata{
	Name:    "zones",
	Columns: []string{"code", "logo", "name", "zonestatus"},
	PartKey: []string{"code"},
}

var zoneTable = table.New(zoneMetadata)


func addZone(zone domain.Zone) domain.Zone {
	session := connect()
	defer session.Close()
	stmt, names := zoneTable.Insert()
	q := gocqlx.Query(session.Query(stmt), names).BindStruct(zone)
	if err := q.ExecRelease(); err != nil {
		log.Fatal(err)
	}
	return zone
}

func getAllZones() domain.Zones {
	session := connect()
	defer session.Close()
	var zones []domain.Zone
	statement, _ := qb.Select("zones").ToCql()
	if err := gocqlx.Select(&zones, session.Query(statement)); err != nil {
		log.Fatal(err)
	}
	return zones
}

func getZone(code string) domain.Zone {
	session := connect()
	defer session.Close()
	var zones []domain.Zone
	var zone   domain.Zone

	statement, _ := qb.Select("zones").Where(qb.Eq("code")).ToCql()
	if err := gocqlx.Select(&zones, session.Query(statement,code)); err != nil {
		log.Fatal(err)
	}
	for _, zoneOne := range zones {
		zone = zoneOne
	}
	return zone
}

func deleteZone()  {

}


