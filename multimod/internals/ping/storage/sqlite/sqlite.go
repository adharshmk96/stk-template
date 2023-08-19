package sqlite

import (
	"database/sql"

	"github.com/adharshmk96/stk-template/multimod/internals/ping/interfaces"
)

type sqliteRepo struct {
	conn *sql.DB
}

func NewSqliteRepo(conn *sql.DB) interfaces.PingStorage {
	return &sqliteRepo{
		conn: conn,
	}
}
