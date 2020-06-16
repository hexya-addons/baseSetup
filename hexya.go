package baseSetup

import (
	_ "github.com/hexya-addons/base"
	_ "github.com/hexya-addons/web"
	"github.com/hexya-erp/hexya/src/server"
)

const MODULE_NAME string = "baseSetup"

func init() {
	server.RegisterModule(&server.Module{
		Name:     MODULE_NAME,
		PreInit:  func() {},
		PostInit: func() {},
	})

}
