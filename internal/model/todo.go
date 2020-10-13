package model

type Todo struct {
	Todo string `json:"todo" form:"todo" query:"todo"`
	Done bool   `json:"done" form:"done" query:"done"`
}
