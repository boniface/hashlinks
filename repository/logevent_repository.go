package repository

import "github.com/scylladb/gocqlx/table"

var logMetadata = table.Metadata{
	Name:    "feeds",
	Columns: []string{"zone", "sitecode", "id", "feedfilter", "feedlink", "feedtype"},
	PartKey: []string{"zone"},
	SortKey: []string{"sitecode","id"},
}


var logTable = table.New(logMetadata)


