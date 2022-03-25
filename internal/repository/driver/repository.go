package driver

import "github.com/Selahattinn/bitaksi-driver/internal/model"

type Reader interface {
	GetDriver(id int64) (*model.Driver, error)
	GetAllDrivers() ([]*model.Driver, error)
	FindSuitableDrivers(RiderPoint *model.Rider, distance float64) ([]*model.Driver, error)
}

type Writer interface {
	CreateDriver(user *model.Driver) (int64, error)
	UpdateDriver(user *model.Driver) (*model.Driver, error)
	DeleteDriver(user *model.Driver) (int64, error)
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}
