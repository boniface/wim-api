package repository

import (
	"github.com/gocql/gocql"
)

func SelectKey() string {
	defer getCluster().Close()
	var key string
	err1 := Session.Query(`SELECT key FROM keychain LIMIT 1`).Consistency(gocql.One).Scan(&key)

	if err1 != nil {
		return ""
	}

	return key
}
