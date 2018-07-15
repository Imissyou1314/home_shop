package service

import "home_shop/models"

type userService struct {

}

func (u *userService) GetTableName() string {
	return "t_user"
}

func (u *userService) GetById(userId int64) (*models.User, error) {
	user := &models.User{}
	user.Id = userId
	err := o.Read(&user)
	return user, err
}

func (u *userService) GetByUserName(userName string) (*models.User, error)  {
	user := &models.User{}
	user.Name = userName
	o.Using(u.GetTableName())
	err := o.Read(&user)
	return user, err
}