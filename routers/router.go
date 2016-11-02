package routers

import (
	"github.com/astaxie/beego"
	"spartasvr/controllers"
)

func init() {
	beego.Router("/test", &controllers.MainController{})
	beego.Router("/sst_list.json", &controllers.SstListCtl{})
	beego.Router("/article_list.json", &controllers.ArticleListCtl{})
	//beego.Router("/all_article_list.json", &controllers.ArticleListCtl{})
}
