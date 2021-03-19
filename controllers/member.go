package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"restful/models"
)

type MemberController struct {
	beego.Controller
}

// @router / [get]
func (c *MemberController) GetAllCNM() {
	data := models.GetAllMember()
	c.Data["json"] = data
	_ = c.ServeJSON()
}

// @router /one [get]
func (c *MemberController) GetOne() {
	data := models.GetDBMember()
	c.Data["json"] = data
	_ = c.ServeJSON()
}
