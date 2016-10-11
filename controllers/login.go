package controllers

import (
	// "beeblog/models"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	if this.Input().Get("exit") == "true" {
		this.Ctx.SetCookie("username", "", -1, "/")
		this.Ctx.SetCookie("password", "", -1, "/")

		this.Redirect("/login", 302) // --
		return
	}

	this.TplName = "login.html"
}

func (this *LoginController) Post() {
	username := this.Input().Get("username")
	password := this.Input().Get("password")
	autoLogin := (this.Input().Get("autoLogin") == "on")

	user, err := GetUser(username, password)
	if err == nil {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}
		this.Ctx.SetCookie("username", username, maxAge, "/")
		this.Ctx.SetCookie("password", password, maxAge, "/")
		this.Redirect("/", 302)

		CurrentUser = user
		return
	}

	this.Redirect("/login", 302)
}
