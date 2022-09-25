package controllers

type TestController struct {
	BaseController
}

func (c *TestController) Get() {
	c.TplName = "iframe_blog_manage.html"

	/*	if c.IsLogin{
			c.TplName="blog_layout.html"
		}else {
			c.Redirect("/",302)
		}*/
}
