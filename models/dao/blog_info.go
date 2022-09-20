package dao

import "time"

/**
表blog_info
*/
type BlogInfo struct {
	BlogId             int       `orm:"auto;description(文章唯一ID)"`
	BlogTitle          string    `orm:"type(char);size(80);description(文章标题)"`
	BlogBrief          string    `orm:"type(char);size(255);description(文章简介)"`
	BlogCreateTime     time.Time `orm:"auto_now_add;type(datetime);description(文章创建时间)"` //auto_now_add 第一次保存时才设置时间   date 是日期
	BlogCreateUser     int       `orm:"size(4);description(文章创建用户ID)"`
	BlogContent        string    `orm:"type(text);description(文章内容html)"`
	BlogImgUrl         string    `orm:"type(char);size(255);description(文章图片地址)"`
	BlogVisibleType    int       `orm:"size(1);description(文字可见类型，0全部可见，1为自己可见)"`
	BlogClassifyType   string    `orm:"type(char);description(文章分类 多个类型用|分开)"`
	BlogLastModifyTime time.Time `orm:"auto_now;type(datetime);description(文章最后一次修改时间)"` //auto_now 每次 model 保存时都会对时间自动更新
	BlogReadCount      int       `orm:"size(10);description(文章阅读量)"`
	BlogState          int       `orm:"size(1);description(文章状态 0发表 1为草稿)"`
}

/**
表user
*/
type User struct {
	UserId       int       `orm:"auto;description(用户唯一ID)"`
	UserName     string    `orm:"unique;type(char);size(20);description(用户登陆账号唯一)"`
	UserPassword string    `orm:"type(char);size(50);description(用户密码)"`
	RegisterTime time.Time `orm:"auto_now_add;type(datetime);description(账号注册时间)"`
	UserGrade    int       `orm:"size(1);default(4);description(用户等级 0为超级管理员 1为普通管理员 3位栏目管理员 4位普通用户)"`
}
