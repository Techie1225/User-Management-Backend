package main

import (
	"user-management-backend/internal/controller"
	"user-management-backend/internal/database"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	database.InitDB("mysql", "root:root@tcp(127.0.0.1:3306)/usermanage")

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:4200"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization},
	}))

	e.GET("/api/v1/users", controller.GetUsers)
	e.GET("/api/v1/users/:id", controller.GetUserByID)
	e.POST("/api/v1/users", controller.AddUser)
	e.PUT("/api/v1/users/:id", controller.UpdateUser)
	e.DELETE("/api/v1/users/:id", controller.DeleteUser)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Start(":8086")
}
