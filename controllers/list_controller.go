package controllers

import (
	"YunWeiBlog/models/dao"
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"math"
)

type ListController struct {
	BaseController
}

func (c *ListController) Get() {
	wd := c.GetString("wd")
	c.Data["Wd"] = wd
	c.TplName = "blog_list.html"
}
func (c *ListController) Post() {
	limit := 8
	var blogInfo []*dao.BlogInfo
	var blogList dao.BlogList
	wd := c.GetString("wd")
	page, _ := c.GetInt("page")
	start := (page - 1) * limit
	cond := orm.NewCondition()
	cond1 := cond.And("BlogVisibleType", 0).And("BlogState", 0).And("BlogTitle__icontains", wd).Or("BlogBrief__icontains", wd)
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
