package controllers

import (
	"github.com/astaxie/beego"
	"spartasvr/models"
	"encoding/json"
)

type SstListCtl struct {
	beego.Controller
}

type ArticleListCtl struct{
	beego.Controller
}

func (ctl *SstListCtl) Get() {
	//可以考虑利用url后面的参数
	//var arg string 
	//c.Ctx.Input.Bind(&arg, "city")	//c为controller，获取url city参数值
	sstlist := models.GetSstList()
	data, _ := json.Marshal(sstlist)
	ctl.Ctx.WriteString(string(data))
}

func (ctl *ArticleListCtl) Get() {
	sstlist := models.GetArticleList()
	data, _ := json.Marshal(sstlist)
	ctl.Ctx.WriteString(string(data))
}

/************************** test *********************/
type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	models.Dbopr() //test orm
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"

}
/************************** test end *********************/
