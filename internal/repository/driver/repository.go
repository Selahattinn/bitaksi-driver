package driver

import "github.com/Selahattinn/bitaksi-driver/internal/model"

type Reader interface {
	GetDriver(id int64) (*model.Driver, error)
}

type Writer interface {
	CreateDriver(user *model.Driver) (interface{}, error)
}

//Repository repository interface
type Repository interface {
	Reader
	Writer
}
