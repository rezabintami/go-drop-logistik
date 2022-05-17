package response

import (
	driverResp "go-drop-logistik/controllers/drivers/response"
	receiptResp "go-drop-logistik/controllers/receipts/response"
	trackResp "go-drop-logistik/controllers/tracks/response"
	"go-drop-logistik/modules/manifest"
)

type Manifest struct {
	ID      int                     `json:"id"`
	Code    string                  `json:"code"`
	Status  string                  `json:"status"`
	Receipt *[]receiptResp.Receipts `json:"receipts"`
	Driver  *driverResp.Drivers     `json:"driver"`
	Tracks  *[]trackResp.Track      `json:"track"`
}

type ManifestResponse struct {
	ID     int                 `json:"id"`
	Code   string              `json:"code"`
	Status string              `json:"status"`
	Driver *driverResp.Drivers `json:"driver"`
}

type ManifestPageResponse struct {
	Manifest *[]ManifestResponse `json:"manifest"`
	Total    int                 `json:"total"`
}

func FromDomain(manifestDomain *manifest.Domain) (res *Manifest) {
	if manifestDomain != nil {
		res = &Manifest{
			ID:      manifestDomain.ID,
			Code:    manifestDomain.Code,
			Status:  manifestDomain.Status,
			Receipt: receiptResp.FromManifestListDomain(&manifestDomain.Receipt),
			Driver:  driverResp.FromDomain(manifestDomain.Driver),
			Tracks:  trackResp.FromListDomain(&manifestDomain.Tracks),
		}
	}
	return res
}

func FromListDomain(manifestDomain []manifest.Domain, Count int) *ManifestPageResponse {
	allManifest := []ManifestResponse{}
	for _, value := range manifestDomain {
		manifest := ManifestResponse{
			ID:     value.ID,
			Code:   value.Code,
			Status: value.Status,
			Driver: driverResp.FromDomain(value.Driver),
		}
		allManifest = append(allManifest, manifest)
	}
	result := ManifestPageResponse{}
	result.Manifest = &allManifest
	result.Total = Count
	return &result
}
