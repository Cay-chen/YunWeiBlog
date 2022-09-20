package dao

type BlogList struct {
	Pages int         `json:"pages"`
	Data  []*BlogInfo `json:"data"`
}
