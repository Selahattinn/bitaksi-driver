package model

import (
	"reflect"
	"testing"
)

func TestNewDriver(t *testing.T) {
	type args struct {
		id   int64
		lat  float64
		long float64
	}
	tests := []struct {
		name string
		args args
		want *Driver
	}{
		{name: "valid", args: args{id: 1, lat: 1.1, long: 1.1}, want: &Driver{ID: 1, Location: GeoJson{Type: "Point", Coordinates: []float64{1.1, 1.1}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDriver(tt.args.id, tt.args.lat, tt.args.long); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDriver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDriver_Validate(t *testing.T) {
	type fields struct {
		ID       int64
		Location GeoJson
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "valid", fields: fields{ID: 1, Location: GeoJson{Type: "Point", Coordinates: []float64{1.1, 1.1}}}, wantErr: false}, {name: "invalid", fields: fields{ID: 1, Location: GeoJson{Type: "Point", Coordinates: []float64{1.1, -181.1}}}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Driver{
				ID:       tt.fields.ID,
				Location: tt.fields.Location,
			}
			if err := d.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Driver.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
