package tracks

import "time"

type Tracks struct {
	ID                 int `gorm:"primary_key"`
	StartAgentID       int
	CurrentAgentID     int
	DestinationAgentID int
	ManifestID         int
	Message            string
	CreatedAt          time.Time
}
