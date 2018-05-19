package model

import (
	"fastgo/model/user_model"
	"fastgo/module/config"
	"fmt"

	beego_config "github.com/astaxie/beego/config"

	"strings"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var confDbStr string

func init() {
	conf, err := config.Reader("database.conf")
	if err != nil {
		fmt.Println("------", err)
	}
	host := conf.String("mysql::host")
	port := conf.String("mysql::port")

	dbUsername := "db_username"
	dbPassword := "db_password"

	confDbStr = dbUsername + ":" + dbPassword + "@(" + host + ":" + port + ")/{{DB_NAME}}?charset=utf8"

	//注册db
	registerDB()

	//注册model
	registerModel()

	// create table
	// orm.RunSyncdb("default", false, true)

	//在测试环境打开mysql 调试
	iniconf, err := beego_config.NewConfig("ini", "config/app.conf")
	if err != nil {
		fmt.Println("------", err)
	} else {
		if iniconf.String("runmode") == "dev" {
			orm.Debug = true
		}
	}
}

func registerDB() {
	// set default database
	orm.RegisterDataBase("default", "mysql", getDbStr(user_model.USER_DB), 30)
}

func registerModel() {
	// register model
	orm.RegisterModel(new(user_model.User))
}

func getDbStr(dbname string) string {
	return strings.Replace(confDbStr, "{{DB_NAME}}", dbname, -1)
}
