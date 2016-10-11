package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302) // --
		return
	}

	this.Data["IsLogin"] = true
	this.Data["IsHome"] = true

	topics, err := GetAllTopics(
		this.Input().Get("cate"),
		this.Input().Get("label"),
		true)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics

	// 获取博客分类
	// categories, err := GetAllCategories()
	// if err != nil {
	// 	beego.Error(err)
	// }
	// this.Data["Categories"] = categories
	this.Data["TopicTypes"] = TopicTypes

	this.TplName = "home.html"
}
