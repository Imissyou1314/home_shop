package service

import (
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"os"
	"home_shop/models"
)

var (
	o				orm.Ormer
	tablePrefix		string
	UserService		*userService
	GoodsService	*goodsService
	ShopService 	*shopService
	OrderService	*orderService

)

func init() {
	dbHost := beego.AppConfig.String("mysqlurls")
	dbPort := beego.AppConfig.String("mysqlport")
	dbUser := beego.AppConfig.String("mysqluser")
	dbPassword := beego.AppConfig.String("mysqlpass")
	dbName := beego.AppConfig.String("mysqldb")
	tablePrefix = beego.AppConfig.String("table.prefix")

	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&loc=Asia/Shanghai"

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dsn)

	orm.RegisterModelWithPrefix(tablePrefix,
		new (models.User),
		new(models.Goods),
		new(models.Order),
		new(models.Shop),
	)

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}

	o = orm.NewOrm()
	orm.RunCommand()

	// 创建代码
	os.Mkdir("logs", 0755)

	initService()

}

func initService() {
	UserService = &userService{}
	GoodsService = &goodsService{}
	OrderService = &orderService{}
	ShopService = &shopService{}
}
