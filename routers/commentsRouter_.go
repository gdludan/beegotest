package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["beegotest/controllers:MainController"] = append(beego.GlobalControllerRouter["beegotest/controllers:MainController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           "/",
			AllowHTTPMethods: []string{"Get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
