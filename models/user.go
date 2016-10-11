package models

import (
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id         int64
	UserName   string
	PassWord   string
	Email      string
	Registered string `orm:"index"`
}

func (user *User) TableName() string {
	return TableName("user")
}

func (user *User) TableEngine() string {
	return "INNODB DEFAULT CHARSET=utf8"
}

func (user *User) Insert() error {
	if _, err := orm.NewOrm().Insert(user); err != nil {
		return err
	}
	return nil
}

func (user *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(user, fields...); err != nil {
		return err
	}
	return nil
}

func (user *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(user, fields...); err != nil {
		return err
	}
	return nil
}

func (user *User) Delete() error {
	if _, err := orm.NewOrm().Delete(user); err != nil {
		return err
	}
	return nil
}

func (user *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(user)
}
