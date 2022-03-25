package service

import (
	"github.com/Selahattinn/bitaksi-driver/internal/repository"
	"github.com/Selahattinn/bitaksi-driver/internal/service/driver"
	"github.com/Selahattinn/bitaksi-driver/internal/service/search"
)

type Provider struct {
	cfg           *Config
	repository    repository.Repository
	DriverService *driver.Service
	SearchService *search.Service
}

func NewProvider(cfg *Config, repo repository.Repository) (*Provider, error) {
	driverService, err := driver.NewService(repo)
	if err != nil {
		return nil, err
	}
	searchService, err := search.NewService(repo)
	if err != nil {
		return nil, err
	}
	return &Provider{
		cfg:           cfg,
		repository:    repo,
		DriverService: driverService,
		SearchService: searchService,
	}, nil
}

func (p *Provider) GetConfig() *Config {
	return p.cfg
}

func (p *Provider) Shutdown() {
	p.repository.Shutdown()
}

func (p *Provider) GetDriverService() *driver.Service {
	return p.DriverService
}

func (p *Provider) GetSearchService() *search.Service {
	return p.SearchService
}
