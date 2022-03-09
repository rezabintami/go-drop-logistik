package manifest

import (
	"time"

	"gorm.io/gorm"
)

type Manifest struct {
	ID        int `gorm:"primary_key"`
	Code      string
	Status    string
	DriverID  int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}