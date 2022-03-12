package response

import (
	"go-drop-logistik/business/manifest"
	"time"
)

type Manifest struct {
	ID        int       `json:"id"`
	Code      string    `json:"code"`
	Status    string    `json:"status"`
	DriverID  int       `json:"driver_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ManifestPageResponse struct {
	Manifest *[]Manifest `json:"manifest"`
	Total    int         `json:"total"`
}

func FromDomain(manifestDomain manifest.Domain) Manifest {
	return Manifest{
		
	}
}

func FromListDomain(agentDomain []manifest.Domain, Count int) *ManifestPageResponse {
	allManifest := []Manifest{}
	for _, value := range agentDomain {
		agent := Manifest{
			ID:              value.ID,
			Code:            value.Code,
			Status:          value.Status,
			DriverID:        value.DriverID,
			CreatedAt:       value.CreatedAt,
			UpdatedAt:       value.UpdatedAt,
		}
		allManifest = append(allManifest, agent)
	}
	result := ManifestPageResponse{}
	result.Manifest = &allManifest
	result.Total = Count
	return &result
}
