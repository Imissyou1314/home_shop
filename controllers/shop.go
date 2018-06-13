package controllers

import (
	"github.com/astaxie/beego"
	"home_shop/models"
)

type ShopController struct {
	beego.Controller
}

// @Title GetAllShop
// @Description get all Shops
// @Success 200 {object} models.Shop
// @router /Shop [get]
func (s *ShopController) GetShopList() {
	shops := models.GetAll()
	s.Data["json"] = shops
	s.ServeJSON()
}