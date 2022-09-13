package main

import (
	"YunWeiBlog/models/dao"
	_ "YunWeiBlog/routers"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func main() {
	orm.Debug = true
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 1800
	err := logs.SetLogger(logs.AdapterFile, `{"filename":"logs/log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	if err != nil {
		return
	}
	beego.Run()

	/*	orm.Debug = true
		o := orm.NewOrm()
		user := dao.User{
			Id:           1,
			IdName:       "admin",
			UserName:     "管理员",
			UserPassword: "123"}
		// insert
		id, err := o.Insert(&user)
		fmt.Printf("ID: %d, ERR: %v\n", id, err)*/

	/*// update
	//user.Id = 1
	//user.UserName = "astaxie"
	num, err := o.Update(&user)
	fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// read one
	u := dao.User{Id: user.Id}
	err := o.Read(&u)
	fmt.Printf("ERR: %v\n", err)
	fmt.Printf("ERR: %v\n", u)*/

	// delete
	/*	u := dao.User{Id: user.Id}

		num, err := o.Delete(&u)
		fmt.Printf("NUM: %d, ERR: %v\n", num, err)*/

	//beego.Run()

	// 默认使用 default，你可以指定为其他数据库
	/*orm.Debug = true
	o := orm.NewOrm()
	profile := new(dao.Profile)
	profile.Age = 33
	user := new(dao.User1)

	user.Profile = profile
	user.Name = "slene123"
	fmt.Println(o.Insert(profile))
	fmt.Println(o.Insert(user))*/
	//o := orm.NewOrm()
	//ids := []int{1, 2, 3}
	//var user []dao.User1

	//var maps []orm.Params
	//	var list orm.ParamsList
	/*	res := make(orm.Params)

		//num, err :=o.Raw("SELECT id FROM user1 WHERE id IN (?, ?, ?)", ids).QueryRows(&user)
		//num, err :=o.Raw("SELECT * FROM user1 WHERE id IN (?, ?, ?)", ids).ValuesFlat(&list)
		num, err :=o.Raw("SELECT id, name FROM user1").RowsToMap(&res, "id", "name")
		if err == nil {
			fmt.Println("user nums: ", num)
		}
		fmt.Println(res[])*/
	/*	for i:=0;i<int(num);i++{
		fmt.Println(list)
	}*/

	/*	res, err := o.Raw("UPDATE user1 SET name = ? where id=?", "your",1).Exec()
		if err == nil {
			num, _ := res.RowsAffected()
			fmt.Println("mysql row affected nums: ", num)
		}
	*/

	/*	var list orm.ParamsList
		num, err := o.Raw("SELECT id FROM user WHERE id < ?", 10).ValuesFlat(&list)
		if err == nil && num > 0 {
			fmt.Println(list) // []{"1","2","3",...}
		}*/

}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dataSource, _ := beego.AppConfig.String("mysql")
	logs.Debug(dataSource)
	//,"root:root@tcp(192.200.102.56:3306)/yunwei?charset=utf8"
	orm.RegisterDataBase("default", "mysql", dataSource) // register model
	orm.SetMaxIdleConns("default", 30)
	orm.SetMaxOpenConns("default", 30)
	orm.DefaultTimeLoc = time.UTC
	orm.RegisterModel(new(dao.BlogInfo))
	//orm.RegisterModel(new(dao.User1), new(dao.Post), new(dao.Profile), new(dao.Tag))

	// create table
	//	orm.RunSyncdb("default", false, true)

}
