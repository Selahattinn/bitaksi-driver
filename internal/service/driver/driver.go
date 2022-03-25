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

func (s *Service) GetUser(driver *model.Driver) (*model.Driver, error) {
	//Get driver from db
	d, err := s.repository.GetDriverRepository().GetDriver(driver.ID)
	if err != nil {

		return nil, err
	}
	return d, nil
}

func (s *Service) CreateDriver(user *model.Driver) (interface{}, error) {
	//Create driver in db
	u, err := s.repository.GetDriverRepository().CreateDriver(user)
	if err != nil {
		return nil, err
	}
	return u, nil
}
