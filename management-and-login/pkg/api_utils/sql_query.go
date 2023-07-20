package api_utils

import (
	"fmt"
	bd "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/mysql"
	"regexp"
	"strings"
)

func (q *Query) BuildSQLFromApiQuery() (*bd.Query, error) {
	sql := bd.Query{
		Limit:  q.Limit,
		Offset: q.Offset,
	}
	if q.Sort != nil {
		sql.Order = camelToUpperSnakeCase(q.Sort.Value) + space + q.Sort.Order
	}

	return conditionBuilder(q.TableName, &sql, q.FilterFields, q.FilterValues), nil
}

func (q *Query) BuildSQLFromApiCheck() (*bd.Query, error) {
	sql := bd.Query{
		Limit:  1,
		Offset: 0,
	}

	return checkBuilder(q.TableName, &sql, q.FilterFields, q.FilterValues), nil
}

// ff - filter fields (name, lastName, pesel)
// fv - filter values (Jan, Kowalski)
// result: (`name` = 'Jan' or `lastName` = 'Jan' or `pesel` = 'Jan') and (`name` = 'Kowalski' or `lastName` = 'Kowalski' or `pesel` = 'Kowalski')
func conditionBuilder(tableName string, s *bd.Query, ff []string, fv []string) *bd.Query {
	condition := &strings.Builder{}
	for i, value := range fv {
		condition.WriteString(leftBracket)
		for j, field := range ff {
			condition.WriteString(fmt.Sprintf("`%s`.`%s` %s '%s%s%s'",
				tableName,
				camelToUpperSnakeCase(field),
				like,
				wildcard,
				value,
				wildcard))
			if j != len(ff)-1 {
				condition.WriteString(or)
			}
		}
		condition.WriteString(rightBracket)
		if i != len(fv)-1 {
			condition.WriteString(or)
		}
	}
	s.Filter = condition.String()

	return s
}

// ff - filter fields (pesel)
// fv - filter values (123456789)
// result: (`PESEL` = '123456789')
func checkBuilder(tableName string, s *bd.Query, ff []string, fv []string) *bd.Query {
	condition := &strings.Builder{}
	for i, value := range fv {
		condition.WriteString(leftBracket)
		for j, field := range ff {
			condition.WriteString(fmt.Sprintf("`%s`.`%s` %s '%s'",
				tableName,
				camelToUpperSnakeCase(field),
				equal,
				value))
			if j != len(ff)-1 {
				condition.WriteString(or)
			}
		}
		condition.WriteString(rightBracket)
		if i != len(fv)-1 {
			condition.WriteString(or)
		}
	}
	s.Filter = condition.String()

	return s
}

func camelToUpperSnakeCase(c string) string {
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(c, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToUpper(snake)
}
