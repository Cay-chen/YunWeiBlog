package controllers

type MyController struct {
	BaseController
}

func (c *MyController) Get() {
	c.TplName = "layout.html"

	/*	if c.IsLogin{
			c.TplName="layout.html"
		}else {
			c.Redirect("/",302)
		}*/
}
