package service

import (
	"win-xp/common"
	"win-xp/models"
	"fmt"
)

type SysMenuService struct {
}

func (this *SysMenuService) ListMenu() []models.SysMenuTree {
	treeList := make([]models.SysMenuTree, 0)
	menu := make([]models.SysMenu, 0)
	err := common.Engine.Find(&menu)
	fmt.Println(err)
	menuMap := make(map[int]models.SysMenuTree)
	//将父菜单添加进去
	for i := 0; i < len(menu); i++ {
		if menu[i].ParentId == 0 {
			menuTree := models.SysMenuTree{}
			menuTree.Menu = menu[i]
			menuTree.SubMenu = make([]models.SysMenu, 0)
			menuMap[menu[i].Id] = menuTree
		}
	}

	//遍历查找对应子菜单
	var parentId int
	for i := 0; i < len(menu); i++ {
		parentId = menu[i].ParentId
		if parentId != 0 {
			menuItem := menuMap[parentId]
			if &menuItem != nil {
				menuItem.SubMenu = append(menuItem.SubMenu, menu[i])
				menuMap[parentId] = menuItem
			}
		}
	}
	for i := 0; i < len(menu); i++ {
		if menu[i].ParentId == 0 {
			treeList = append(treeList, menuMap[menu[i].Id])
		}
	}
	return treeList
}
