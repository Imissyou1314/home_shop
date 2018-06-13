package models

import "time"

type OrderState int

const (
	_ OrderState = iota
	Loacked				// 已锁定
	Payment				// 待付款
	Paid				// 已付款
	StaySend			// 待发货
	Sending				// 已发货
	Successed			// 交易成功
	Colsed				// 关闭交易
	Failured			// 交易失败
)

type Order struct {
	Id 				int
	ShopId			int
	UserId          int
	ShopNumber		int64			`orm:"default(0)"`
	Price			int64			`orm:"default(0)"`
	Desc 			string			`orm:"size(60);null"`
	CreateTime 		time.Time		`orm:"auto_now_add;type(datetime)"`
	EndTime			time.Time       `orm:"type(datetime);null"`
	State			int				`orm:"default(1)"`

	MailCode		string			`orm:"size(15);null"`					// 邮件码
	MailType		string 			`orm:"size(19);null"`					// 邮件类型
}
