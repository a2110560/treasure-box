package api

import (
	DataBase "Swagger/Db"
	"github.com/labstack/echo/v4"
	"net/http"
)

// @Summary 會員
//
//
type user struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

// AddData
// @Summary 取得資料
// @Tags User
// @Accept       json
// @Produce      json
// @Param        username   body   string  true  "使用者資料"
// @Success      200  {object}  map[string]interface{}  "成功回傳"
// @Failure      400  {string}  string "格式不對"
// @Failure      404  {string}  string "找不到"
// @Failure      500  {string}  string "資料庫錯誤"
// @version 1.0
// @Router /user [post]
func AddData(c echo.Context) error {
	var input user
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	//if err := echo.QueryParamsBinder(c).
	//	String("username", &input.Username).
	//	BindError(); err != nil {
	//	return echo.NewHTTPError(http.StatusInternalServerError, err)
	//}
	if result := DataBase.GetDB().Debug().Create(&input); result.Error != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, result.Error)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "add",
		"data":   input,
	})
}
func UpdateData(c echo.Context) error {
	var input user
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	//if err := echo.QueryParamsBinder(c).
	//	Int("id", &input.Id).
	//	String("username", &input.Username).
	//	BindError(); err != nil {
	//	return err
	//}
	param := c.Param("id")
	if result := DataBase.GetDB().Debug().
		Model(&input).
		Find(&input, param).
		Where("id=?", param).
		Updates(&input); result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, result.Error)
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": "update",
		"data":   input,
	})
}
func DeleteData(c echo.Context) error {
	var input user
	//if err := echo.QueryParamsBinder(c).
	//	Int("id", &input.Id).
	//	BindError(); err != nil {
	//	return err
	//}
	param := c.Param("id")
	if result := DataBase.GetDB().Debug().Model(&input).Find(&input,
		param).Delete(&input, param); result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, result.Error)
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": "del ok",
	})
}

// SearchData
// @Summary      Show an account
// @Description  get string by ID
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]
func SearchData(c echo.Context) error {
	var input user
	//if err := echo.QueryParamsBinder(c).
	//	Int("id", &input.Id).
	//	BindError(); err != nil {
	//	return err
	//}
	param := c.Param("id")
	if result := DataBase.GetDB().Debug().Model(&input).Find(&input,
		param); result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, result.Error)
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": "get",
		"data":   input,
	})
}
func ListData(c echo.Context) error {
	var input user
	if result := DataBase.GetDB().Debug().Find(&input); result.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, result.Error)
	}
	return c.JSON(http.StatusOK, echo.Map{
		"status": "get",
		"data":   input,
	})
}
func (user) TableName() string {
	return "swagger_test"
}
