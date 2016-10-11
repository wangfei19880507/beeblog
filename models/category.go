package models

import (
	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id              int64
	UserId          int64
	UserName        string
	CategoryName    string
	Created         string `orm:"index"`
	Views           int64  `orm:"index"`
	TopicTime       string `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

func (category *Category) TableName() string {
	return TableName("category")
}

func (category *Category) TableEngine() string {
	return "INNODB DEFAULT CHARSET=utf8"
}

func (category *Category) Insert() error {
	if _, err := orm.NewOrm().Insert(category); err != nil {
		return err
	}
	return nil
}

func (category *Category) Read(fields ...string) error {
	if err := orm.NewOrm().Read(category, fields...); err != nil {
		return err
	}
	return nil
}

func (category *Category) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(category, fields...); err != nil {
		return err
	}
	return nil
}

func (category *Category) Delete() error {
	if _, err := orm.NewOrm().Delete(category); err != nil {
		return err
	}
	return nil
}

func (category *Category) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(category)
}
