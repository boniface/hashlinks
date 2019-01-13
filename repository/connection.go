package repository

import "github.com/gocql/gocql"

func connect() *gocql.Session {
	// connect to the cluster
	cluster := gocql.NewCluster("172.17.0.2")
	cluster.Keyspace = "hashdata"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	return session
}
