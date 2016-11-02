// models
package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//"time"
)

var Ormer1 orm.Ormer

func init() {
	orm.RegisterModel(new(test))
}

func GetSstList() (sst SstList){
	//操作数据库获取公众号信息
	Ormer1 = orm.NewOrm()

	var ssts [] SstInfo
	//Ormer1.QueryTable("t_gzh").Limit(1).All(&sstinfos, "wechat_id", "wechat_name", "author_img_url")
	Ormer1.Raw("select wechat_id, wechat_name, author_img_url from t_gzh limit 10").QueryRows(&ssts)
	sst.Sst_list = append(sst.Sst_list, ssts...)

	return sst
}

func GetArticleList() (atl ArticleList){
	Ormer1 = orm.NewOrm()
	//TODO : parse the upload json
	
	var atsBuf [] ArticleInfo
	Ormer1.Raw("select  from t_article").QueryRows(&atsBuf)	//TODO : sql

	//TODO :atsBuf is a buffer,根据客户端上传的文章数量和文章索引来获取buffer中相应的数据
	var ats [] ArticleInfo
	atl.Article_list = append(atl.Article_list, ats...)

	return atl
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
