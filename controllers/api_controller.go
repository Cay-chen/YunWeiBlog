package controllers

import (
	"YunWeiBlog/models/dao"
	"YunWeiBlog/util"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"math/rand"
	"path"
	"strconv"
	"time"
)

type ApiController struct {
	beego.Controller
}

func (c *ApiController) Get() {

}

func (c *ApiController) Post() {
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
		blogInfo.BlogCreateTime = util.GetNowTime()
		blogInfo.BlogLastModifyTime = util.GetNowTime()
		blogInfo.BlogCreateUser = "2"
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
		var blogInfo dao.BlogInfo
		data := c.Ctx.Input.RequestBody
		logs.Error(string(data))
		err := json.Unmarshal(data, &blogInfo)
		if err != nil {
			logs.Error("json.Unmarshal is err:", err.Error())
		}
		blogInfo.BlogCreateTime = util.GetNowTime()
		blogInfo.BlogLastModifyTime = util.GetNowTime()
		blogInfo.BlogCreateUser = "2"
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
	}
}
