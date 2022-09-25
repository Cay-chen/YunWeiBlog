package controllers

import (
	"YunWeiBlog/models"
	"fmt"
	"github.com/beego/beego/v2/core/logs"
)

type ListController struct {
	BaseController
}

func (c *ListController) Get() {
	wd := c.GetString("wd")
	c.Data["Wd"] = wd
	c.Data["ClassifyList"] = models.GetClassifyList()
	c.TplName = "blog_list.html"
}
func (c *ListController) Post() {
	wd := c.GetString("wd")
	page, _ := c.GetInt("page")
	//classify, _ := c.GetInt("classify")
	backRes, err := models.GetApiBlogInfoList(wd, page, -1, 0, 0, -1)
	if err != nil {
		logs.Error(err.Error())
		backRes = models.GetNilBlogInfoList()
	}
	fmt.Println(backRes)
	c.Ctx.WriteString(backRes)
}
