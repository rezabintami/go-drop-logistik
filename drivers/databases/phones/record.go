package phones

import "time"

type Phones struct {
	ID        int `gorm:"primary_key"`
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}