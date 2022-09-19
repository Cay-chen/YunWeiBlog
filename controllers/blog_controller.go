package controllers

import (
	"YunWeiBlog/models/dao"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
)

type BlogController struct {
	BaseController
}

func (c *BlogController) Get() {
	id := c.Ctx.Input.Param(":id")
	o := orm.NewOrm()
	id1, _ := strconv.Atoi(id)
	u := dao.BlogInfo{BlogId: id1}
	err := o.Read(&u)
	if err == nil {
		c.Data["Data"] = u
	} else {
		c.Data["Html"] = "<p>没内容！</p>"
	}
	o.QueryTable("BlogInfo").Filter("BlogId", u.BlogId).Update(orm.Params{
		"BlogReadCount": orm.ColValue(orm.ColAdd, 1),
	})
	c.TplName = "blog.html"
}
