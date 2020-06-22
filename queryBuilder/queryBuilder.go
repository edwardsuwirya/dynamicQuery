package queryBuilder

import (
	"strings"
)

type IDynamicQuery interface {
	Where(exps IWhereExpression) IDynamicQuery
	And(exps ...interface{}) IDynamicQuery
	Or(exps ...interface{}) IDynamicQuery
	Limit() IDynamicQuery
	OrderBy(orderExp ...IOrderExpression) IDynamicQuery
	Select(fieldNames string, tableName string) string
}
type DynamicQuery struct {
	whereResult string
	orderResult string
	sqlResult   string
}

func (d DynamicQuery) Where(exp IWhereExpression) IDynamicQuery {
	d.whereResult = exp.WhereResult()
	return d
}

func (d DynamicQuery) andOrGenerator(sqlOperator string, exps ...interface{}) string {
	sliceOfExpression := make([]string, 0)
	for _, exp := range exps {
		switch exp.(type) {
		case IWhereExpression:
			sliceOfExpression = append(sliceOfExpression, exp.(IWhereExpression).WhereResult())
		default:
			sliceOfExpression = append(sliceOfExpression, "("+exp.(DynamicQuery).whereResult+")")
		}

	}
	if len(exps) == 1 {
		return sliceOfExpression[0]
	} else {
		return strings.Join(sliceOfExpression, sqlOperator)
	}
}
func (d DynamicQuery) And(exps ...interface{}) IDynamicQuery {
	d.whereResult = d.andOrGenerator(" AND ", exps...)
	return d
}

func (d DynamicQuery) Or(exps ...interface{}) IDynamicQuery {
	d.whereResult = d.andOrGenerator(" OR ", exps...)
	return d
}

func (d DynamicQuery) Limit() IDynamicQuery {
	panic("implement me")
}

func (d DynamicQuery) OrderBy(orderExp ...IOrderExpression) IDynamicQuery {
	sliceOfOrder := make([]string, 0)
	for _, ob := range orderExp {
		sliceOfOrder = append(sliceOfOrder, ob.OrderResult())
	}
	d.orderResult = strings.Join(sliceOfOrder, ",")
	return d
}

func (d DynamicQuery) Select(fieldNames string, tableName string) string {
	tempResult := "SELECT " + fieldNames + " FROM " + tableName
	if len(d.whereResult) > 1 {
		tempResult = tempResult + " WHERE " + d.whereResult
	}

	if len(d.orderResult) > 1 {
		tempResult = tempResult + " ORDER BY " + d.orderResult
	}
	d.sqlResult = tempResult

	return d.sqlResult

}
