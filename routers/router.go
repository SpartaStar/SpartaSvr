package routers

import (
	"spartasvr/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/test", &controllers.MainController{})
	beego.Router("/sstlist", &controllers.SstlistCtl{})
}
