package model

import (
	"reflect"
	"testing"
)

func TestNewDriver(t *testing.T) {
	type args struct {
		lat  float64
		long float64
		ID   int64
	}
	tests := []struct {
		name string
		args args
		want *Driver
	}{
		{name: "valid Driver", args: args{ID: 12, lat: -180, long: -90}, want: &Driver{ID: 12, Lat: -180, Long: -90}}, {name: "valid Driver", args: args{ID: 12, lat: 180, long: -90}, want: &Driver{ID: 12, Lat: 180, Long: -90}}, {name: "valid Driver", args: args{ID: 12, lat: 22, long: -22}, want: &Driver{ID: 12, Lat: 22, Long: -22}}, {name: "invalid Driver", args: args{ID: 12, lat: -181, long: 91}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDriver(tt.args.ID, tt.args.lat, tt.args.long); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDriver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDriver_Validate(t *testing.T) {
	type fields struct {
		Lat  float64
		Long float64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{name: "valid Driver", fields: fields{Lat: -180, Long: -90}, wantErr: false}, {name: "valid Driver", fields: fields{Lat: 180, Long: -90}, wantErr: false}, {name: "valid Driver", fields: fields{Lat: 22, Long: -22}, wantErr: false}, {name: "invalid Driver", fields: fields{Lat: -181, Long: 91}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Driver{
				Lat:  tt.fields.Lat,
				Long: tt.fields.Long,
			}
			if err := d.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Driver.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
