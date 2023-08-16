package flags

import (

	"fmt"
	"sso/global"
)

func Version() {
	fmt.Printf("项目版本：%s", global.CONFIG.System.Version)
}
