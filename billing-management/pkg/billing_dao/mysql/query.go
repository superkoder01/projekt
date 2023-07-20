package mysql

type Query struct {
	Limit  int         `json:"limit,omitempty"`
	Offset int         `json:"offset,omitempty"`
	Order  string      `json:"order,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}
