package controllers

import (
	"beego-FaceRecognition/src/common"
	"fmt"
	"strings"

	"github.com/astaxie/beego"
)

const (
	SUCCESS = "success"
)

type BaseController struct {
	beego.Controller
	controllerName string          // 控制器名
	actionName     string          // 动作名
	openPerm       map[string]bool // 公开的权限
}

/**
初始化开放权限(不需要权限校验的操作,后续如果有不需要权限校验的操作都可以写在这里)
*/
func (this *BaseController) initOpenPerm() {
	this.openPerm = map[string]bool{
		"MainController.LeftMenu": true,
		"MainController.Norole":   true,
	}
}

/**
判断是否是不需要鉴权的公共操作
*/
func (this *BaseController) isOpenPerm() bool {
	//如果是登陆相关操作则不进行登陆鉴权和权限鉴权等操作
	if strings.EqualFold(this.controllerName, "logincontroller") {
		return true
	}
	this.initOpenPerm()
	key := this.controllerName + "." + this.actionName
	if this.openPerm[key] {
		return true
	}
	return false
}

/*
指定页面，并且返回公共参数
*/
func (this *BaseController) show(url string) {
	this.Data["staticUrl"] = beego.AppConfig.String("staticUrl")
	this.TplName = url
}

/**
把需要返回的结构序列化成json 输出
*/
func (this *BaseController) jsonResult(result interface{}) {
	this.Data["json"] = result
	this.ServeJSON()
	this.StopRun()
}

/**
获取IP
*/
func (this *BaseController) getClientIp() string {
	ip := this.Ctx.Request.Header.Get("Remote_addr")
	if ip == "" {
		ip = this.Ctx.Request.RemoteAddr
	}
	fmt.Println(ip)
	if strings.Contains(ip, ":") {
		ip = common.Substr(ip, 0, strings.Index(ip, ":"))
	}
	fmt.Println(ip)
	return ip
}

type Empty struct {
}

/*
 用于分页展示列表的时候的 输出json
*/
func (this *BaseController) jsonResultPager(count int, roles interface{}) {
	beego.Debug("分页数据：", count, roles)
	resultMap := make(map[string]interface{}, 1)
	if count == 0 || roles == nil {
		beego.Debug("查询分页数据为空，返回默认json")
		//这里默认totle设置为1是因为easyui分页控件如果totle 为0会出现错乱
		resultMap["total"] = 1
		resultMap["rows"] = make([]Empty, 0)
	} else {
		resultMap["total"] = count
		resultMap["rows"] = roles
	}
	this.Data["json"] = resultMap
	this.ServeJSON()
	this.StopRun()
}
