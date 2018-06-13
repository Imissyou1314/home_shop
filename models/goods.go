package models

import "time"

type GoodsType int

const (
	_ GoodsType = iota
	Fruit				// 水果
	Seafood				// 海鲜
	Dry					// 干货
)

type Goods struct {
	Id					int				`orm:"auto"`			// ID
	Name				string			`orm:"size(50)"`		// 商品名
	Desc 				string			`orm:"size(50)"`		// 商品简介
	DescLink 			string 		`orm:"size(30)"`			// 商定链接
	Images   			string		`orm:"size(100)"`			// 图片链接
	Type				int

	Unit    			string									// 单位
	Price				int64
	TotalSales 			int64 		`orm:"size(20)"`			// 销售量

	Postage 			int64									// 邮费
	HasPostage 			bool		`orm:"default(false)"`		// 包不包邮
	LowestPostage		int64		`orm:"default(0)"`			// 最低包邮费用

	HasConditions		bool		`orm:"default(false)"`		// 是否有条件
	LowestConditions	int64									// 最低打折条件
	Discount 			float32		`orm:"default(0.0)"`		// 打折折扣
	Preferential   		int64		`orm:"default(0)"`			// 优惠价格
	ActivityDesc		string 		`orm:"size(40)"`			// 活动说明
	IsActivity			bool		`orm:"default(false)"`		// 是否活动
	StartActivityTime	time.Time	`orm:"type(datetime)"`		// 活动时间
	EndActivityTime		time.Time	`orm:"type(datetime)"`		// 结束活动时间
}
