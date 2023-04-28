package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"school-manager/config"
	"school-manager/controller"
)

func main() {

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Connect To Database
	config.DatabaseInit()
	gorm := config.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	err = dbGorm.Ping()
	if err != nil {
		return
	}

	userRoute := e.Group("/login")
	userRoute.POST("/", controller.SignIn)

	studentRoute := e.Group("/students")
	studentRoute.GET("", controller.GetStudents)

	e.Logger.Fatal(e.Start(":9090"))
}
