package main

import (
	"fmt"
	"github.com/edwardsuwirya/dynamicQuery/queryBuilder"
)

func main() {
	var dbq queryBuilder.DynamicQuery

	//result := dbq.Or(
	//	queryBuilder.Exp{"fname", "=", "'Edo'"},
	//	queryBuilder.Exp{"age", ">=", "17"},
	//	queryBuilder.ExpBetween{"order_date", "2019-01-01", "2019-02-03"},
	//	//queryBuilder.ExpIn{"order_date", "2019-01-01", "2019-02-03"},
	//	dbq.And(
	//		queryBuilder.Exp{"IDCARD", "=", "XXX"},
	//		queryBuilder.Exp{"SIM", "=", "YYY"},
	//	),
	//).OrderBy(
	//	queryBuilder.ExpOrderBy{"IDCARD", "ASC"},
	//	queryBuilder.ExpOrderBy{"fname", ""}).Select("*", "m_customer")
	result := dbq.Where(
		queryBuilder.ExpIn{"order_date", nil, 3},
	).OrderBy(
		queryBuilder.ExpOrderBy{"IDCARD", "ASC"},
		queryBuilder.ExpOrderBy{"fname", ""}).Select("*", "m_customer")

	fmt.Printf("%v\n", result)
}
