package models

type Shop struct {
	Id      	int64   `orm:"auto"`						// 商店ID
	TotalNum 	int64  	`orm:"default(100)"`				// 总数
	Name    	string 	`orm:"size(100)"`					// 商店名
	Gender  	string 	`orm:"size(10)"`					// 性别
	Desc    	string 	`orm:"size(32)"`					// 商店简介
	Address 	string 	`orm:"size(64)"`					// 商店地址
	Uuid    	string 	`orm:"size(128)"`					// uuid
	Img     	string 	`orm:"size(30)"`					// 商店图片

	PayAccount  string  `orm:"size(20)"`					// 支付账号
	PayImg		string	`orm:"size(20)"`					// 支付二维码
	UserName	string										// 用户名
	Contact 	string										// 内容

}
