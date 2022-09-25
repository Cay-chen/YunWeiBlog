package controllers

import "YunWeiBlog/models"

type PersonController struct {
	BaseController
}

func (c *PersonController) Get() {

	c.TplName = "iframe_person_manage.html"

}
func (c *PersonController) Post() {
	limit, _ := c.GetInt("limit")
	page, _ := c.GetInt("page")
	backRes, err := models.GetUsersList(limit, page)
	if err != nil {
		backRes = models.GetNilBlogInfoList()
	}
	c.Ctx.WriteString(backRes)
}
