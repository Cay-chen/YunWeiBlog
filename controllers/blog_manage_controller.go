package controllers

import (
	"YunWeiBlog/models"
	"YunWeiBlog/models/dao"
	"encoding/json"
)

type BlogManageController struct {
	BaseController
}

func (c *BlogManageController) Get() {
	if c.IsLogin {
		wd := c.GetString("wd")
		//beginDate := c.GetString("beginDate")
		//endDate := c.GetString("endDate")
		xs, err := c.GetInt("xs")
		if err != nil {
			c.Data["Xs"] = "-1"
		} else {
			c.Data["Xs"] = xs
		}
		switch xs {
		case 0:
			c.Data["Kj0"] = "selected"
			break
		case 1:
			c.Data["Kj1"] = "selected"
			break
		case 2:
			c.Data["Kj2"] = "selected"

			break
		case -1:
			c.Data["Kj9"] = "selected"
			break
		}
		c.Data["Wd"] = wd
		c.TplName = "iframe_blog_manage.html"
	} else {
		c.Redirect("/error/600", 302)
	}
}
func (c *BlogManageController) Post() {
	if c.IsLogin {
		wd := c.GetString("wd")
		page, _ := c.GetInt("page")
		xs, _ := c.GetInt("xs")
		backRes, err := models.GetApiBlogInfoList(wd, page, -1, xs, 0, c.User.UserId)
		if err != nil {
			backRes = models.GetNilBlogInfoList()
		}
		c.Ctx.WriteString(backRes)
	} else {
		var jsonResult dao.JsonResult
		jsonResult.Code = -1
		jsonResult.Msg = "没登陆啊！"
		backRes, _ := json.Marshal(jsonResult)
		c.Ctx.WriteString(string(backRes))
	}

}
