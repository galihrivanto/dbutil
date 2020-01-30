package util

import (
	"fmt"

	"github.com/galihrivanto/dbutil"
	"github.com/globalsign/mgo"
)

func MongoDial(config dbutil.Config) (*mgo.Session, error) {
	dsn := fmt.Sprintf("mongodb://%s:%d/%s",
		config.GetString("host", "localhost"),
		config.GetInt("port", 27017),
		config.GetString("dbname"),
	)

	session, err := mgo.Dial(dsn)
	if err != nil {
		return nil, err
	}

	// TODO: add session setting

	return session, nil
}
