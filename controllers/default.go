package controllers

import (
	"github.com/astaxie/beego"
	"spartasvr/models"
	"encoding/json"
)

type MainController struct {
	beego.Controller
}

type SstlistCtl struct {
	beego.Controller
}

func (c *MainController) Get() {
	models.Dbopr() //test orm
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

}

func (ctl *SstlistCtl) Get() {
	//可以考虑利用url后面的参数
	//var arg string 
	//c.Ctx.Input.Bind(&arg, "city")	//c为controller，获取url city参数值
	sstlist := models.GetSstList()
	data, _ := json.Marshal(sstlist)
	ctl.Ctx.WriteString(string(data))
}
