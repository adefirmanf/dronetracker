package main

import (
	"os"

	"github.com/SawitProRecruitment/UserService/generated"
	"github.com/SawitProRecruitment/UserService/handler"
	"github.com/SawitProRecruitment/UserService/repository"
	"github.com/SawitProRecruitment/UserService/service/drone"
	"github.com/SawitProRecruitment/UserService/service/estate"
	"github.com/SawitProRecruitment/UserService/service/tree"

	repoEstate "github.com/SawitProRecruitment/UserService/repository/estate"
	repoTree "github.com/SawitProRecruitment/UserService/repository/tree"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	var server generated.ServerInterface = newServer()

	generated.RegisterHandlers(e, server)
	e.Use(middleware.Logger())
	e.Logger.Fatal(e.Start(":1323"))
}

func newServer() *handler.Server {
	dbDsn := os.Getenv("DATABASE_URL")

	db := repository.NewRepository(repository.NewRepositoryOptions{
		Dsn: dbDsn,
	})

	// Initiate Estate Repository
	estateRepo := repoEstate.NewRepository(db)
	treeRepo := repoTree.NewRepository(db)

	opts := handler.NewServerOptions{
		Validator:     handler.NewValidator(),
		EstateService: estate.NewEstateService(estateRepo),
		TreeService:   tree.NewTreeService(treeRepo, estateRepo),
		DroneService:  drone.NewDroneService(drone.DroneOpts{}),
	}

	return handler.NewServer(opts)
}
