package controllers

import (
	"YunWeiBlog/models"
	"YunWeiBlog/models/dao"
	"encoding/json"
	"fmt"
	"math/rand"
	"path"
	"strconv"
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
		/**
		上传图片
		*/
		case "upload_img":
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
			Success.Data.Url = "/static/upload/" + timeString + randString + ext
			Success.Data.Alt = "/static/upload/" + timeString + randString + ext
			Success.Data.Href = "/static/upload/" + timeString + randString + ext
			//返回图片地址
			res, _ := json.Marshal(Success)
			c.Ctx.WriteString(string(res))
			break
		case "delete_blog":
			var backRes dao.BackResult
			blogId, _ := c.GetInt("blogId")
			isSu, err := models.DeleteBlog(blogId, c.User.UserId)
			if err != nil {
				backRes.State = -1
				backRes.Msg = err.Error()
			} else {
				if isSu {
					backRes.State = 0
					backRes.Msg = "删除成功！"
				} else {
					backRes.State = -1
					backRes.Msg = "删除失败！"
				}
			}
			res, _ := json.Marshal(backRes)
			c.Ctx.WriteString(string(res))
			break

		}
	} else {
		var jsonBack dao.JsonResult
		jsonBack.Code = -1
		jsonBack.Msg = "都没有登陆？"
		res, _ := json.Marshal(jsonBack)
		c.Ctx.WriteString(string(res))
	}
}
