package routers

import (
	"sso/api"
)

func (router RouterGroup) UserRouter() {
	app := api.ApiGroupApp.UserApi
	router.POST("login", app.UserLoginView) // 用户登录
}