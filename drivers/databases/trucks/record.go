package trucks

type Trucks struct {
	ID           int `gorm:"primary_key"`
	Name         string
	Type         string
	LicensePlate string
}