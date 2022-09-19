package controllers

import (
	"YunWeiBlog/models/dao"
	"YunWeiBlog/models/utils"
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	"strings"
)

// LoginController /*登录控制器*/
type LoginController struct {
	BaseController
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	if c.IsLogin {
		c.Redirect("/", 302)
	} else {
		var user dao.User
		var jsonBack dao.JsonResult
		data := c.Ctx.Input.RequestBody
		err := json.Unmarshal(data, &user)
		user.UserPassword = utils.Md5(user.UserPassword)
		if err != nil {
			jsonBack.Code = -1
			jsonBack.Msg = err.Error()
		} else {
			o := orm.NewOrm()
			qs := o.QueryTable("user")
			errO := qs.Filter("UserName", user.UserName).Filter("UserPassword", user.UserPassword).One(&user)
			if errO != nil {
				if strings.Contains(errO.Error(), "no row found") {
					jsonBack.Code = -1
					jsonBack.Msg = "账号或密码错误！"
				} else {
					jsonBack.Code = -1
					jsonBack.Msg = errO.Error()
				}
			} else {
				if user.UserId > 0 {
					user.UserPassword = ""
					jsonUser, _ := json.Marshal(user)
					errSe := c.SetSession("loginUser", string(jsonUser))
					if errSe != nil {
						jsonBack.Code = -1
						jsonBack.Msg = "登陆失败！"
					} else {
						jsonBack.Code = 0
						jsonBack.Msg = "登陆成功！"
					}

				} else {
					if c.GetSession("loginUser") != nil {
						c.DestroySession()
					}
					jsonBack.Code = -1
					jsonBack.Msg = "登陆失败！"
				}
			}
		}
		res, _ := json.Marshal(jsonBack)
		c.Ctx.WriteString(string(res))
	}

}
