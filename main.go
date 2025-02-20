package main

import (
	"beego_blog/models"
	_ "beego_blog/routers"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	models.Init()
	web.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	web.Run()
}
