package search

import (
	"math"

	"github.com/Selahattinn/bitaksi-driver/internal/model"
	"github.com/Selahattinn/bitaksi-driver/internal/repository"
)

const radius = 6371

type Service struct {
	repository repository.Repository
}

func NewService(repo repository.Repository) (*Service, error) {
	return &Service{
		repository: repo,
	}, nil
}

func (s *Service) FindSuitableDrivers(searchInfo *model.Search) ([]*model.SearchResult, error) {
	//Get suitable drivers in db
	d, err := s.repository.GetDriverRepository().FindSuitableDrivers(&searchInfo.Rider, searchInfo.MaxDistance*1000)
	if err != nil {
		return nil, err
	}
	var searchResults []*model.SearchResult

	for _, driver := range d {
		distance := calculateDistance(&searchInfo.Rider, driver)
		searchResult := &model.SearchResult{
			Driver:   *driver,
			Distance: distance,
		}
		searchResults = append(searchResults, searchResult)
	}

	return searchResults, nil
}

func calculateDistance(rider *model.Rider, driver *model.Driver) float64 {

	// calculate distance beetween two points
	degreesLat := degrees2radians(driver.Location.Coordinates[0] - rider.Lat)
	degreesLong := degrees2radians(driver.Location.Coordinates[1] - rider.Long)
	a := (math.Sin(degreesLat/2)*math.Sin(degreesLat/2) +
		math.Cos(degrees2radians(rider.Lat))*
			math.Cos(degrees2radians(driver.Location.Coordinates[0]))*math.Sin(degreesLong/2)*
			math.Sin(degreesLong/2))
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := radius * c
	return distance
}

func degrees2radians(degrees float64) float64 {
	return degrees * math.Pi / 180
}
func (s *Service) Shutdown() {
	s.repository.Shutdown()
}
