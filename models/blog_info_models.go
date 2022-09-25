package models

import (
	"YunWeiBlog/models/dao"
	"encoding/json"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"math"
	"strconv"
	"time"
)

func GetBlogInfo(BlogId int) (dao.BlogInfo, error) {
	o := orm.NewOrm()
	blogInfo := dao.BlogInfo{}
	err := o.QueryTable("BlogInfo").Filter("BlogId", BlogId).RelatedSel().One(&blogInfo)
	return blogInfo, err
}
func ReadAdd(BlogId int) {
	o := orm.NewOrm()
	o.QueryTable("BlogInfo").Filter("BlogId", BlogId).Update(
		orm.Params{
			"BlogReadCount": orm.ColValue(orm.ColAdd, 1),
		})
}
func GetApiBlogInfoList(wd string, page, classify, visible, state, userId int) (string, error) {
	limit := 8
	var blogInfo []*dao.BlogInfo
	var blogList dao.ApiBlogList
	start := (page - 1) * limit
	cond := orm.NewCondition()
	//var condS *orm.Condition
	condS := cond.And("BlogDelState", 0)
	condS = condS.And("BlogState", state)
	if visible != -1 {
		condS = condS.And("BlogVisibleType", visible)
	}
	if userId != -1 {
		condS = condS.And("BlogCreateUser", userId)
	}
	if classify != -1 {
		condS = condS.And("BlogClassifyType__icontains", strconv.Itoa(classify)+"|")
	}
	if wd != "" {
		condS = condS.AndCond(cond.And("BlogTitle__icontains", wd).Or("BlogBrief__icontains", wd))
	}

	//	cond1 := cond.And("BlogVisibleType", 0).And("BlogState", 0).AndCond(cond.And("BlogTitle__icontains", wd).Or("BlogBrief__icontains", wd))
	o := orm.NewOrm()
	qs := o.QueryTable("BlogInfo")
	_, err := qs.SetCond(condS).RelatedSel("User").Limit(limit, start).All(&blogInfo, "BlogId", "BlogTitle", "BlogBrief", "BlogCreateTime", "BlogCreateUser", "BlogImgUrl", "BlogClassifyType", "BlogReadCount")
	if err != nil {
		return "", err
	}
	count, err1 := qs.SetCond(condS).RelatedSel().Count()
	if err1 != nil {
		return "", err1
	}
	pages := int(math.Ceil(float64(count) / float64(limit)))
	blogList.Pages = pages
	blogList.Data = blogInfo
	backRes, err2 := json.Marshal(blogList)
	if err2 != nil {
		return "", err2
	}
	return string(backRes), err2
}
func GetNilBlogInfoList() string {
	var blogList dao.ApiBlogList
	backRes, _ := json.Marshal(blogList)
	return string(backRes)
}

func InsertAndEditBlogInfo(blogId, userId int, data []byte) (string, error) {
	var blogInfo dao.BlogInfo
	var jsonBack dao.JsonResult
	fmt.Println(userId)
	fmt.Println(string(data))
	err := json.Unmarshal(data, &blogInfo)
	if err != nil {
		return "", err
	}
	o := orm.NewOrm()
	if blogId != -1 {
		num, err2 := o.QueryTable("BlogInfo").Filter("BlogId", blogId).Filter("BlogCreateUser", userId).Update(orm.Params{
			"BlogTitle":          blogInfo.BlogTitle,
			"BlogBrief":          blogInfo.BlogBrief,
			"BlogContent":        blogInfo.BlogContent,
			"BlogImgUrl":         blogInfo.BlogImgUrl,
			"BlogVisibleType":    blogInfo.BlogVisibleType,
			"BlogClassifyType":   blogInfo.BlogClassifyType,
			"BlogState":          blogInfo.BlogState,
			"BlogLastModifyTime": time.Now(),
		})
		if err2 != nil {
			jsonBack.Code = -1
			jsonBack.Msg = err2.Error()
		} else {
			if num > 0 {
				jsonBack.Code = 0
				jsonBack.Msg = "修改成功"
			} else {
				jsonBack.Code = -1
				jsonBack.Msg = "失败"
			}
		}
	} else {
		blogInfo.BlogCreateUser = userId
		var users *dao.User
		users = new(dao.User)
		users.UserId = userId
		blogInfo.User = users
		insert, err1 := o.Insert(&blogInfo)
		if err1 != nil {
			jsonBack.Code = -1
			jsonBack.Msg = err1.Error()
		} else {
			if insert > 0 {
				jsonBack.Code = 0
				jsonBack.Msg = "发表成功"
			} else {
				jsonBack.Code = -1
				jsonBack.Msg = "失败"
			}
		}
	}
	res, err5 := json.Marshal(jsonBack)
	return string(res), err5
}
func DeleteBlog(blogId, userId int) (bool, error) {
	o := orm.NewOrm()
	num, err := o.QueryTable("BlogInfo").Filter("BlogId", blogId).Filter("BlogCreateUser", userId).Update(orm.Params{
		"BlogDelState": 1,
		"BlogDelTime":  time.Now(),
	})
	/*	num, err := o.QueryTable("BlogInfo").Filter("BlogId", blogId).Filter("BlogCreateUser",userId).Delete()
	 */
	if err != nil {
		return false, err
	}
	if num > 0 {
		return true, nil
	} else {
		er1 := fmt.Errorf("没找到删除的内容")
		return false, er1
	}

}
