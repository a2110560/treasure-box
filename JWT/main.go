package main

import (
	"context"
	"github.com/go-redis/redis/v9"
	"jwt_mysql/JWT"
	"jwt_mysql/Redis"
	"jwt_mysql/Service"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var redisDb *redis.Client
var ctx context.Context

func main() {
	//初始化
	e := echo.New()
	ctx = context.Background()
	//初始化redis
	redisDb = Redis.InitClient()
	{
		e.POST("/register", Register)
		e.POST("/login", Login)
	}

	{
		r := e.Group("/restricted")
		r.Use(middleware.JWTWithConfig(middleware.JWTConfig{
			Claims:     &JWT.Claims{},
			SigningKey: []byte("secret"),
		}))
		r.GET("/user", getInfo)
	}
	e.Logger.Fatal(e.Start("localhost:1323"))
}

func Login(c echo.Context) error {
	var err error
	var token string
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	//格式不對
	if err = echo.QueryParamsBinder(c).
		String("username", &input.Username).
		String("password", &input.Password).
		BindError(); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Status": "Invalid format...",
		})
	}
	//一般要token會先與資料庫比對身分，沒有身分就無法下一步
	user, err := Service.FindUserByUsername(input.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Status": "User Not Found...",
		})
	}
	//對密碼
	if user.Password != input.Password {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Status": "Wrong password...",
		})
	}
	redisToken, err := redisDb.Get(ctx, input.Username).Result()
	if err == redis.Nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Status": "Key is not exist...",
		})
	} else {
		if len(redisToken) > 0 {
			//redis裡面有token(還沒過期)，並且使用者的名字也對
			return c.JSON(http.StatusOK, echo.Map{
				"response_code":    200,
				"response_message": "Success",
				"token":            redisToken,
			})
		} else {
			//redis裡面沒有token(過期)，需refresh token
			if token, err = JWT.GenerateToken(input.Username); err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"Status": err,
				})
			}
			//加入redis快取 (與jwt_token過期時間一致)
			if err = redisDb.Set(ctx, input.Username, token, 10*time.Minute).Err(); err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"Status": err,
				})
			}
			return c.JSON(http.StatusOK, echo.Map{
				"response_code":    200,
				"response_message": "Success",
				"token":            token,
			})
		}
	}
}
func Register(c echo.Context) error {
	var st Service.UserCredentials
	var err error
	var token string
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err = echo.QueryParamsBinder(c).
		String("username", &input.Username).
		String("password", &input.Password).
		BindError(); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Status": "Invalid format...",
		})
	}
	user, _ := Service.FindUserByUsername(input.Username)
	if st == user {
		//代表資料庫沒資料，可以+
		if user, err = Service.InsertUser(input.Username, input.Password); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Status": err,
			})
		}
	} else {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Status": "Duplicate Users....",
		})
	}
	//產生token
	if token, err = JWT.GenerateToken(input.Username); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Status": err,
		})
	}
	//加入redis快取 (與jwt_token過期時間一致)
	if err = redisDb.Set(ctx, input.Username, token, 10*time.Minute).Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Status": err,
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"response_code":    200,
		"response_message": "Success",
		"token":            token,
	})
}

func getInfo(c echo.Context) error {
	var err error
	var input struct {
		Username string `json:"username"`
	}
	//格式不對
	if err = echo.QueryParamsBinder(c).
		String("username", &input.Username).
		BindError(); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Status": "Invalid format...",
		})
	}
	user, err := Service.FindUserByUsername(input.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"Status": "User Not Found...",
		})
	}
	return c.JSON(http.StatusOK, echo.Map{
		"Data": user,
	})
}
