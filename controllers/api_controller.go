package controllers

import (
	"YunWeiBlog/models/dao"
	"YunWeiBlog/models/utils"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"math/rand"
	"path"
	"strconv"
	"strings"
	"time"
)

type ApiController struct {
	BaseController
}

func (c *ApiController) Get() {

}

func (c *ApiController) Post() {
	if c.IsLogin {
		id := c.Ctx.Input.Param(":id")
		switch id {
		case "upload_img":
			fmt.Println("aaa")
			var Success dao.SuccessJson
			var Fail dao.FailJson
			//获取上传的文件
			file, head, er := c.GetFile("custom-fileName")
			if er != nil {
				Fail.Errno = 1
				Fail.Message = "上传出错"
				res, _ := json.Marshal(Fail)
				fmt.Println(string(res))
				c.Ctx.WriteString(string(res))
				return
			}
			ext := path.Ext(head.Filename)
			//限制上传文件类型
			var FileAllow map[string]bool = map[string]bool{
				".jpg":  true,
				".txt":  true,
				".png":  true,
				".PNG":  true,
				".go":   true,
				".webp": true,
			}
			if _, ok := FileAllow[ext]; !ok {
				Fail.Errno = 1
				Fail.Message = "文件后缀名不符合上传要求"
				res, _ := json.Marshal(Fail)
				fmt.Println(string(res))
				c.Ctx.WriteString(string(res))
				return
			}
			defer file.Close()
			//获取当前时间
			timeString := time.Now().Format("20060102150405")
			//生产一个四位数转string
			randString := strconv.FormatInt(int64(rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000)), 10)
			//保持图片
			err := c.SaveToFile("custom-fileName", "static/upload/"+timeString+randString+ext)
			if err != nil {
				Fail.Errno = 1
				Fail.Message = "上传文件失败"
				res, _ := json.Marshal(Fail)
				c.Ctx.WriteString(string(res))
				return
			}
			Success.Errno = 0
			Success.Data.Url = "http://127.0.0.1/static/upload/" + timeString + randString + ext
			Success.Data.Alt = "/static/upload/" + timeString + randString + ext
			Success.Data.Href = "http://127.0.0.1/static/upload/" + timeString + randString + ext
			//返回图片地址
			res, _ := json.Marshal(Success)
			c.Ctx.WriteString(string(res))
			break
		case "edit_submit":
			var blogInfo dao.BlogInfo
			data := c.Ctx.Input.RequestBody
			logs.Error(string(data))
			err := json.Unmarshal(data, &blogInfo)
			if err != nil {
				logs.Error("json.Unmarshal is err:", err.Error())
			}
			blogInfo.BlogCreateUser = c.User.UserId
			o := orm.NewOrm()
			insert, err := o.Insert(&blogInfo)
			var jsonBack dao.JsonResult

			if err != nil {
				logs.Error(insert)
				logs.Error(err.Error())
				jsonBack.Code = -1
				jsonBack.Msg = err.Error()
			} else {
				if insert > 0 {
					jsonBack.Code = 0
					jsonBack.Msg = "乘车"
				} else {
					jsonBack.Code = -1
					jsonBack.Msg = "失败"
				}
			}
			res, err := json.Marshal(jsonBack)
			c.Ctx.WriteString(string(res))
			break
		case "sign_in":
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
						logs.Error("用户注册失败！->" + err1.Error())
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
			break
			/*		case "login":
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
								jsonBack.Code = 0
								jsonBack.Msg = "登陆成功！"
							} else {
								jsonBack.Code = -1
								jsonBack.Msg = "登陆失败！"
							}
						}
					}
					res, _ := json.Marshal(jsonBack)
					c.Ctx.WriteString(string(res))
					break*/
		}
	} else {
		var jsonBack dao.JsonResult
		jsonBack.Code = -1
		jsonBack.Msg = "都没有登陆？"
		res, _ := json.Marshal(jsonBack)
		c.Ctx.WriteString(string(res))
	}
}
