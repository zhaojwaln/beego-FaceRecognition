package service

import (
	"beego-FaceRecognition/src/model"
	"strings"

	"github.com/astaxie/beego"
)

type roleService struct{}

func (this *roleService) LoadMenu() []model.RoleTree {

	var roles []model.RoleTree
	selectSql := "SELECT t.id, pid, name, roleurl , ismenu, des from t_role t where t.id != 0 and t.ismenu = 0"
	if _, err := o.Raw(selectSql).QueryRows(&roles); err != nil {
		beego.Error("查询权限树的role列表异常，error message：", err.Error())
		return roles
	}

	pidMap := make(map[int64]bool, 10)
	for _, role := range roles {
		pidMap[role.Pid] = true
	}

	for i, role := range roles {
		//展开所有父节点
		if pidMap[role.Id] {
			roles[i].Open = true
			continue
		}
		if !strings.EqualFold(role.Roleurl, "") {
			click := "click: addTab('" + roles[i].Name + "','" + roles[i].Roleurl + "')"
			roles[i].Click = click
		}
	}

	return roles
}