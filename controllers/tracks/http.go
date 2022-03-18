package tracks

import (
	"net/http"
	"strconv"

	"go-drop-logistik/app/middleware"
	"go-drop-logistik/business/trackmanifest"
	"go-drop-logistik/business/tracks"
	"go-drop-logistik/controllers/tracks/request"
	base_response "go-drop-logistik/helper/response"

	echo "github.com/labstack/echo/v4"
)

type TracksController struct {
	trackUsecase         tracks.Usecase
	trackManifestUsecase trackmanifest.Usecase
}

func NewTracksController(uc tracks.Usecase, tr trackmanifest.Usecase) *TracksController {
	return &TracksController{
		trackUsecase:         uc,
		trackManifestUsecase: tr,
	}
}

func (controller *TracksController) CreateTrack(c echo.Context) error {
	ctx := c.Request().Context()

	name := middleware.GetUser(c).Name
	manifestId, _ := strconv.Atoi(c.Param("id"))

	req := request.Track{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	trackId, err := controller.trackUsecase.StoreTrack(ctx, req.ToDomain(), name)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err = controller.trackManifestUsecase.Store(ctx, manifestId, trackId)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessInsertResponse(c, "Successfully inserted")
}
