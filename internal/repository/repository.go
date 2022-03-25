package repository

import "github.com/Selahattinn/bitaksi-driver/internal/repository/driver"

type Repository interface {
	Shutdown()
	GetDriverRepository() driver.Repository
}
