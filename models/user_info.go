package models

import (
	"YunWeiBlog/models/dao"
	"encoding/json"
	"github.com/beego/beego/v2/client/orm"
)

func GetUsersList(limit, page int) (string, error) {
	var backRes dao.BackLayuiTableData
	var usersList []*dao.User
	start := (page - 1) * limit
	o := orm.NewOrm()
	qs := o.QueryTable("User")
	_, err := qs.Limit(limit, start).All(&usersList, "UserId", "UserName", "RegisterTime", "UserGrade", "UserState")
	if err != nil {
		return "", err
	}
	count, err1 := qs.Count()
	if err1 != nil {
		return "", err1
	}
	backRes.Data = usersList
	backRes.Code = 0
	backRes.Count = int(count)
	backRes.Msg = "成功！"
	backRes1, err2 := json.Marshal(backRes)
	if err2 != nil {
		return "", err2
	}
	return string(backRes1), nil
}
