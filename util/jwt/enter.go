package jwt

import (
	"blogx/global"
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 自定义的三个字段
type MyClaim struct {
	UserID   uint   `json:"userID"`
	Username string `json:"username"`
	Role     int8   `json:"role"`
}

// 传入jwt的字段
type Claim struct {
	MyClaim
	jwt.StandardClaims
}

func GetToken(myClaim MyClaim) (string, error) {
	claim :=	Claim{
		MyClaim: myClaim,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Duration(global.Conf.Jwt.Expire) * time.Second).Unix(), // 过期时间
			Issuer:    global.Conf.Jwt.Issuer,                                                     // 签发人
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(global.Conf.Jwt.Secret))
}


func ParseToken(token string) (*Claim, error) {
	if token == "" {
		return nil, errors.New("请登录")
	}
	_token, err := jwt.ParseWithClaims(token, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Conf.Jwt.Secret), nil
	})

	if err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			return nil, errors.New("token过期")
		}
		if strings.Contains(err.Error(), "signature is invalid") {
			return nil, errors.New("token无效")
		}
		if strings.Contains(err.Error(), "token contains an invalid") {
			return nil, errors.New("token非法")
		}
		return nil, err
	}
	if claims, ok := _token.Claims.(*Claim); ok && _token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

func ParseTokenByGin(c *gin.Context) (*Claim, error) {
	token := c.GetHeader("token")
	if token == "" {
		token = c.Query("token")
	}

	return ParseToken(token)
}
