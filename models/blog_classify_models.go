package models

import (
	"YunWeiBlog/models/dao"
	"github.com/beego/beego/v2/client/orm"
)

func GetClassifyList() []dao.BlogClassify {
	var classify []dao.BlogClassify
	//isClassifyChecked :=""
	o := orm.NewOrm()
	qs := o.QueryTable("BlogClassify")
	qs.Filter("ClassifyState", 0).All(&classify)
	return classify
}
