package controllers

import (
	"YunWeiBlog/models/dao"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

type ListController struct {
	BaseController
}

func (c *ListController) Get() {
	var blogInfo []*dao.BlogInfo
	o := orm.NewOrm()
	qs := o.QueryTable("blogInfo")
	num, err := qs.Filter("BlogVisibleType", 0).Limit(20).All(&blogInfo)
	if err != nil {
		logs.Error(err.Error())
	}
	logs.Notice(num)
	c.Data["BlogInfoList"] = blogInfo
	c.TplName = "blog_list.html"
}
