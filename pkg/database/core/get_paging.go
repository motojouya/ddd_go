package core

import (
	"errors"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/go-gorp/gorp/v3"
	basic "github.com/motojouya/ddd_go/pkg/basic/core"
)

func GetPaging(executer gorp.SqlExecutor, records interface{}, conditions map[string]interface{}, orders []Order, pager basic.Pager) ([]interface{}, error) {
	tableMap, err := tableFor(executer, records)
	if err != nil {
		// TODO errorはもう少しわかりやすいのにする
		return nil, errors.New("record is not table")
	}

	selectExpressions := make([]interface{}, 0, len(tableMap.Columns))
	for _, col := range tableMap.Columns {
		if col.Transient {
			continue
		}
		selectExpressions = append(selectExpressions, goqu.C(col.ColumnName))
	}

	counter := 0
	whereExpressions := make([]goqu.Expression, len(conditions))
	for key, condition := range conditions {
		exitCol := basic.Some(tableMap.Columns, func(col *gorp.ColumnMap) bool {
			return col.ColumnName == key
		})
		if !exitCol {
			// TODO errorはもう少しわかりやすいのにする
			return nil, errors.New("record is not table")
		}

		if condition == nil {
			whereExpressions[counter] = goqu.C(key).IsNull()
		} else {
			whereExpressions[counter] = goqu.C(key).Eq(condition)
		}
		counter++
	}

	orderExpressions := make([]exp.OrderedExpression, len(orders))
	for i, order := range orders {
		exitCol := basic.Some(tableMap.Columns, func(col *gorp.ColumnMap) bool {
			return col.ColumnName == order.Column
		})
		if !exitCol {
			// TODO errorはもう少しわかりやすいのにする
			return nil, errors.New("record is not table")
		}

		if order.Ascending {
			orderExpressions[i] = goqu.I(order.Column).Asc()
		} else {
			orderExpressions[i] = goqu.I(order.Column).Desc()
		}
	}

	query := Dialect.
		Select(selectExpressions...).
		From(goqu.T(tableMap.TableName)).
		Where(whereExpressions...).
		Order(orderExpressions...).
		Offset(uint(pager.Cursor - 1)).
		Limit(uint(pager.Limit))

	sql, args, err := query.Prepared(true).ToSQL()
	if err != nil {
		return nil, err
	}

	return executer.Select(records, sql, args...)
}

func (o *ORP) GetPaging(records interface{}, conditions map[string]interface{}, orders []Order, pager basic.Pager) ([]interface{}, error) {
	return GetPaging(o, records, conditions, orders, pager)
}
