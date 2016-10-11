package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	this.Data["IsLogin"] = true
	this.Data["IsTopic"] = true

	topics, err := GetAllTopics("", "", false)
	if err != nil {
		beego.Error(err)
	}
	this.Data["Topics"] = topics

	this.TplName = "topic.html"
}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	this.Data["IsLogin"] = true
	this.Data["IsTopic"] = true

	tid := this.Input().Get("tid")
	topicTitle := this.Input().Get("topictitle")
	label := this.Input().Get("label")
	topicContent := this.Input().Get("content-markdown-doc")
	categoryName := this.Input().Get("categoryname")

	_, fh, err := this.GetFile("attachment")
	if err != nil {
		beego.Error(err)
	}

	var attachment string
	if fh != nil {
		attachment = fh.Filename
		beego.Info(attachment)
		err = this.SaveToFile("attachment", path.Join("attachment", attachment))
		if err != nil {
			beego.Error(err)
		}
	}

	if len(tid) == 0 {
		err = AddTopic(topicTitle, categoryName, label, topicContent, attachment)
	} else {
		err = ModifyTopic(tid, topicTitle, categoryName, label, topicContent, attachment)
	}
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)
}

func (this *TopicController) Add() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	this.Data["IsLogin"] = true
	this.Data["IsTopic"] = true

	var err error
	this.Data["Categories"], err = GetAllCategories()
	if err != nil {
		beego.Error(err)
	}

	this.TplName = "topic_add.html"
}

func (this *TopicController) Delete() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	err := DeleteTopic(this.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)
}

func (this *TopicController) Modify() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	this.Data["IsLogin"] = true
	this.Data["IsTopic"] = true

	var err error
	this.Data["Categories"], err = GetAllCategories()
	if err != nil {
		beego.Error(err)
	}

	tid := this.Input().Get("tid")
	topic, err := GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Tid"] = tid
	this.Data["Topic"] = topic

	this.TplName = "topic_modify.html"
}

func (this *TopicController) View() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	this.Data["IsLogin"] = true
	this.Data["IsTopic"] = true

	reqUrl := this.Ctx.Request.RequestURI
	i := strings.LastIndex(reqUrl, "/")
	tid := reqUrl[i+1:]
	topic, err := GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["Labels"] = strings.Split(topic.Labels, " ")

	comments, err := GetAllComments(tid)
	if err != nil {
		beego.Error(err)
		return
	}
	this.Data["Comments"] = comments

	this.TplName = "topic_view.html"
}

func AddTopic(topicTitle, categoryName, label, topicContent, attachment string) error {
	label = "$" + strings.Join(strings.Split(label, " "), "#$") + "#"

	topic := &models.Topic{
		UserId:       CurrentUser.Id,
		UserName:     CurrentUser.UserName,
		CategoryName: categoryName,
		TopicTitle:   topicTitle,
		TopicContent: topicContent,
		Labels:       label,
		Attachment:   attachment,
		Created:      time.Now().Format("2006-01-02 15:04:05"),
		Updated:      time.Now().Format("2006-01-02 15:04:05"),
		CommentTime:  time.Date(1, 1, 1, 0, 0, 0, 0, time.UTC).Format("2006-01-02 15:04:05"),
	}

	err := topic.Insert()
	if err != nil {
		return err
	}

	cate := new(models.Category)
	qs := cate.Query()
	err = qs.Filter("category_name", categoryName).One(cate)
	if err == nil {
		cate.TopicTime = time.Now().Format("2006-01-02 15:04:05")
		cate.TopicCount++
		cate.TopicLastUserId = CurrentUser.Id
		err = cate.Update()
	}

	return err
}

func DeleteTopic(tid string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	/*	删除 User 该 Topic 下的所有 Comment - DeleteComment()	*/

	var oldCate string
	topic := &models.Topic{Id: tidNum}
	if topic.Read() == nil {
		oldCate = topic.CategoryName
		err = topic.Delete()
		if err != nil {
			return err
		}
	}
	if len(oldCate) > 0 {
		cate := new(models.Category)
		qs := cate.Query()
		err = qs.Filter("category_name", oldCate).One(cate)
		if err == nil {
			cate.TopicCount--
			err = cate.Update()
		}
	}

	return err
}

func GetTopic(tid string) (*models.Topic, error) {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return nil, err
	}

	topic := new(models.Topic)
	qs := topic.Query()
	err = qs.Filter("id", tidNum).One(topic)
	if err != nil {
		return nil, err
	}

	topic.Views++
	err = topic.Update()

	topic.Labels = strings.Replace(
		strings.Replace(topic.Labels, "#", " ", -1),
		"$",
		"",
		-1)

	return topic, nil
}

func GetAllTopics(categoryName, label string, isDesc bool) (topics []*models.Topic, err error) {
	topics = make([]*models.Topic, 0)

	topic := new(models.Topic)
	qs := topic.Query()

	if isDesc {
		if len(categoryName) > 0 {
			qs = qs.Filter("category_name", categoryName)
		}
		if len(label) > 0 {
			qs = qs.Filter("labels__contains", "$"+label+"#")
		}
		_, err = qs.OrderBy("-created").All(&topics)

	} else {
		_, err = qs.All(&topics)
	}

	return topics, err
}

func ModifyTopic(tid, topicTitle, categoryName, label, topicContent, attachment string) error {
	tidNum, err := strconv.ParseInt(tid, 10, 64)
	if err != nil {
		return err
	}

	label = "$" + strings.Join(strings.Split(label, " "), "#$") + "#"

	var oldCate, oldAttach string
	topic := &models.Topic{Id: tidNum}

	if topic.Read() == nil {
		oldCate = topic.CategoryName
		oldAttach = topic.Attachment
		topic.CategoryName = categoryName
		topic.TopicTitle = topicTitle
		topic.TopicContent = topicContent
		topic.Labels = label
		topic.Attachment = attachment
		topic.Updated = time.Now().Format("2006-01-02 15:04:05")

		err = topic.Update()
		if err != nil {
			return err
		}
	}

	if categoryName != oldCate {
		cate := new(models.Category)
		qs := cate.Query()
		err = qs.Filter("category_name", oldCate).One(cate)
		if err == nil {
			cate.TopicCount--
			err = cate.Update()
		}

		err = qs.Filter("category_name", categoryName).One(cate)
		if err == nil {
			cate.TopicCount++
			err = cate.Update()
		}
	}
	// 删除旧的附件
	if len(oldAttach) > 0 {
		os.Remove(path.Join("attachment", oldAttach))
	}
	cate := new(models.Category)
	qs := cate.Query()
	err = qs.Filter("category_name", categoryName).One(cate)
	if err == nil {
		cate.TopicCount++
		err = cate.Update()
	}

	return nil
}
