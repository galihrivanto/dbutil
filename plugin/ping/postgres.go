package ping

import (
	"github.com/galihrivanto/dbutil"
	"github.com/galihrivanto/dbutil/util"
)

type postgres struct{}

func (p *postgres) Ping(config dbutil.Config) error {
	db, err := util.PgDial(config)
	if err != nil {
		return err
	}

	return db.Ping()
}

func init() {
	dbSupports["postgres"] = &postgres{}
}
