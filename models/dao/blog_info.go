package dao

type BlogInfo struct {
	BlogId             int `orm:"auto"`
	BlogTitle          string
	BlogBrief          string
	BlogCreateTime     string
	BlogCreateUser     string
	BlogContent        string
	BlogImgUrl         string
	BlogVisibleType    int
	BlogClassifyType   string
	BlogLastModifyTime string
	BlogReadCount      int
	BlogState          int
}
