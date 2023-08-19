package service

func (s *pingService) PingService() string {
	err := s.pingStorage.Ping()
	if err != nil {
		return "error"
	}
	return "pong"
}
