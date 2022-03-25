package service

import (
	"github.com/Selahattinn/bitaksi-driver/internal/service/driver"
	"github.com/Selahattinn/bitaksi-driver/internal/service/search"
)

type Config struct{}

type Service interface {
	GetConfig() *Config
	GetDriverService() *driver.Service
	GetSearchService() *search.Service
	Shutdown()
}
