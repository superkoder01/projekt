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
	"net/url"
	"strconv"
	"strings"
)

type QueryResult struct {
	Amount   int64       `json:"amount,omitempty"`
	Elements interface{} `json:"elements,omitempty"`
}

func (c *QueryResult) String() string {
	return fmt.Sprintf("%s", *c)
}

const (
	// query params
	limit  = "limit"
	offset = "offset"
	page   = "page"
	sort   = "sort"

	// order by
	ASC  = "asc"
	DESC = "desc"

	//// filter operator
	//LTE = "lte"
	//GTE = "gte"
	//EQ  = "eq"

	// filter
	filterFields = "filterFields"
	filterValues = "filterValues"
)

type Query struct {
	//Filter map[string][]*FilterValue `json:"filter,omitempty"`
	FilterFields []string `json:"filterFields,omitempty"`
	FilterValues []string `json:"filterValues,omitempty"`
	Limit        int      `json:"limit,omitempty"`
	Offset       int      `json:"offset,omitempty"`
	Page         int      `json:"page,omitempty"`
	Sort         *Sort    `json:"sort,omitempty"`
}

type Sort struct {
	// ascending, descending
	Order int
	// value to order by
	Value string
}

type FilterValue struct {
	// lte, gte, eq
	Operator string
	// value to filter by
	Value string
}

func NewQuery() *Query {
	return &Query{}
}

func (q *Query) init() {
	q.Limit = 10
	q.Offset = 0
}

func ParseQuery(urlQuery string) *Query {
	query := NewQuery()
	query.init()

	urlString, err := url.Parse(urlQuery)
	if err != nil {
		return query
	}

	for k, v := range urlString.Query() {
		switch k {
		case limit:
			if ps, err := strconv.Atoi(v[0]); err == nil && ps >= 1 {
				query.Limit = ps
			}
		case offset:
			if of, err := strconv.Atoi(v[0]); err == nil && of >= 0 {
				query.Offset = of
			}
		case page:
			if pn, err := strconv.Atoi(v[0]); err == nil && pn >= 1 {
				query.Page = pn
				query.Offset = query.Limit * (query.Page - 1)
			}
		case sort:
			query.Sort = parseSort(v[0])
		case filterFields:
			query.FilterFields = strings.Split(v[0], ",")
		case filterValues:
			query.FilterValues = strings.Split(v[0], ",")
		}
	}

	return query
}

func parseSort(val string) *Sort {
	if len(val) < 5 {
		return nil
	}

	s := Sort{}
	if val[:3] == ASC {
		s.Order = 1
		s.Value = val[4:]
	} else if val[:4] == DESC {
		s.Order = -1
		s.Value = val[5:]
	} else {
		return nil
	}

	return &s
}
