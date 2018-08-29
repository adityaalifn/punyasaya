package repository

import (
	"time"

	"github.com/tokopedia/sqlt"
)

type postgreArticleRepo struct {
	DbMaster *sqlt.DB
	DbSlave  *sqlt.DB
	Timeout  time.Duration
}

func NewArticleRepository(dbMaster *sqlt.DB, dbSlave *sqlt.DB, timeout time.Duration) *postgreArticleRepo {
	return &postgreArticleRepo{
		DbMaster: dbMaster,
		DbSlave:  dbSlave,
		Timeout:  timeout,
	}
}

