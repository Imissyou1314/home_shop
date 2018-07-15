package service

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego"
	"os"
	"home_shop/models"
	"github.com/astaxie/beego/logs"
	"fmt"
)

var (
	o            orm.Ormer
	tablePrefix  string
	UserService  *userService
	GoodsService *goodsService
	ShopService  *shopService
	OrderService *orderService
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
		new(models.User),
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

func CheckErr(err error) bool{
	if nil != err {
		logs.Error(err)
		return false
	}
	return true
}

// New Error 生成异常
func NewErr(errMsg string) error {
	return fmt.Errorf("Error Msg: %s", errMsg)
}

