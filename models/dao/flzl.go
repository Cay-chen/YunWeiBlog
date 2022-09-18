package dao

type Flzl struct {
	ZlId   int    `orm:"auto"`
	ZlName string `orm:"size(50)"`
}
