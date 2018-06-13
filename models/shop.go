package models

type Shop struct {
	Id      	int    	`orm:"auto"`						// 商店ID
	TotalNum 	int64  	`orm:"default(100)"`				// 总数
	Name    	string 	`orm:"size(100)"`					// 商店名
	Gender  	string 	`orm:"size(10)"`					// 性别
	Desc    	string 	`orm:"size(32)"`
	Address 	string 	`orm:"size(64)"`
	Uuid    	string 	`orm:"size(128)"`
	Img     	string 	`orm:"size(30)"`

	PayAccount  string  `orm:"size(20)"`					// 支付账号
	PayImg		string	`orm:"size(20)"`					// 支付二维码
	UserName	string
	Contact 	string

}

func GetAll() Shop{

}
