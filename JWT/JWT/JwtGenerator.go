package JWT

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func GenerateToken(name string) (string, error) {
	var err error
	//設立claims
	claims := &Claims{
		name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
		},
	}
	//創一個用claims生成的token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//把token轉成string回傳
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return t, nil
}
