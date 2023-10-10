package service

func (s *pingService) PingService() string {
	err := s.storage.Ping()
	if err != nil {
		return "error"
	}
	return "pong"
}
