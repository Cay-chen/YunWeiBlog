package controllers

import (
	"YunWeiBlog/models/dao"
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"math"
)

type BlogManageController struct {
	BaseController
}

func (c *BlogManageController) Get() {
	if c.IsLogin {
		wd := c.GetString("wd")
		xs, err := c.GetInt("xs")
		logs.Notice("asdasdasdasdas")
		if err != nil {
			c.Data["Xs"] = "9"
		} else {
			c.Data["Xs"] = xs
		}
		c.Data["Wd"] = wd
		c.TplName = "blog_manage.html"
	} else {
		c.Redirect("/", 302)
	}
}
func (c *BlogManageController) Post() {
	if c.IsLogin {
		limit := 8
		var blogInfo []*dao.BlogInfo
		var blogList dao.BlogList
		wd := c.GetString("wd")
		page, _ := c.GetInt("page")
		xs, _ := c.GetInt("xs")
		start := (page - 1) * limit
		end := page * limit
		cond := orm.NewCondition()
		logs.Notice("aaaaaa" + wd)
		var cond1 *orm.Condition
		switch xs {
		case 0:
			cond1 = cond.And("BlogVisibleType", 0).And("BlogCreateUser", c.User.UserId).And("BlogState", 0).AndCond(cond.And("BlogTitle__icontains", wd).Or("BlogBrief__icontains", wd))
			break
		case 1:
			cond1 = cond.And("BlogVisibleType", 1).And("BlogCreateUser", c.User.UserId).And("BlogState", 0).AndCond(cond.And("BlogTitle__icontains", wd).Or("BlogBrief__icontains", wd))
			break
		case 9:
			cond1 = cond.And("BlogCreateUser", c.User.UserId).And("BlogState", 0).AndCond(cond.And("BlogTitle__icontains", wd).Or("BlogBrief__icontains", wd))
			break
		}
		o := orm.NewOrm()
		qs := o.QueryTable("blogInfo")
		qs.SetCond(cond1).Limit(end, start).All(&blogInfo, "BlogId", "BlogTitle", "BlogBrief", "BlogCreateTime", "BlogCreateUser", "BlogImgUrl", "BlogClassifyType", "BlogReadCount")
		count, _ := qs.SetCond(cond1).Count()
		pages := int(math.Ceil(float64(count) / float64(limit)))
		blogList.Pages = pages
		blogList.Data = blogInfo
		backRes, _ := json.Marshal(blogList)
		c.Ctx.WriteString(string(backRes))
	} else {
		var jsonResult dao.JsonResult
		jsonResult.Code = -1
		jsonResult.Msg = "没登陆啊！"
		backRes, _ := json.Marshal(jsonResult)
		c.Ctx.WriteString(string(backRes))
	}

}
