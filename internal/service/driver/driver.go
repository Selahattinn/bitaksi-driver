package driver

import (
	"errors"

	"github.com/Selahattinn/bitaksi-driver/internal/model"
	"github.com/Selahattinn/bitaksi-driver/internal/repository"
)

type Service struct {
	repository repository.Repository
}

var (
	ErrorUserNotFound = errors.New("driver not found")
)

func NewService(repo repository.Repository) (*Service, error) {
	return &Service{
		repository: repo,
	}, nil
}

func (s *Service) CreateDriver(driver *model.Driver) (int64, error) {
	//Create driver in db
	id, err := s.repository.GetDriverRepository().CreateDriver(driver)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (s *Service) GetDriver(id int64) (*model.Driver, error) {
	//Get driver in db
	d, err := s.repository.GetDriverRepository().GetDriver(id)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (s *Service) UpdateDriver(driver *model.Driver) (*model.Driver, error) {
	//Update driver in db
	d, err := s.repository.GetDriverRepository().UpdateDriver(driver)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (s *Service) DeleteDriver(driver *model.Driver) (int64, error) {
	//Delete driver in db
	id, err := s.repository.GetDriverRepository().DeleteDriver(driver)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (s *Service) GetAllDrivers() ([]*model.Driver, error) {
	//Get all drivers from db
	d, err := s.repository.GetDriverRepository().GetAllDrivers()
	if err != nil {
		return nil, err
	}
	return d, nil
}

func (s *Service) Shutdown() {
	s.repository.Shutdown()
}
