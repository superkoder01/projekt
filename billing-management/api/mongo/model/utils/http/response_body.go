package http

type ResponseBody[T any] struct {
	Amount   int64 `json:"amount"`
	Elements []T   `json:"elements"`
}
