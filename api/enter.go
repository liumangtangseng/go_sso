package api

import "sso/api/user_api"

type ApiGroup struct {
	UserApi user_api.UserApi
}

var ApiGroupApp = new(ApiGroup)