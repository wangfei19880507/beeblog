package models

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbname := beego.AppConfig.String("dbname")

	if dbport == "" {
		dbport = "3306"
	}

	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" +
		dbport + ")/" + dbname + "?charset=utf8"

	orm.RegisterDataBase("default", "mysql", dsn)

	orm.RegisterModel(new(User), new(TopicType), new(Category), new(Topic), new(Comment))

	force := false
	verbose := true
	name := "default"

	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println("models.base ====", err)
	}
}

func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func TableName(str string) string {
	return fmt.Sprintf("%s%s", beego.AppConfig.String("dbprefix"), str)
}
