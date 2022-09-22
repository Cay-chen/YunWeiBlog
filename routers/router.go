package routers

import (
	"YunWeiBlog/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.ListController{})
	beego.Router("/api/:id", &controllers.ApiController{})
	beego.Router("/edit", &controllers.EditController{})
	beego.Router("/blog/:id", &controllers.BlogController{})
	/*	beego.Router("/blog_list", &controllers.ListController{})
	 */beego.Router("/my", &controllers.MyController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})
	beego.Router("/test", &controllers.TestController{})
	beego.Router("/blog_manage", &controllers.BlogManageController{})
	beego.Router("/exit", &controllers.ExitControoler{})

}
