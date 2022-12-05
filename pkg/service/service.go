package service

import (
	"ginapis/pkg/config"
	"ginapis/pkg/dao"
)

type Service struct {
	cfg *config.Config
	d   *dao.Dao
}

func New(cfg *config.Config, d *dao.Dao) *Service {
	svc := &Service{
		cfg: cfg,
		d:   d,
	}

	return svc
}

func (svc *Service) Close() {}
