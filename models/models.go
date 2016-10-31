// models
package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var Ormer1 orm.Ormer

func init() {
	orm.RegisterModel(new(test))
}

func GetSstList() Sst {
	//操作数据库获取公众号信息
	//TODO
	Ormer1 = orm.NewOrm()
	var sst Sst
	
	var sstinfos [] SstInfo
	Ormer1.QueryTable("t_gzh").Limit(20).All(&sstinfos, "wechat_id", "wechat_name", "author_img_url")
	sst.Sst_list = append(sst.Sst_list, sstinfos...)
	return sst //
}

/****************** test *****************************/
type test struct {
	Id   int
	Name string
}

func Dbopr() { //test orm
	Ormer1 = orm.NewOrm()

//	te := new(test)
//	te.Id = 1001
//	te.Name = "hello"

//	fmt.Println(Ormer1.Insert(te))
//	te.Name = "world"
//	fmt.Println(Ormer1.Update(te))

//	te1 := test{Id: 1001}

//	err := Ormer1.Read(&te1)
//	if err == orm.ErrNoRows {
//		fmt.Println("查询不到")
//	} else if err == orm.ErrMissPK {
//		fmt.Println("找不到主键")
//	} else {
//		fmt.Println(te1.Id, te1.Name)
//	}

	var tes []test
	Ormer1.QueryTable("test").Limit(5).All(&tes,"name", "id")
	fmt.Println(tes)
}
/****************** test end *****************************/
