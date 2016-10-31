package routers

import (
	"github.com/astaxie/beego"
	"spartasvr/controllers"
)

func init() {
	beego.Router("/test", &controllers.MainController{})
	beego.Router("/sst_list", &controllers.SstlistCtl{})
}
