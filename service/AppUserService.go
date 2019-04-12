package service

import (
	"win-xp/common"
	"win-xp/models"
)

type AppUserService struct {
}

func (this *AppUserService) SelectById(id int) *models.AppUser {
	user := &models.AppUser{}
	user.Id = id
	is, _ := common.Engine.Get(user)
	if !is {
		return nil
	}
	return user
}

func (this *AppUserService) UserLogin(userName string, password string) *models.AppUser {
	user := &models.AppUser{}
	is, _ := common.Engine.Where("nick_name=?", userName).Where("password=?", password).Get(user)
	if is {
		return user
	}
	return nil
}
