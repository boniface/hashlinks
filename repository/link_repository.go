package repository

import (
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
	"github.com/scylladb/gocqlx/table"
	"hashlinks/domain"
	"log"
	"time"
)


var linkMetadata = table.Metadata{
	Name:    "links",
	Columns: []string{"zone", "datepublished", "linkhash", "linksite", "linksitecode", "linktitle","linktype","linkurl"},
	PartKey: []string{"zone"},
	SortKey: []string{"datepublished","linkhash"},
}


var linkTable = table.New(linkMetadata)



func addLink(link  domain.Link) domain.Link {
	session := connect()
	defer session.Close()
	stmt, names := linkTable.Insert()
	q := gocqlx.Query(session.Query(stmt), names).BindStruct(link)
	if err := q.ExecRelease(); err != nil {
		log.Fatal(err)
	}
	return link

}

func getZoneLinks(zone string) domain.Links {

	session := connect()
	defer session.Close()
	var links []domain.Link
	statement, _ := qb.Select("links").Where(qb.Eq("zone")).ToCql()
	if err := gocqlx.Select(&links, session.Query(statement,zone)); err != nil {
		log.Fatal(err)
	}
	return links
}

func getAllLinks() domain.Links {
	session := connect()
	defer session.Close()
	var links []domain.Link
	statement, _ := qb.Select("links").ToCql()
	if err := gocqlx.Select(&links, session.Query(statement)); err != nil {
		log.Fatal(err)
	}
	return links

}

func getLatestLinks(zone string) domain.Links {
	session := connect()
	defer session.Close()
	var links []domain.Link
	var time = time.Now().Add(-10*time.Hour)
	statement, _ := qb.Select("links").Where(qb.Eq("zone"),qb.Gt("datepublished")).ToCql()
	if err := gocqlx.Select(&links, session.Query(statement,zone,time)); err != nil {
		log.Fatal(err)
	}
	return links

}