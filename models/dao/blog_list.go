package dao

type ApiBlogList struct {
	Pages int         `json:"pages"`
	Data  []*BlogInfo `json:"data"`
}
