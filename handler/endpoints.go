package handler

import (
	"errors"
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/types"
	"github.com/labstack/echo/v4"
)

const errorBindRequest = "failed to bind request"
const errorServerResponse = "internal server error"

func (s *Server) PostEstate(ctx echo.Context) error {
	var req generated.CreateEstateRequestPayload

	if err := ctx.Bind(&req); err != nil {
		var badRequestResp generated.BadRequest
		badRequestResp.Error = errorBindRequest

		return ctx.JSON(http.StatusBadRequest, badRequestResp)
	}

	if err := s.Validator.Validate(&req); err != nil {
		var badRequestResp generated.BadRequest
		badRequestResp.Error = err.Error()

		return ctx.JSON(http.StatusBadRequest, badRequestResp)
	}

	id, err := s.EstateService.CreateNewEstate(ctx.Request().Context(), req.Width, req.Length)
	if err != nil {
		var internalServerErrorResp generated.InternalServerError
		internalServerErrorResp.Error = errorServerResponse

		return ctx.JSON(http.StatusInternalServerError, internalServerErrorResp)
	}

	var resp generated.CreateValidResponse
	resp.Id = id
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) PostEstateEstateIdTree(ctx echo.Context, estateId string) error {
	var req generated.CreateTreeRequestPayload
	type Path struct {
		EstateID string `validate:"uuid"`
	}

	// Validate UUID Format
	if err := s.Validator.Validate(&Path{EstateID: estateId}); err != nil {
		var badRequestResp generated.BadRequest
		badRequestResp.Error = err.Error()

		return ctx.JSON(http.StatusBadRequest, badRequestResp)
	}

	if err := ctx.Bind(&req); err != nil {
		var badRequestResp generated.BadRequest
		badRequestResp.Error = errorBindRequest

		return ctx.JSON(http.StatusBadRequest, badRequestResp)
	}

	if err := s.Validator.Validate(&req); err != nil {
		var badRequestResp generated.BadRequest
		badRequestResp.Error = err.Error()

		return ctx.JSON(http.StatusBadRequest, badRequestResp)
	}

	var resp generated.CreateValidResponse
	if id, err := s.TreeService.CreateNewTree(ctx.Request().Context(), estateId, req.X, req.Y, req.Height); err != nil {
		// Handling error
		switch {
		// Handling error if estate not found
		case errors.Is(err, types.ErrorEstateNotFound):
			var notFoundResp generated.ErrorNotFound
			notFoundResp.Error = err.Error()

			return ctx.JSON(http.StatusNotFound, notFoundResp)

		case
			// Handling error if the tree is out of bound
			errors.Is(err, types.ErorrTreeOutOfBound),
			// Handling error if the tree is already planted
			errors.Is(err, types.ErrorTreeAlreadyPlanted):
			var badRequestResp generated.BadRequest
			badRequestResp.Error = err.Error()

			return ctx.JSON(http.StatusBadRequest, badRequestResp)

		default:
			var internalServerErrorResp generated.InternalServerError
			internalServerErrorResp.Error = errorServerResponse

			return ctx.JSON(http.StatusInternalServerError, internalServerErrorResp)
		}
	} else {
		resp.Id = id
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) GetEstateEstateIdDronePlan(ctx echo.Context, estateID string, params generated.GetEstateEstateIdDronePlanParams) error {
	type Path struct {
		EstateID string `validate:"uuid"`
	}

	if err := s.Validator.Validate(&Path{EstateID: estateID}); err != nil {
		var badRequestResp generated.BadRequest
		badRequestResp.Error = err.Error()

		return ctx.JSON(http.StatusBadRequest, badRequestResp)
	}

	estate, err := s.EstateService.RetrieveEstate(ctx.Request().Context(), estateID)
	if err != nil {
		var internalServerErrorResp generated.InternalServerError
		internalServerErrorResp.Error = errorServerResponse

		return ctx.JSON(http.StatusInternalServerError, internalServerErrorResp)
	}

	trees, err := s.TreeService.RetrievesByEstateID(ctx.Request().Context(), estateID)
	if err != nil {
		var internalServerErrorResp generated.InternalServerError
		internalServerErrorResp.Error = errorServerResponse

		return ctx.JSON(http.StatusInternalServerError, internalServerErrorResp)
	}
	var resp generated.GetDronePlanResponse
	stats := s.DroneService.GetDronePlane(estate, trees)

	resp.Distance = stats
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) GetEstateEstateIdStats(ctx echo.Context, estateId string) error {
	type Path struct {
		EstateID string `validate:"uuid"`
	}

	if err := s.Validator.Validate(&Path{EstateID: estateId}); err != nil {
		var badRequestResp generated.BadRequest
		badRequestResp.Error = err.Error()

		return ctx.JSON(http.StatusBadRequest, badRequestResp)
	}

	var resp generated.GetStatEstateResponse
	if stats, err := s.TreeService.GetStats(ctx.Request().Context(), estateId); err != nil {
		// Return default value if error
		return ctx.JSON(http.StatusOK, resp)
	} else {
		resp.Count = stats.Count
		resp.Max = stats.Max
		resp.Min = stats.Min
		resp.Median = stats.Median
	}
	return ctx.JSON(http.StatusOK, resp)

}
