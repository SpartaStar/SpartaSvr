package main

import (
	_ "spartasvr/routers"
	_ "spartasvr/models"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
)

func init(){
	//orm.RegisterDataBase("sparta", "mysql", "sparta:sparta123@tcp(52.197.76.174:3306)/sparta?charset=utf8")
}

func main() {
	//Ormer1 := orm.NewOrm()
	//Ormer1.Using("sparta")

	beego.SetLogger("file", `{"filename":"test.log"}`)
	//beego.SetLevel(beego.LevelInformational)
	beego.SetLogFuncCall(true)

	beego.Informational("Sparta start......")
	beego.Run()
}

