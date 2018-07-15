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
	Id 				int64
	ShopId			int														// 商店关联ID
	UserId          int														// 用户关联iD
	ShopNumber		int64			`orm:"default(0)"`						// 商品数量
	Price			int64			`orm:"default(0)"`						// 商品价格
	Desc 			string			`orm:"size(60);null"`					// 购买说明
	CreateTime 		time.Time		`orm:"auto_now_add;type(datetime)"`		// 购买时间
	EndTime			time.Time       `orm:"type(datetime);null"`				// 订单结束时间
	State			int				`orm:"default(1)"`						// 订单状态

	MailCode		string			`orm:"size(15);null"`					// 邮件码
	MailType		string 			`orm:"size(19);null"`					// 邮件类型
}
