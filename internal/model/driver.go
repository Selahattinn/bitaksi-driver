package model

import (
	"errors"
)

var (
	ErrInvalidCoordinateType   = errors.New("invalid cordinate type")
	ErrInvalidCoordinatesArray = errors.New("invalid cordinate array")
	ErrInvalidLatitude         = errors.New("invalid latitude")
	ErrInvalidLongitude        = errors.New("invalid longitude")
)

type Driver struct {
	ID       int64   `bson:"_id" json:"id"`
	Location GeoJson `bson:"location" json:"location"`
}

// swagger:model
type GeoJson struct {
	Type string `json:"type"`

	// example: [-122.083739,37.423021]
	Coordinates []float64 `json:"coordinates"`
}

func NewDriver(id int64, lat float64, long float64) *Driver {
	if lat < -90 || lat > 90 {
		return nil
	}
	if long < -180 || long > 180 {
		return nil
	}
	return &Driver{
		ID: id,
		Location: GeoJson{
			Type:        "Point",
			Coordinates: []float64{lat, long},
		},
	}
}

func (d *Driver) Validate() error {
	if len(d.Location.Coordinates) != 2 {
		return ErrInvalidCoordinatesArray
	}
	if d.Location.Type != "Point" {
		return ErrInvalidCoordinateType
	}
	if d.Location.Coordinates[0] < -180 || d.Location.Coordinates[0] > 180 {
		return ErrInvalidLatitude
	}
	if d.Location.Coordinates[1] < -90 || d.Location.Coordinates[1] > 90 {
		return ErrInvalidLongitude
	}
	return nil
}
