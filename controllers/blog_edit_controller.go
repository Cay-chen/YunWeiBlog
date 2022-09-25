package controllers

import (
	"YunWeiBlog/models"
	"strconv"
	"strings"
)

type EditController struct {
	BaseController
}

func (c *EditController) Get() {
	if c.IsLogin {
		blogId, err9 := c.GetInt("BlogId")
		if err9 != nil {
			c.Data["IsXiu"] = true
			c.Data["Title"] = "文章编辑"
			c.Data["ButtonSubmit"] = "发布博客"
		} else {
			c.Data["Title"] = "修改文章"
			c.Data["IsXiu"] = false
			c.Data["ButtonSubmit"] = "修改博客"
		}
		classifyBox := ""
		classify := models.GetClassifyList()
		blogInfo, err := models.GetBlogInfo(blogId)
		if err == nil {
			c.Data["IsEdit"] = true
			switch blogInfo.BlogVisibleType {
			case 0:
				c.Data["AllKj"] = "checked"
				break
			case 1:
				c.Data["JwKj"] = "checked"
				break
			case 2:
				c.Data["DlKj"] = "checked"
				break
			}
			for i := 0; i < len(classify); i++ {
				if strings.Contains(blogInfo.BlogClassifyType, strconv.Itoa(classify[i].ClassifyId)+"|") {
					classifyBox = classifyBox + "<input type=\"checkbox\" name=\"flzl" + strconv.Itoa(classify[i].ClassifyId) + "\" lay-skin=\"primary\" value=\"" + strconv.Itoa(classify[i].ClassifyId) + "\" title=\"" + classify[i].ClassifyName + "\" checked>\n                    "
				} else {
					classifyBox = classifyBox + "<input type=\"checkbox\" name=\"flzl" + strconv.Itoa(classify[i].ClassifyId) + "\" lay-skin=\"primary\" value=\"" + strconv.Itoa(classify[i].ClassifyId) + "\" title=\"" + classify[i].ClassifyName + "\">\n                    "
				}
			}
			c.Data["Data"] = blogInfo
		} else {
			for i := 0; i < len(classify); i++ {
				classifyBox = classifyBox + "<input type=\"checkbox\" name=\"flzl" + strconv.Itoa(classify[i].ClassifyId) + "\" lay-skin=\"primary\" value=\"" + strconv.Itoa(classify[i].ClassifyId) + "\" title=\"" + classify[i].ClassifyName + "\">\n                    "
			}
			c.Data["IsEdit"] = false
			c.Data["AllKj"] = "checked"
		}
		c.Data["ClaaifyBox"] = classifyBox
		c.TplName = "blog_edit.html"
	} else {
		c.Redirect("/", 302)
	}
}
func (c *EditController) Post() {
	blogId, _ := c.GetInt("blog_id")
	data := c.Ctx.Input.RequestBody
	resResult, err := models.InsertAndEditBlogInfo(blogId, c.User.UserId, data)
	if err != nil {
		resResult = models.GetNilBlogInfoList()
	}
	c.Ctx.WriteString(resResult)
}
