package handler

import (
	"fmt"
	"net/http"

	"github.com/SawitProRecruitment/UserService/generated"
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
	return nil
}

func (s *Server) PostEstateEstateIdTree(ctx echo.Context, estateId string) error {
	return nil
}

func (s *Server) GetEstateEstateIdDronePlane(echo.Context, string, generated.GetEstateEstateIdDronePlaneParams) error {
	return nil
}

func (s *Server) GetEstateEstateIdStats(ctx echo.Context, estateId string) error {
	return nil
}
