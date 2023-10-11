package pingStorage

import "github.com/adharshmk96/stk-template/singlemod/internals/core/serr"

// Repository Methods
func (s *sqliteRepo) Ping() error {
	err := s.conn.Ping()
	if err != nil {
		return serr.ErrPingFailed
	}
	return nil
}
