package repository

import (
	"github.com/scylladb/gocqlx"
	"github.com/scylladb/gocqlx/qb"
	"github.com/scylladb/gocqlx/table"
	"hashlinks/domain"
	"log"
)

var feedMetadata = table.Metadata{
	Name:    "feeds",
	Columns: []string{"zone", "sitecode", "id", "feedfilter", "feedlink", "feedtype"},
	PartKey: []string{"zone"},
	SortKey: []string{"sitecode","id"},
}


var feedTable = table.New(feedMetadata)


func addFeed(feed domain.Feed) domain.Feed {
	session := connect()
	defer session.Close()
	stmt, names := feedTable.Insert()
	q := gocqlx.Query(session.Query(stmt), names).BindStruct(feed)
	if err := q.ExecRelease(); err != nil {
		log.Fatal(err)
	}
	return feed

}


func getZoneFeeds(zone string) domain.Feeds {
	session := connect()
	defer session.Close()
	var feeds []domain.Feed
	statement, _ := qb.Select("feeds").Where(qb.Eq("zone")).ToCql()
	if err := gocqlx.Select(&feeds, session.Query(statement,zone)); err != nil {
		log.Fatal(err)
	}
	return feeds

}

func  getSiteFeeds(zone string, sitecode string) domain.Feeds {
	session := connect()
	defer session.Close()
	var feeds []domain.Feed
	statement, _ := qb.Select("feeds").Where(qb.Eq("zone")).Where(qb.Eq("sitecode")).ToCql()
	if err := gocqlx.Select(&feeds, session.Query(statement,zone,sitecode)); err != nil {
		log.Fatal(err)
	}
	return feeds

}

func getFeedById(zone string, sitecode string,id string) domain.Feed {
	session := connect()
	defer session.Close()
	var feed domain.Feed
	var feeds []domain.Feed
	statement, _ := qb.Select("feeds").Where(qb.Eq("zone"),qb.Eq("sitecode"),qb.Eq("id")).ToCql()
	if err := gocqlx.Select(&feeds, session.Query(statement,zone,sitecode,id)); err != nil {
		log.Fatal(err)
	}
	for _, feedOne := range feeds {
		feed = feedOne
	}
	return feed
}





