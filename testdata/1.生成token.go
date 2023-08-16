package main

import (
	"fmt"
	"sso/utils/jwts"
)

func main()  {
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: "chunfei",
		Role: 2,
		UserID: 1,
	})

	fmt.Println(token,err)
}


