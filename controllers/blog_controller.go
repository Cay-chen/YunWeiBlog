package controllers

import (
	"YunWeiBlog/models/dao"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"strconv"
)

type BlogController struct {
	beego.Controller
}

func (c *BlogController) Get() {
	id := c.Ctx.Input.Param(":id")
	o := orm.NewOrm()
	id1, _ := strconv.Atoi(id)
	u := dao.BlogInfo{BlogId: id1}
	err := o.Read(&u)
	if err == nil {
		c.Data["Html"] = u.BlogContent
		c.Data["Title"] = u.BlogTitle
	} else {
		c.Data["Html"] = "<p>没内容！</p>"

	}
	c.TplName = "blog.html"
}
