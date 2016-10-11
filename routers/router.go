package routers

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"
	"os"
)

func init() {
	beego.Router("/", &controllers.HomeController{})

	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/register", &controllers.RegisterController{})

	beego.Router("/category", &controllers.CategoryController{})

	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})

	beego.Router("/comment/add", &controllers.CommentController{}, "*:Add")
	beego.Router("/comment/delete", &controllers.CommentController{}, "*:Delete")

	os.Mkdir("attachment", os.ModePerm)
	beego.Router("/attachment/:all", &controllers.AttachController{})
}
