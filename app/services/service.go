package services

import "github.com/spadesk1991/skip-micro/app/dao"

type Service struct {
	*dao.Dao
}

func (s *Service) Ping() (err error) {
	err = s.Dao.Ping()
	return
}

func New() *Service {
	return &Service{dao.New()}
}
