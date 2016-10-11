package controllers

import (
	"beeblog/models"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type CommentController struct {
	beego.Controller
}

func (this *CommentController) Add() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	topicid := this.Input().Get("topicid")
	commentContent := this.Input().Get("commentcontent")

	err := AddComment(topicid, commentContent)
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic/view/"+topicid, 302)
}

func (this *CommentController) Delete() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	topicid := this.Input().Get("topicid")

	err := DeleteComment(this.Input().Get("cid"))
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic/view/"+topicid, 302)
}

func AddComment(topicid, commentContent string) error {
	tidNum, err := strconv.ParseInt(topicid, 10, 64)
	if err != nil {
		return err
	}

	comment := &models.Comment{
		TopicId:        tidNum,
		UserName:       CurrentUser.UserName,
		CommentContent: commentContent,
		Created:        time.Now().Format("2006-01-02 15:04:05"),
	}

	err = comment.Insert()
	if err != nil {
		return err
	}

	topic := &models.Topic{Id: tidNum}
	if topic.Read() == nil {
		topic.CommentTime = time.Now().Format("2006-01-02 15:04:05")
		topic.CommentCount++
		topic.CommentLastUserId = CurrentUser.Id
		err = topic.Update()
	}

	return err
}

func DeleteComment(cid string) error {
	cidNum, err := strconv.ParseInt(cid, 10, 64)
	if err != nil {
		return err
	}

	var tidNum int64
	comment := &models.Comment{Id: cidNum}
	if comment.Read() == nil {
		tidNum = comment.TopicId
		err = comment.Delete()
		if err != nil {
			return err
		}
	}

	comments := make([]*models.Comment, 0)
	qs := comment.Query()
	_, err = qs.Filter("topic_id", tidNum).OrderBy("-created").All(&comments)
	if err != nil {
		return err
	}

	topic := &models.Topic{Id: tidNum}
	if topic.Read() == nil {
		topic.CommentTime = comments[0].Created
		topic.CommentCount = int64(len(comments))
		err = topic.Update()
	}

	return err
}

func GetAllComments(topicid string) (comments []*models.Comment, err error) {
	tidNum, err := strconv.ParseInt(topicid, 10, 64)
	if err != nil {
		return nil, err
	}

	var comment models.Comment
	qs := comment.Query()
	_, err = qs.Filter("topic_id", tidNum).All(&comments)

	return comments, err
}
