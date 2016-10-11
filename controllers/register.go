package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
	"time"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get() {
	this.TplName = "register.html"
}

func (this *RegisterController) Post() {
	username := this.Input().Get("username")
	password := this.Input().Get("password")
	email := this.Input().Get("email")

	err := AddUser(username, password, email)
	if err != nil {
		beego.Error(err)
		return
	}

	this.Redirect("/login", 302)
}

func AddUser(username, password, email string) error {
	user := &models.User{
		UserName:   username,
		PassWord:   password,
		Email:      email,
		Registered: time.Now().Format("2006-01-02 15:04:05"),
	}

	qs := user.Query()
	err := qs.Filter("user_name", username).One(user)
	if err == nil {
		return err
	}

	err = user.Insert()
	if err != nil {
		return err
	}

	return nil
}
