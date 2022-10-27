package main

import (
	"Swagger/api"
	_ "Swagger/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	swagger "github.com/swaggo/echo-swagger"
)

// @title API測試
// @version 1.0
// @description API
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey  Bearer
// @in                          header
// @name                        Authorization
// @description					Description for what is this security definition being used

// @host localhost:1323
// @BasePath /api/v1
// @schemes http

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))
	//預設是index.html
	v1 := e.Group("/api/v1")
	{
		user := v1.Group("/user")
		{
			user.POST("", api.AddData)
			user.PUT("/:id", api.UpdateData)
			user.DELETE("/:id", api.DeleteData)
			user.GET("/:id", api.SearchData)
			user.GET("", api.ListData)
		}

	}
	e.GET("/swagger/*", swagger.WrapHandler)

	e.Logger.Fatal(e.Start("127.0.0.1:1323"))

}
