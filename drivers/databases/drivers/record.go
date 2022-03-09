package drivers

type Drivers struct {
	ID      int `gorm:"primary_key"`
	Name    string
	Phone   string
	Address string
	TruckID int
}