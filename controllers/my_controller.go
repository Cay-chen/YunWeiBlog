package controllers

type MyController struct {
	BaseController
}

func (c *MyController) Get() {
	if c.IsLogin {
		wd := c.GetString("wd")
		xs, err := c.GetInt("xs")
		if err != nil {
			c.Data["Xs"] = 9
		} else {
			c.Data["Xs"] = xs
		}
		c.Data["Wd"] = wd
		c.TplName = "blog_layout.html"
	} else {
		c.Redirect("/error/600", 302)
	}

}
func (c *MyController) Post() {

}
