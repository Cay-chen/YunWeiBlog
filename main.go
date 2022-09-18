package main

import (
	"YunWeiBlog/models/dao"
	"YunWeiBlog/models/utils"
	_ "YunWeiBlog/routers"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"strings"
	"time"
)

func main() {
	orm.Debug = true
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 600
	err := logs.SetLogger(logs.AdapterFile, `{"filename":"logs/log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	if err != nil {
		logs.Error("日志配置出错！->" + err.Error())
		os.Exit(0)
	}
	beego.BConfig.WebConfig.StaticExtensionsToGzip = []string{".css", ".js"}
	beego.Run()
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dataSource, e := beego.AppConfig.String("mysql")
	if e != nil {
		logs.Error("获取数据库配置文件失败！->" + e.Error())
		os.Exit(0)
	}
	//,"root:root@tcp(192.200.102.56:3306)/yunwei?charset=utf8"
	errR := orm.RegisterDataBase("default", "mysql", dataSource)
	if errR != nil {
		logs.Error("连接数据库失败！->" + errR.Error())
		os.Exit(0)
	} // register model
	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 30)
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterModel(new(dao.BlogInfo), new(dao.User))
	//orm.RegisterModel(new(dao.User1), new(dao.Post), new(dao.Profile), new(dao.Tag))
	isInstall, _ := beego.AppConfig.String("isInstall")
	if len(isInstall) < 1 {
		// create table
		err := orm.RunSyncdb("default", false, true)
		if err != nil {
			logs.Error("创建表失败！->" + err.Error())
			os.Exit(0)
		}
		var user dao.User
		name, _ := beego.AppConfig.String("adminName")
		password, _ := beego.AppConfig.String("adminPassword")
		user.UserName = name
		user.UserGrade = 0
		user.UserPassword = utils.Md5(password)
		if len(strings.TrimSpace(name)) < 1 {
			logs.Error("获取管理员用户名错误！")
			os.Exit(0)
		}
		if len(strings.TrimSpace(password)) < 1 {
			logs.Error("获取管理员密码名错误！")
			os.Exit(0)
		}
		o := orm.NewOrm()
		_, err1 := o.Insert(&user)
		if err1 != nil {
			logs.Error("创建管理员用户失败！->" + err1.Error())
			os.Exit(0)
		}

		file, errC := os.OpenFile("conf/app.conf", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
		if errC != nil {
			logs.Error("打开安装文件失败！->" + errC.Error())
			os.Exit(0)
		}
		_, errW := file.WriteString("isInstall=0")
		if errW != nil {
			return
		}
		/*	path := "conf/install.conf"
			_, err5 := os.OpenFile(path, os.O_APPEND|os.O_CREATE, 0644)
			if err5 != nil {
				logs.Error("创建安装文件失败！->" + err5.Error())
				return
			} else {
				logs.Notice("创建成功！")
				file, errC := os.OpenFile("conf/install.conf", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
				if errC != nil {
					logs.Error("打开安装文件失败！->" + errC.Error())
					os.Exit(0)
				}
				_, errW := file.WriteString("isInstall=0")
				if errW != nil {
					return
				}
			}*/
		/*
			errS := beego.AppConfig.Set("isInstall", "1")
			if errS != nil {
				logs.Error("修改安装状态失败！->" + errS.Error())
				os.Exit(0)
			}
			isInstall2, _ := beego.AppConfig.String("isInstall")
			logs.Notice(isInstall2)*/
	} else {
		logs.Notice("初始化完成！")

	}
}
