package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var CurrentUser = new(models.User)
var TopicTypes = make([]*models.TopicType, 0)

func init() {
	topictype := new(models.TopicType)
	topictypes := []models.TopicType{
		{TopicTypeName: "TopicTypeName1"},
		{TopicTypeName: "TopicTypeName2"},
		{TopicTypeName: "TopicTypeName3"},
		{TopicTypeName: "TopicTypeName4"},
		{TopicTypeName: "TopicTypeName5"},
	}

	err := topictype.InsertMulti(topictypes)
	if err != nil {
		beego.Error(err)
	}

	_, err = topictype.Query().All(&TopicTypes)
	if err != nil {
		beego.Error(err)
	}
}

func checkAccount(ctx *context.Context) bool {
	cookie, err := ctx.Request.Cookie("username")
	if err != nil {
		return false
	}
	username := cookie.Value

	cookie, err = ctx.Request.Cookie("password")
	if err != nil {
		return false
	}
	password := cookie.Value
	_, err = GetUser(username, password)

	return err == nil
}

func GetUser(username, password string) (*models.User, error) {
	user := new(models.User)
	qs := user.Query()
	err := qs.Filter("user_name", username).Filter("pass_word", password).One(user)

	return user, err
}
