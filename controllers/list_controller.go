package controllers

type ListController struct {
	BaseController
}

func (c *ListController) Get() {
	/*	var blogInfo []*dao.BlogInfo
		wd:= c.GetString("wd")
		cond :=orm.NewCondition()
		cond1 := cond.And("BlogVisibleType", 0).And("BlogTitle__icontains",wd).Or("BlogBrief__icontains",wd)
		o := orm.NewOrm()
		qs := o.QueryTable("blogInfo")
		num, err := qs.SetCond(cond1).All(&blogInfo)

		//num, err := qs.Filter("BlogVisibleType", 0).All(&blogInfo)
		if err != nil {
			logs.Error(err.Error())
		}
		logs.Notice(num)
		c.Data["BlogInfoList"] = blogInfo*/
	c.TplName = "blog_list.html"
}
