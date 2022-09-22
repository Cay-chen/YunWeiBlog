package controllers

import (
	"YunWeiBlog/models/dao"
	"github.com/beego/beego/v2/client/orm"
	"strconv"
	"strings"
)

type EditController struct {
	BaseController
}

func (c *EditController) Get() {
	if c.IsLogin {
		var classify []dao.BlogClassify
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
		//isClassifyChecked :=""
		o := orm.NewOrm()
		u := dao.BlogInfo{BlogId: blogId}
		err := o.Read(&u)
		qs := o.QueryTable("BlogClassify")
		qs.Filter("ClassifyState", 0).All(&classify)
		if err == nil {
			c.Data["IsEdit"] = true
			switch u.BlogVisibleType {
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
			//strings.Split(u.BlogClassifyType,"|")
			for i := 0; i < len(classify); i++ {
				if strings.Contains(u.BlogClassifyType, strconv.Itoa(classify[i].ClassifyId)+"|") {
					classifyBox = classifyBox + "<input type=\"checkbox\" name=\"flzl" + strconv.Itoa(classify[i].ClassifyId) + "\" lay-skin=\"primary\" value=\"" + strconv.Itoa(classify[i].ClassifyId) + "\" title=\"" + classify[i].ClassifyName + "\" checked>\n                    "
				} else {
					classifyBox = classifyBox + "<input type=\"checkbox\" name=\"flzl" + strconv.Itoa(classify[i].ClassifyId) + "\" lay-skin=\"primary\" value=\"" + strconv.Itoa(classify[i].ClassifyId) + "\" title=\"" + classify[i].ClassifyName + "\">\n                    "
				}
				//isClassifyChecked= isClassifyChecked +"if (typeof (data.field.flzl"+ strconv.Itoa(classify[i].ClassifyId)  +") != 'undefined') {\n                flzl = flzlFun(flzl, data.field.flzl"+ strconv.Itoa(classify[i].ClassifyId) +")\n            }\n"

			}
			c.Data["Data"] = u
		} else {
			for i := 0; i < len(classify); i++ {
				classifyBox = classifyBox + "<input type=\"checkbox\" name=\"flzl" + strconv.Itoa(classify[i].ClassifyId) + "\" lay-skin=\"primary\" value=\"" + strconv.Itoa(classify[i].ClassifyId) + "\" title=\"" + classify[i].ClassifyName + "\">\n                    "
				//isClassifyChecked= isClassifyChecked +"if (typeof (data.field.flzl"+ strconv.Itoa(classify[i].ClassifyId)  +") != 'undefined') {\n                flzl = flzlFun(flzl, data.field.flzl"+ strconv.Itoa(classify[i].ClassifyId) +")\n            }\n"
			}
			c.Data["IsEdit"] = false
			c.Data["AllKj"] = "checked"
		}
		//c.Data["IsClassifyChecked"] = isClassifyChecked
		c.Data["ClaaifyBox"] = classifyBox
		c.TplName = "blog_edit.html"
	} else {
		c.Redirect("/", 302)
	}
}
