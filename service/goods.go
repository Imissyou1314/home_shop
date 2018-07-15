package service

import "home_shop/models"

type goodsService struct {
}

func (g *goodsService) GetTabelName() string{
	return "t_good"
}

func (g *goodsService) GetById(goodId int64) (*models.Goods, error) {
	good := &models.Goods{}
	good.Id = goodId
	err := o.Read(&good)
	return good, err
}

func (g *goodsService) GetAll() ([]*models.Goods, error) {
	var goodList []*models.Goods
	_, err := o.QueryTable(g.GetTabelName()).Filter("is_show", true).All(&goodList)
	return goodList, err
}

func (g *goodsService) CreateGoods(good *models.Goods) (*models.Goods, error) {
	if created, id, err := o.ReadOrCreate(&good, "Id"); err == nil {
		if created {
			return g.GetById(id)
		} else {
			return nil, NewErr("该订单已经存在")
		}
	} else {
		return nil, err
	}
}

func (g *goodsService) UpdateGoods(good *models.Goods) (*models.Goods, error) {
	if id, err := o.Update(&good); err != nil {
		return nil, err
	} else {
		return g.GetById(id)
	}
}

