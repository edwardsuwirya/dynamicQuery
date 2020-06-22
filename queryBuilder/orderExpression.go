package queryBuilder

import "fmt"

type IOrderExpression interface {
	OrderResult() string
}

type ExpOrderBy struct {
	FieldName string
	Direction string
}

func (o ExpOrderBy) OrderResult() string {
	switch o.Direction {
	case "ASC", "DESC":
		return fmt.Sprintf("%v %v", o.FieldName, o.Direction)
	default:
		return fmt.Sprintf("%v %v", o.FieldName, "ASC")
	}
}
