package model

import "errors"

var (
	ErrInvalidLatitude  = errors.New("invalid latitude")
	ErrInvalidLongitude = errors.New("invalid longitude")
)

type Driver struct {
	ID   int64   `json:"id"`
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

func NewDriver(id int64, lat float64, long float64) *Driver {
	if lat < -180 || lat > 180 {
		return nil
	}
	if long < -90 || long > 90 {
		return nil
	}

	return &Driver{
		Lat:  lat,
		Long: long,
		ID:   id,
	}
}

func (d *Driver) Validate() error {
	if d.Lat < -180 || d.Lat > 180 {
		return ErrInvalidLatitude
	}
	if d.Long < -90 || d.Long > 90 {
		return ErrInvalidLongitude
	}
	return nil
}
