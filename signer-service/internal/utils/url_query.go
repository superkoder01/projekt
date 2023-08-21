/*
BMSFES. 
Copyright (C) 2022-2035 

This file is part of the BMSFES solution. 
BMSFES is free software: you can redistribute it and/or modify 
it under the terms of the GNU Affero General Public License as 
published by the Free Software Foundation, either version 3 of the 
License, or (at your option) any later version.
 
BMSFES solution is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of 
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the 
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License 
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
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
