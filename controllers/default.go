package controllers

import (
	"github.com/astaxie/beego"
	//"spartasvr/models"
)

type MainController struct {
	beego.Controller
}

type SstlistCtl struct{
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

	//models.Dbopr()	//test orm
}

func (SstlistCtl) Get(){
	//可以考虑利用url后面的参数
	//TODO
}
