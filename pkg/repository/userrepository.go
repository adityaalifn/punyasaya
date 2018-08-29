package repository

import (
	"time"

	"github.com/tokopedia/sqlt"
)

type postgreUserRepo struct {
	DbMaster *sqlt.DB
	DbSlave  *sqlt.DB
	Timeout  time.Duration
}

func NewUserRepository(dbMaster *sqlt.DB, dbSlave *sqlt.DB, timeout time.Duration) *postgreUserRepo {
	return &postgreUserRepo{
		DbMaster: dbMaster,
		DbSlave:  dbSlave,
		Timeout:  timeout,
	}
}
