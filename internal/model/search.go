package model

import "errors"

var (
	ErrInvalidMaxDistance = errors.New("invalid max distance")
)

type Search struct {
	Rider       Rider   `json:"rider"`
	MaxDistance float64 `json:"max_distance"` // in km
}

type Rider struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

func NewRider(lat, long float64) *Rider {
	if lat < -90 || lat > 90 {
		return nil
	}
	if long < -180 || long > 180 {
		return nil
	}
	return &Rider{
		Lat:  lat,
		Long: long,
	}
}

func (s *Search) Validate() error {
	if err := s.Rider.Validate(); err != nil {
		return err
	}
	if s.MaxDistance < 0 {
		return ErrInvalidMaxDistance
	}
	return nil
}

func (r *Rider) Validate() error {
	if r.Lat < -90 || r.Lat > 90 {
		return ErrInvalidLatitude
	}
	if r.Long < -180 || r.Long > 180 {
		return ErrInvalidLongitude
	}
	return nil
}
