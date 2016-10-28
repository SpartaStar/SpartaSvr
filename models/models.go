// models
package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

type test struct{
	Id int
	Name string
}

var Ormer1 orm.Ormer

func init(){
	orm.RegisterModel(new(test))
}

func Dbopr(){	//test orm
	te := new(test)
	te.Id = 1001
	te.Name = "hello"

	fmt.Println(Ormer1.Insert(te))
	te.Name = "world"
	fmt.Println(Ormer1.Update(te))

	te1 := test{Id: 1001}
	
	err := Ormer1.Read(&te1)
	if err == orm.ErrNoRows {
    		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
    		fmt.Println("找不到主键")
	} else {
    		fmt.Println(te1.Id, te1.Name)
	}
}
