package dao

type BackResult struct {
	State int
	Msg   string
	Data  interface{}
}
type BackLayuiTableData struct {
	Code  int     `json:"code"`
	Msg   string  `json:"msg"`
	Count int     `json:"count"`
	Data  []*User `json:"data"`
}
