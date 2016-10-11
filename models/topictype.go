package models

import (
	"github.com/astaxie/beego/orm"
)

type TopicType struct {
	Id            int64
	TopicTypeName string
}

func (topictype *TopicType) TableName() string {
	return TableName("topictype")
}

func (topictype *TopicType) TableEngine() string {
	return "INNODB DEFAULT CHARSET=utf8"
}

func (topictype *TopicType) Insert() error {
	if _, err := orm.NewOrm().Insert(topictype); err != nil {
		return err
	}
	return nil
}
func (topictype *TopicType) InsertMulti(topictypes []TopicType) error {
	if _, err := orm.NewOrm().InsertMulti(len(topictypes), topictypes); err != nil {
		return err
	}
	return nil
}

func (topictype *TopicType) Read(fields ...string) error {
	if err := orm.NewOrm().Read(topictype, fields...); err != nil {
		return err
	}
	return nil
}

func (topictype *TopicType) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(topictype, fields...); err != nil {
		return err
	}
	return nil
}

func (topictype *TopicType) Delete() error {
	if _, err := orm.NewOrm().Delete(topictype); err != nil {
		return err
	}
	return nil
}

func (topictype *TopicType) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(topictype)
}
