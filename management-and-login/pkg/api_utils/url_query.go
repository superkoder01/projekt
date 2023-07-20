package api_utils

import (
	"net/url"
	"strconv"
	"strings"
)

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
	TableName string `json:"tableName,omitempty"`
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
	Order string
	// value to order by
	Value string
}

type FilterValue struct {
	// lte, gte, eq
	Operator string
	// value to filter by
	Value string
}

func NewQuery(tableName string) *Query {
	return &Query{TableName: tableName}
}

func (q *Query) init() {
	q.Limit = 10
	q.Offset = 0
}

func ParseQuery(tableName string, urlQuery string) *Query {
	query := NewQuery(tableName)
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
		s.Order = ASC
		s.Value = val[4:]
	} else if val[:4] == DESC {
		s.Order = DESC
		s.Value = val[5:]
	} else {
		return nil
	}

	return &s
}

//func parseFilter(values []string) []*FilterValue {
//	var fvs []*FilterValue
//	for _, val := range values {
//		if len(val) < 4 || !strings.Contains(val, ":") {
//			return nil
//		}
//
//		fv := FilterValue{}
//
//		if val[:2] == EQ {
//			fv.Operator = EQ
//			fv.Value = val[3:]
//		} else if val[:3] == LTE {
//			fv.Operator = LTE
//			fv.Value = val[4:]
//		} else if val[:3] == GTE {
//			fv.Operator = GTE
//			fv.Value = val[4:]
//		}
//
//		fvs = append(fvs, &fv)
//	}
//
//	return fvs
//}
