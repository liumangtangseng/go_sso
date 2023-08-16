package main

import (
	"fmt"
	"sso/utils/jwts"
)

func main()  {
	claims, err := jwts.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuaWNrX25hbWUiOiJjaHVuZmVpIiwicm9sZSI6MiwidXNlcl9pZCI6MSwiYXZhdGFyIjoiIiwiZXhwIjoxNjkyMTYxNjgwLjE4OTcwNSwiaXNzIjoiYWtyb3N0YXIifQ.UnOfz4u_-rGzo2RvvUfwP_2sw5A-5-JYHe-db8EWNU0")
	fmt.Println(claims,err)
}
