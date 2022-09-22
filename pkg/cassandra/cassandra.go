package cassandra

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

type CassandraClient struct {
	Session  gocqlx.Session
	KeySpace string
}

func NewCassandraClient(clusterUrls []string, keyspace string) (*CassandraClient, error) {

	cluster := gocql.NewCluster(clusterUrls...)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Any

	session, err := gocqlx.WrapSession(cluster.CreateSession())

	if err != nil {
		return nil, err
	}

	return &CassandraClient{
		Session:  session,
		KeySpace: keyspace,
	}, nil
}

func (c *CassandraClient) CloseConnection() {
	c.Session.Close()
}

type CassandraRepository struct {
	CassandraClient CassandraClient
}
