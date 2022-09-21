package controllers

type TestController struct {
	BaseController
}

func (c *TestController) Get() {
	c.TplName = "blog_manage.html"

	/*	if c.IsLogin{
			c.TplName="layout.html"
		}else {
			c.Redirect("/",302)
		}*/
}
