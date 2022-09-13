package controllers

import beego "github.com/beego/beego/v2/server/web"

type EditController struct {
	beego.Controller
}

func (c *EditController) Get() {
	c.TplName = "blog_edit.html"
}
