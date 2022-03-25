package model

import (
	"testing"
)

func TestSearch_Validate(t *testing.T) {
	type fields struct {
		Rider       Rider
		MaxDistance float64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "test1", fields: fields{Rider: Rider{Lat: 40.730610, Long: -73.935242}, MaxDistance: 0.0}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Search{
				Rider:       tt.fields.Rider,
				MaxDistance: tt.fields.MaxDistance,
			}
			if err := s.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Search.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRider_Validate(t *testing.T) {
	type fields struct {
		Lat  float64
		Long float64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "test1", fields: fields{Lat: 40.730610, Long: -73.935242}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rider{
				Lat:  tt.fields.Lat,
				Long: tt.fields.Long,
			}
			if err := r.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Rider.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
