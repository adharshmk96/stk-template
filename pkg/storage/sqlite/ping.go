package sqlite

import "github.com/adharshmk96/stk-template/pkg/core/serr"

func (s *sqliteRepo) Ping() error {
	err := s.conn.Ping()
	if err != nil {
		return serr.ErrPingFailed
	}
	return nil
}
