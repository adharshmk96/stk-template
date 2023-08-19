package sqlite

import (
	"database/sql"

	"github.com/adharshmk96/stk-template/singlemod/internals/core"
)

type sqliteRepo struct {
	conn *sql.DB
}

func NewSqliteRepo(conn *sql.DB) core.PingStorage {
	return &sqliteRepo{
		conn: conn,
	}
}
