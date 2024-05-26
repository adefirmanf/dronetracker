package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/types"
	"github.com/labstack/echo/v4"
)

// This is just a test endpoint to get you started. Please delete this endpoint.
// (GET /hello)
func (s *Server) GetHello(ctx echo.Context, params generated.GetHelloParams) error {
	var resp generated.HelloResponse
	resp.Message = fmt.Sprintf("Hello User %d", params.Id)
	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) PostEstate(ctx echo.Context) error {
	var req generated.CreateEstateRequestPayload

	if err := ctx.Bind(&req); err != nil {
		var badRequestResp generated.BadRequest
		badRequestResp.Error = err.Error()

		return ctx.JSON(http.StatusBadRequest, badRequestResp)
	}

	if err := s.Validator.Validate(&req); err != nil {
		var badRequestResp generated.BadRequest
		badRequestResp.Error = err.Error()

		return ctx.JSON(http.StatusBadRequest, badRequestResp)
	}

	id, err := s.EstateService.CreateNewEstate(ctx.Request().Context(), req.Width, req.Length)
	if err != nil {
		return err
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
		badRequestResp.Error = err.Error()

		return ctx.JSON(http.StatusBadRequest, badRequestResp)
	}

	if err := s.Validator.Validate(&req); err != nil {
		var badRequestResp generated.BadRequest
		badRequestResp.Error = err.Error()

		return ctx.JSON(http.StatusBadRequest, badRequestResp)
	}

	var resp generated.CreateValidResponse
	if id, err := s.TreeService.CreateNewTree(ctx.Request().Context(), estateId, req.X, req.Y, req.Height); err != nil {
		log.Println(err)
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
			return ctx.JSON(http.StatusInternalServerError, "internal server error")
		}
	} else {
		resp.Id = id
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (s *Server) GetEstateEstateIdDronePlane(echo.Context, string, generated.GetEstateEstateIdDronePlaneParams) error {
	return nil
}

func (s *Server) GetEstateEstateIdStats(ctx echo.Context, estateId string) error {
	return nil
}
