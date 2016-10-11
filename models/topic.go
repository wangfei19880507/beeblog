package models

import (
	"github.com/astaxie/beego/orm"
)

type Topic struct {
	Id                int64
	UserId            int64
	UserName          string
	CategoryId        int64
	CategoryName      string
	TopicTitle        string
	TopicContent      string `orm:"size(5000)"`
	Labels            string
	Attachment        string
	Views             int64  `orm:"index"`
	Created           string `orm:"index"`
	Updated           string `orm:"index"`
	CommentTime       string `orm:"index"`
	CommentCount      int64
	CommentLastUserId int64
}

func (topic *Topic) TableName() string {
	return TableName("topic")
}

func (topic *Topic) TableEngine() string {
	return "INNODB DEFAULT CHARSET=utf8"
}

func (topic *Topic) Insert() error {
	if _, err := orm.NewOrm().Insert(topic); err != nil {
		return err
	}
	return nil
}

func (topic *Topic) Read(fields ...string) error {
	if err := orm.NewOrm().Read(topic, fields...); err != nil {
		return err
	}
	return nil
}

func (topic *Topic) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(topic, fields...); err != nil {
		return err
	}
	return nil
}

func (topic *Topic) Delete() error {
	if _, err := orm.NewOrm().Delete(topic); err != nil {
		return err
	}
	return nil
}

func (topic *Topic) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(topic)
}
