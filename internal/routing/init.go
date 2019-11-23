package routing

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	v1 "github.com/sergey-suslov/trechit-server/internal/routing/api/v1"
)

// Init routing
func Init() {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "☗ ${method} => ${uri} = ${status}\n",
	}))

	apiV1 := e.Group("/v1")
	v1.InitAPI(apiV1)

	e.Logger.Fatal(e.Start(":1323"))
}
