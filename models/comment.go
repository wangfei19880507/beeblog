package models

import (
	"github.com/astaxie/beego/orm"
)

type Comment struct {
	Id             int64
	UserId         int64
	UserName       string
	TopicId        int64
	TopicTitle     string
	CommentContent string `orm:"size(1000)"`
	Created        string `orm:"index"`
}

func (comment *Comment) TableName() string {
	return TableName("comment")
}

func (comment *Comment) TableEngine() string {
	return "INNODB DEFAULT CHARSET=utf8"
}

func (comment *Comment) Insert() error {
	if _, err := orm.NewOrm().Insert(comment); err != nil {
		return err
	}
	return nil
}

func (comment *Comment) Read(fields ...string) error {
	if err := orm.NewOrm().Read(comment, fields...); err != nil {
		return err
	}
	return nil
}

func (comment *Comment) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(comment, fields...); err != nil {
		return err
	}
	return nil
}

func (comment *Comment) Delete() error {
	if _, err := orm.NewOrm().Delete(comment); err != nil {
		return err
	}
	return nil
}

func (comment *Comment) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(comment)
}
