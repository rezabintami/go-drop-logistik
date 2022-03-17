package response

import (
	"go-drop-logistik/business/manifest"
	receiptResp "go-drop-logistik/controllers/receipts/response"
	"time"
)

type Manifest struct {
	ID        int                     `json:"id"`
	Code      string                  `json:"code"`
	Status    string                  `json:"status"`
	Receipt   *[]receiptResp.Receipts `json:"receipts"`
	DriverID  int                     `json:"driver_id"`
	CreatedAt time.Time               `json:"created_at"`
	UpdatedAt time.Time               `json:"updated_at"`
}

type ManifestResponse struct {
	ID        int       `json:"id"`
	Code      string    `json:"code"`
	Status    string    `json:"status"`
	DriverID  int       `json:"driver_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ManifestPageResponse struct {
	Manifest *[]ManifestResponse `json:"manifest"`
	Total    int                 `json:"total"`
}

func FromDomain(manifestDomain *manifest.Domain) (res *Manifest) {
	if manifestDomain != nil {
		res = &Manifest{
			ID:        manifestDomain.ID,
			Code:      manifestDomain.Code,
			Status:    manifestDomain.Status,
			Receipt:   receiptResp.FromManifestListDomain(&manifestDomain.Receipt),
			DriverID:  manifestDomain.DriverID,
			CreatedAt: manifestDomain.CreatedAt,
			UpdatedAt: manifestDomain.UpdatedAt,
		}
	}
	return res
}

func FromListDomain(manifestDomain []manifest.Domain, Count int) *ManifestPageResponse {
	allManifest := []ManifestResponse{}
	for _, value := range manifestDomain {
		manifest := ManifestResponse{
			ID:        value.ID,
			Code:      value.Code,
			Status:    value.Status,
			DriverID:  value.DriverID,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
		}
		allManifest = append(allManifest, manifest)
	}
	result := ManifestPageResponse{}
	result.Manifest = &allManifest
	result.Total = Count
	return &result
}
