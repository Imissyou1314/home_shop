package models

type Shop struct {
	Id     int    `orm:"auto"`
	Name   string `orm:"size(100)"`
	Gender string `orm:"size(10)"`
	desc  string `orm:"size(32)"`
	price  int64 `orm:default(0)`
	Address string `orm:"size(64)"`
	Uuid string `orm:"size(128)"`
	Password string `orm:"size(128)"`
	Salt string `orm:"size(6)"`
}
