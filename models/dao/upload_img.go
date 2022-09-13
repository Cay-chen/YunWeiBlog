package dao

type SuccessJson struct {
	Errno int `json:"errno"`
	Data  struct {
		Url  string `json:"url"`
		Alt  string `json:"alt"`
		Href string `json:"href"`
	} `json:"data"`
}

type FailJson struct {
	Errno   int    `json:"errno"`
	Message string `json:"message"`
}
type JsonResult struct {
	Code  int         `json:"code"`  // 响应编码：0成功 401请登录 403无权限 500错误
	Msg   string      `json:"msg"`   // 消息提示语
	Data  interface{} `json:"data"`  // 数据对象
	Count int64       `json:"count"` // 记录总数
}
