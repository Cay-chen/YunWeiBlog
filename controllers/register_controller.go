package controllers

import (
	"YunWeiBlog/models/dao"
	"YunWeiBlog/models/utils"
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"strings"
)

type RegisterController struct {
	beego.Controller
}

func (c *RegisterController) Post() {
	var user dao.User
	var jsonBack dao.JsonResult
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &user)
	user.UserPassword = utils.Md5(user.UserPassword)
	user.UserGrade = 4
	if err != nil {
		jsonBack.Code = -1
		jsonBack.Msg = err.Error()
	} else {
		o := orm.NewOrm()
		_, err1 := o.Insert(&user)
		if err1 != nil {
			if strings.Contains(err1.Error(), "Duplicate entry") {
				jsonBack.Code = -1
				jsonBack.Msg = "该用户名已注册！"
			} else {
				jsonBack.Code = -1
				jsonBack.Msg = "用户注册失败！"
			}
		} else {
			jsonBack.Code = 0
			jsonBack.Msg = "注册成功！"
		}
	}
	res, _ := json.Marshal(jsonBack)
	c.Ctx.WriteString(string(res))
}
