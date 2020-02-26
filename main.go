package main

import (
	"github.com/astaxie/beego/orm"
	"hello/models"
	_ "hello/routers"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

func init()  {
	// register model
	orm.RegisterModel(new(models.User))

	// set default database
	orm.RegisterDataBase("default", "mysql", "root:root@/majoo?charset=utf8", 30)
}

func main() {
	beego.Run()
}

