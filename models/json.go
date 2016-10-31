// json
package models


type SstInfo struct{
	Wechat_id int
	Wechat_name string
	Author_img_url string
}

type Sst struct{		//subscription list
	Sst_list [] SstInfo
}
