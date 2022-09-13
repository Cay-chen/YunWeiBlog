package dao

type User22 struct {
	Id           int    `orm:"auto"`
	IdName       string `orm:"size(50)"`
	UserName     string `orm:"size(100)"`
	UserPassword string
}
