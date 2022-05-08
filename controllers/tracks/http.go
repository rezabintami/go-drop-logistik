package tracks

import (
	"net/http"
	"strconv"

	"go-drop-logistik/app/middleware"
	"go-drop-logistik/constants"
	"go-drop-logistik/controllers/tracks/request"
	helpers "go-drop-logistik/helpers"
	"go-drop-logistik/modules/manifest"
	"go-drop-logistik/modules/manifestreceipt"
	"go-drop-logistik/modules/trackmanifest"
	"go-drop-logistik/modules/tracks"

	echo "github.com/labstack/echo/v4"
)

type TracksController struct {
	trackUsecase           tracks.Usecase
	trackManifestUsecase   trackmanifest.Usecase
	manifestreceiptUsecase manifestreceipt.Usecase
	manifestUsecase        manifest.Usecase
}

func NewTracksController(uc tracks.Usecase, tr trackmanifest.Usecase, mr manifestreceipt.Usecase, ma manifest.Usecase) *TracksController {
	return &TracksController{
		trackUsecase:           uc,
		trackManifestUsecase:   tr,
		manifestreceiptUsecase: mr,
		manifestUsecase:        ma,
	}
}

//! NEED GOROUTINES WHEN STORE TRACK IS CALLED AND CHANGES STATUS RECEIPTS TO SHIPPING
func (controller *TracksController) CreateTrack(c echo.Context) error {
	ctx := c.Request().Context()

	name := middleware.GetUser(c).Name
	manifestId, _ := strconv.Atoi(c.Param("manifestId"))

	req := request.Track{}
	if err := c.Bind(&req); err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	validateMessage, validate, err := helpers.Validate(&req)

	if validate {
		return helpers.ErrorValidateResponse(c, http.StatusBadRequest, err, validateMessage)
	}

	trackId, err := controller.trackUsecase.StoreTrack(ctx, req.ToDomain(), name)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	err = controller.trackManifestUsecase.Store(ctx, manifestId, trackId)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	if req.Message != constants.TRACKING_SHIPPING_MESSAGE {
		err = controller.manifestreceiptUsecase.UpdateStatusByManifest(ctx, constants.SHIPPING, manifestId)
		if err != nil {
			return helpers.ErrorResponse(c, http.StatusBadRequest, err)
		}

		err = controller.manifestUsecase.Update(ctx, &manifest.Domain{Status: constants.SHIPPING}, manifestId)
		if err != nil {
			return helpers.ErrorResponse(c, http.StatusBadRequest, err)
		}
	}

	return helpers.SuccessResponse(c, http.StatusCreated, nil)
}

func (controller *TracksController) DeleteTrack(c echo.Context) error {
	ctx := c.Request().Context()

	manifestId, _ := strconv.Atoi(c.Param("manifestId"))
	trackId, _ := strconv.Atoi(c.Param("trackId"))

	err := controller.trackManifestUsecase.Delete(ctx, manifestId, trackId)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	err = controller.trackUsecase.Delete(ctx, trackId)
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, nil)
}

func (controller *TracksController) UpdateTrack(c echo.Context) error {
	ctx := c.Request().Context()

	agentName := middleware.GetUser(c).Name
	trackId, _ := strconv.Atoi(c.Param("trackId"))

	req := request.Track{}
	if err := c.Bind(&req); err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	// validateMessage, validate, err := helpers.Validate(&req)

	// if validate {
	// 	return helpers.ErrorValidateResponse(c, http.StatusBadRequest, err, validateMessage)
	// }

	err := controller.trackUsecase.Update(ctx, trackId, agentName, req.ToDomain())
	if err != nil {
		return helpers.ErrorResponse(c, http.StatusBadRequest, err)
	}

	return helpers.SuccessResponse(c, http.StatusOK, nil)
}
