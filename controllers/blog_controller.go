package controllers

import (
	"YunWeiBlog/models"
	"strconv"
)

type BlogController struct {
	BaseController
}

func (c *BlogController) Get() {
	id := c.Ctx.Input.Param(":id")
	id1, err2 := strconv.Atoi(id)
	if err2 != nil {
		c.Redirect("/", 302)
	} else {
		blogInfo, err := models.GetBlogInfo(id1)
		if err == nil {
			c.Data["Data"] = blogInfo
			models.ReadAdd(id1)
			c.TplName = "blog.html"
		} else {
			c.Redirect("/", 302)
		}
	}
}
