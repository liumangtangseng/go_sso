package jwts

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/sirupsen/logrus"
	"sso/global"
)

// ParseToken 解析 token
func ParseToken(tokenStr string) (*CustomClaims, error) {
	//global.CONFIG = core.ConfInit()
	MySecret = []byte(global.CONFIG.Jwt.Secret)
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})
	if err != nil {
		logrus.Error(fmt.Sprintf("token parse err: %s", err.Error()))
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}