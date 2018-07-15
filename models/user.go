package models

type User struct {
	Id     		int64  `orm:"auto"`
	Name   		string `orm:"size(100)"`		// 用户名
	Gender 		string `orm:"size(10)"`
	Phone  		string `orm:"size(32)"`			// 用户手机
	Email  		string `orm:"size(32)"`			// 用户邮箱
	Address 	string `orm:"size(64)"`			// 用户地址
	Uuid 		string `orm:"size(128)"`		// uuid
	Password 	string `orm:"size(128)"`		// 用户密码
	Salt 		string `orm:"size(6)"`			// 加密盐值
}