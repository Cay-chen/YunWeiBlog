package controllers

import (
	"YunWeiBlog/models/dao"
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
	IsLogin bool
	/*	IsSuperAdmin bool
		IsAdmin      bool
		IsLmAdmin    bool*/
	User *dao.User
}

/**
判断是否登录
*/
func (c *BaseController) Prepare() {
	loginUser := c.GetSession("loginUser")
	if loginUser != nil {
		logs.Notice(loginUser.(string))
		res := &dao.User{}
		err := json.Unmarshal([]byte(loginUser.(string)), &res)
		if err != nil {
			c.IsLogin = false
			return
		} else {
			c.User = res
			c.IsLogin = true
			c.Data["IsLogin"] = true
			c.Data["User"] = res
			switch res.UserGrade {
			case 0:
				c.Data["IsSuperAdmin"] = true
				logs.Notice("IsSuperAdmin")
				break
			case 1:
				c.Data["IsAdmin"] = true
				break
			case 2:
				c.Data["IsLmAdmin"] = true
				break

			}
		}
		/*		evnV,_:= beego.AppConfig.String("evnV")
				if evnV=="1" {
					c.Data["Evn"] = "生产环境"
				}else {
					c.Data["Evn"] = "测试环境"
				}*/
	} else {
		c.IsLogin = false
	}
}
