package queryBuilder

import (
	"fmt"
	"strings"
)

type IWhereExpression interface {
	WhereResult() string
}

type Exp struct {
	FieldName string
	Operator  string
	Value     string
}

func (e Exp) WhereResult() string {
	if len(e.Value) == 0 {
		return fmt.Sprintf("%s %s %s", e.FieldName, e.Operator, "?")
	} else {
		return fmt.Sprintf("%s %s %s", e.FieldName, e.Operator, e.Value)
	}

}

type ExpBetween struct {
	FieldName string
	Value1    string
	Value2    string
}

func (e ExpBetween) WhereResult() string {
	if len(e.Value1) == 0 || len(e.Value2) == 0 {
		return fmt.Sprintf("%s BETWEEN %s AND %s", e.FieldName, "?", "?")
	} else {
		return fmt.Sprintf("%s BETWEEN %s AND %s", e.FieldName, e.Value1, e.Value2)
	}

}

type ExpIn struct {
	FieldName string
	Values    []string
	Length    int
}

func (e ExpIn) WhereResult() string {
	if len(e.Values) == 0 && e.Length == 0 {
		return ""
	} else if len(e.Values) == 1 {
		return fmt.Sprintf("%s IN (%s)", e.FieldName, e.Values[0])
	} else if len(e.Values) > 1 {
		return fmt.Sprintf("%s IN (%s)", e.FieldName, strings.Join(e.Values, ","))
	} else {
		sql := `IN(?` + strings.Repeat(",?", e.Length-1) + `)`
		return fmt.Sprintf("%s %s", e.FieldName, sql)
	}
}
