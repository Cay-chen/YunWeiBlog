package routers

import (
	"YunWeiBlog/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/:id", &controllers.ApiController{})
	beego.Router("/edit", &controllers.EditController{})
	beego.Router("/blog/:id", &controllers.BlogController{})
	beego.Router("/blog_list", &controllers.ListController{})
	beego.Router("/login", &controllers.LoginController{})

}
