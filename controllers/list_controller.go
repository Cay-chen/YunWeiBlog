package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type ListController struct {
	beego.Controller
}

func (c *ListController) Get() {
	c.TplName = "blog_list.html"
}
