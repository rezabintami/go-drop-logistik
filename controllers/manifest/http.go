package manifest

import (
	"net/http"
	"strconv"

	"go-drop-logistik/constants"
	"go-drop-logistik/controllers/manifest/request"
	"go-drop-logistik/controllers/manifest/response"
	base_response "go-drop-logistik/helpers"
	"go-drop-logistik/modules/manifest"
	"go-drop-logistik/modules/manifestreceipt"
	"go-drop-logistik/modules/trackmanifest"

	echo "github.com/labstack/echo/v4"
)

type ManifestController struct {
	manifestUsecase        manifest.Usecase
	manifestreceiptUsecase manifestreceipt.Usecase
	trackManifestUsecase   trackmanifest.Usecase
}

func NewManifestController(uc manifest.Usecase, mr manifestreceipt.Usecase, tr trackmanifest.Usecase) *ManifestController {
	return &ManifestController{
		manifestUsecase:        uc,
		manifestreceiptUsecase: mr,
		trackManifestUsecase:   tr,
	}
}

func (controller *ManifestController) CreateManifest(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Manifest{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := controller.manifestUsecase.StoreManifest(ctx, req.ToDomain())
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return base_response.NewSuccessInsertResponse(c, "Successfully inserted")
}

func (controller *ManifestController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, _ := strconv.Atoi(c.Param("id"))

	manifest, err := controller.manifestUsecase.GetByID(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	receipt, err := controller.manifestreceiptUsecase.GetAllByManifestID(ctx, manifest.ID)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	for _, value := range receipt {
		manifest.Receipt = append(manifest.Receipt, *value.Receipt)
	}

	tracks, err := controller.trackManifestUsecase.GetAllByManifestID(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	for _, value := range tracks {
		manifest.Tracks = append(manifest.Tracks, *value.Track)
	}

	return base_response.NewSuccessResponse(c, response.FromDomain(&manifest))
}

func (controller *ManifestController) Fetch(c echo.Context) error {
	ctx := c.Request().Context()

	page, _ := strconv.Atoi(c.QueryParam("page"))
	perpage, _ := strconv.Atoi(c.QueryParam("per_page"))

	manifest, count, err := controller.manifestUsecase.Fetch(ctx, page, perpage)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromListDomain(manifest, count))
}

func (controller *ManifestController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id, _ := strconv.Atoi(c.Param("id"))

	err := controller.manifestUsecase.Delete(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err = controller.manifestreceiptUsecase.DeleteByManifest(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err = controller.trackManifestUsecase.DeleteByManifest(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, "Delete Successfully")
}

func (controller *ManifestController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id, _ := strconv.Atoi(c.Param("id"))

	req := request.ManifestUpdate{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := controller.manifestUsecase.Update(ctx, req.ToDomain(), id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, "Update Successfully")
}

func (controller *ManifestController) UpdateStatus(c echo.Context) error {
	ctx := c.Request().Context()

	id, _ := strconv.Atoi(c.Param("id"))

	err := controller.manifestUsecase.Update(ctx, &manifest.Domain{Status: constants.SUCCESS}, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err = controller.manifestreceiptUsecase.UpdateStatusByManifest(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, "Update Successfully")
}
