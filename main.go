package main

import (
	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/config"
	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/factory"
	"github.com/GP-2-Kelompok-4/Immersive-Dashboard-App/utils/database/mysql"

	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	mysql.MigrateDB(db)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	factory.InitFactory(e, db)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))
}
