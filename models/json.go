// json
package models


type t_gzh struct{
	Wechat_id int
	Wechat_name string
	Author_img_url string
}

type Sst struct{		//subscription list
	Sst_list [] t_gzh
}
