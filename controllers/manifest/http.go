package manifest

import (
	"net/http"
	"strconv"

	"go-drop-logistik/business/manifest"
	"go-drop-logistik/controllers/manifest/request"
	"go-drop-logistik/controllers/manifest/response"
	base_response "go-drop-logistik/helper/response"

	echo "github.com/labstack/echo/v4"
)

type ManifestController struct {
	manifestUsecase manifest.Usecase
}

func NewManifestController(uc manifest.Usecase) *ManifestController {
	return &ManifestController{
		manifestUsecase: uc,
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

	user, err := controller.manifestUsecase.GetByID(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromDomain(user))
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
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	err := controller.manifestUsecase.Delete(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, "Delete Successfully")
}
