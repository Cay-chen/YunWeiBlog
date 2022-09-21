package controllers

import (
	"YunWeiBlog/models/dao"
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"math"
)

type MyController struct {
	BaseController
}

func (c *MyController) Get() {
	wd := c.GetString("wd")
	xs, err := c.GetInt("xs")
	logs.Notice(xs)
	if err != nil {
		c.Data["Xs"] = 9
	} else {
		c.Data["Xs"] = xs
	}
	c.Data["Wd"] = wd

	c.TplName = "layout.html"
}
func (c *MyController) Post() {
	limit := 8
	var blogInfo []*dao.BlogInfo
	var blogList dao.BlogList
	wd := c.GetString("wd")
	page, _ := c.GetInt("page")
	xs, _ := c.GetInt("xs")
	start := (page - 1) * limit
	cond := orm.NewCondition()
	var cond1 *orm.Condition
	switch xs {
	case 0:
		cond1 = cond.And("BlogVisibleType", 0).And("BlogCreateUser", c.User.UserName).And("BlogState", 0).And("BlogTitle__icontains", wd).Or("BlogBrief__icontains", wd)
		break
	case 1:
		cond1 = cond.And("BlogVisibleType", 1).And("BlogCreateUser", c.User.UserName).And("BlogState", 0).And("BlogTitle__icontains", wd).Or("BlogBrief__icontains", wd)
		break
	case 9:
		cond1 = cond.And("BlogCreateUser", c.User.UserName).And("BlogState", 0).And("BlogTitle__icontains", wd).Or("BlogBrief__icontains", wd)
		break
	}
	o := orm.NewOrm()
	qs := o.QueryTable("blogInfo")
	qs.SetCond(cond1).Limit(limit, start).All(&blogInfo, "BlogId", "BlogTitle", "BlogBrief", "BlogCreateTime", "BlogCreateUser", "BlogImgUrl", "BlogClassifyType", "BlogReadCount")
	count, _ := qs.SetCond(cond1).Count()
	pages := int(math.Ceil(float64(count) / float64(limit)))
	blogList.Pages = pages
	blogList.Data = blogInfo
	backRes, _ := json.Marshal(blogList)
	c.Ctx.WriteString(string(backRes))
}
