package search

import (
	"testing"

	"github.com/Selahattinn/bitaksi-driver/internal/model"
)

func Test_calculateDistance(t *testing.T) {
	type args struct {
		rider  *model.Rider
		driver *model.Driver
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "test1", args: args{rider: &model.Rider{Lat: 40.730610, Long: -73.935242}, driver: &model.Driver{Location: model.GeoJson{Type: "Point", Coordinates: []float64{40.730610, -73.935242}}}}, want: 0.0}, {name: "test2", args: args{rider: &model.Rider{Lat: 2.990353, Long: 101.533913}, driver: &model.Driver{Location: model.GeoJson{Type: "Point", Coordinates: []float64{2.960148, 101.577888}}}}, want: 5.926734214010531},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateDistance(tt.args.rider, tt.args.driver); got != tt.want {
				t.Errorf("calculateDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_degrees2radians(t *testing.T) {
	type args struct {
		degrees float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "test1", args: args{degrees: 0.0}, want: 0.0}, {name: "test2", args: args{degrees: 90.0}, want: 1.5707963267948966},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := degrees2radians(tt.args.degrees); got != tt.want {
				t.Errorf("degrees2radians() = %v, want %v", got, tt.want)
			}
		})
	}
}
