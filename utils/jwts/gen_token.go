package jwts

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"sso/global"
	"time"
)

func GenToken(user JwtPayLoad) (string, error)  {
	//global.CONFIG = core.ConfInit()
	MySecret = []byte(global.CONFIG.Jwt.Secret)
	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.CONFIG.Jwt.Expires))), // 默认2小时过期
			Issuer:    global.CONFIG.Jwt.Issuer,                                                     // 签发人
		},
	}
	fmt.Println(global.CONFIG.Jwt.Expires)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}
