package controllers

type EditController struct {
	BaseController
}

func (c *EditController) Get() {
	if c.IsLogin {
		c.TplName = "blog_edit.html"
	} else {
		c.Redirect("/", 302)
	}
}
