package tracks

import (
	"net/http"
	"strconv"

	"go-drop-logistik/business/tracks"
	"go-drop-logistik/controllers/tracks/request"
	"go-drop-logistik/controllers/tracks/response"
	base_response "go-drop-logistik/helper/response"

	echo "github.com/labstack/echo/v4"
)

type TracksController struct {
	trackUsecase        tracks.Usecase
}

func NewTracksController(uc tracks.Usecase) *TracksController {
	return &TracksController{
		trackUsecase:        uc,
	}
}

func (controller *TracksController) CreateTrack(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Track{}
	if err := c.Bind(&req); err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := controller.trackUsecase.StoreTrack(ctx, req.ToDomain())
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return base_response.NewSuccessInsertResponse(c, "Successfully inserted")
}

func (controller *TracksController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, _ := strconv.Atoi(c.Param("id"))

	manifest, err := controller.trackUsecase.GetByID(ctx, id)
	if err != nil {
		return base_response.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return base_response.NewSuccessResponse(c, response.FromDomain(manifest))
}
