// json
package models

/*************************** Downstream Data ********************************/
type SstInfo struct{
	Wechat_id string
	Wechat_name string
	Author_img_url string
}

type SstList struct{		//subscription list
	Sst_list [] SstInfo
}

type ArticleInfo struct{
	Title string
	//picture url???
}

type ArticleList struct{
	Article_list [] ArticleInfo
}

/*************************** Upstream Data ********************************/