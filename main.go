package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "spartasvr/models"
	_ "spartasvr/routers"
)

func init() {
	orm.RegisterDataBase("default", "mysql", "sparta:sparta123@tcp(52.197.76.174:3306)/sparta?charset=utf8")
}

func main() {

	beego.SetLogger("file", `{"filename":"test.log"}`)
	//beego.SetLevel(beego.LevelInformational)
	beego.SetLogFuncCall(true)

	beego.Informational("Sparta start......")
	beego.Run()
}
