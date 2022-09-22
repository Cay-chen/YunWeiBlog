package controllers

type ExitControoler struct {
	BaseController
}

func (c *ExitControoler) Get() {
	c.DelSession("loginUser")
	c.Redirect("/", 302)

}
