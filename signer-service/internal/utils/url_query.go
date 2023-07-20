package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
)

var (
	logger = logging.MustGetLogger("url_query")
)

type QueryResult struct {
	Amount   int64       `json:"amount,omitempty"`
	Elements interface{} `json:"elements,omitempty"`
}

func (c *QueryResult) String() string {
	return fmt.Sprintf("%s", *c)
}

type Query struct {
	Sort *Sort `json:"sort,omitempty"`
}

type Sort struct {
	// ascending, descending
	Order int
	// value to order by
	Value string
}

func NewQuery() *Query {
	return &Query{}
}

func ParseQuery(ctx *gin.Context) (*Query, error) {
	logger.Debugf("Start parsing Query")
	query := NewQuery()

	logger.Debugf("Query parsed")
	return query, nil
}
