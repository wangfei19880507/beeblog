package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type CategoryController struct {
	beego.Controller
}

func (this *CategoryController) Get() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	this.Data["IsLogin"] = true
	this.Data["IsCategory"] = true

	op := this.Input().Get("op")
	switch op {
	case "add":
		categoryName := this.Input().Get("categoryname")
		if len(categoryName) == 0 {
			break
		}

		err := AddCategory(categoryName)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/category", 302)
		return

	case "del":
		id := this.Input().Get("id")
		if len(id) == 0 {
			break
		}

		err := DeleteCategory(id)
		if err != nil {
			beego.Error(err)
		}

		this.Redirect("/category", 302)
		return
	}

	var err error
	this.Data["Categories"], err = GetAllCategories()
	if err != nil {
		beego.Error(err)
	}

	this.TplName = "category.html"
}

func AddCategory(categoryName string) error {
	cate := &models.Category{
		UserId:       CurrentUser.Id,
		UserName:     CurrentUser.UserName,
		CategoryName: categoryName,
		Created:      time.Now().Format("2006-01-02 15:04:05"),
		TopicTime:    time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02 15:04:05"),
	}

	category := new(models.Category)
	err := category.Query().Filter("user_name", CurrentUser.UserName).Filter("category_name", categoryName).One(cate)
	if err == nil {
		return err
	}

	err = cate.Insert()

	return err
}

func DeleteCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	/*	删除 User 该 Category 下的所有 Topic - DeleteTopic()
		topics := make([]*models.Topic, 0)
		topic := new(models.Topic)
		_, err = topic.Query().Filter("category")
	*/

	category := &models.Category{Id: cid}
	err = category.Delete()

	return err
}

func GetAllCategories() ([]*models.Category, error) {
	categories := make([]*models.Category, 0)

	category := new(models.Category)
	_, err := category.Query().All(&categories)

	return categories, err
}
