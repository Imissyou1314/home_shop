package service

import "home_shop/models"

type goodsService struct {
}

func getTabelName() string{
	return "t_good"
}

func GetById(goodId int64) (*models.Goods, error) {
	good := &models.Goods{}
	good.Id = goodId
	err := o.Read(&good)
	return good, err
}

func GetAll() ([]*models.Goods, error) {
	var goodList []*models.Goods
	_, err := o.QueryTable(getTabelName()).Filter("is_show", true).All(&goodList)
	return goodList, err
}

