package main

import (
	"Gorm/database"
	"Gorm/model"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()
	conn := database.GetDB()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}))
	userGroup := e.Group("/gorm")
	userGroup.POST("", func(c echo.Context) error {
		var user []model.User
		if err := c.Bind(&user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		//conn.Debug().Create(&user);

		if result := conn.Debug().CreateInBatches(user, 5); result.Error != nil {
			return echo.NewHTTPError(http.StatusBadRequest, result.Error)
		}
		return c.JSON(http.StatusOK, user)
	})

	userGroup.GET("", func(c echo.Context) error {
		var input struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
		}
		if err := echo.FormFieldBinder(c).Int64("id", &input.ID).String("name", &input.Name).BindError(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}

		user := &model.User{}
		users, err := user.List(model.User{
			ID:   input.ID,
			Name: input.Name,
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		fmt.Println(users)
		return c.JSON(http.StatusOK, users)
	})
	userGroup.DELETE("", func(c echo.Context) error {
		var input int64
		if err := echo.FormFieldBinder(c).Int64("id", &input).BindError(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		user := &model.User{}
		err := user.Delete(input)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, "isok")
	})
	userGroup.PUT("", func(c echo.Context) error {
		var input struct {
			ID   int64  `json:"id"`
			Name string `json:"name"`
		}
		if err := echo.FormFieldBinder(c).Int64("id", &input.ID).String("name", &input.Name).BindError(); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		user := &model.User{}
		err := user.Update(model.User{
			ID:   input.ID,
			Name: input.Name,
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, "ok")
	})

	e.Start("127.0.0.1:1323")
}
