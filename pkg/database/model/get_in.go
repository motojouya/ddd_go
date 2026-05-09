package model

import (
	"errors"

	"github.com/doug-martin/goqu/v9"
	"github.com/go-gorp/gorp/v3"
	basic "github.com/motojouya/ddd_go/pkg/basic/model"
)

func GetIn(executer gorp.SqlExecutor, records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error) {
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

		whereExpressions[counter] = goqu.C(key).In(condition)
		counter++
	}

	query := Dialect.
		Select(selectExpressions...).
		From(goqu.T(tableMap.TableName)).
		Where(whereExpressions...)

	if forLock {
		query = query.ForUpdate(goqu.Wait)
	}

	sql, args, err := query.Prepared(true).ToSQL()
	if err != nil {
		return nil, err
	}

	return executer.Select(records, sql, args...)
}

func (o *ORP) GetIn(records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error) {
	return GetIn(o, records, conditions, forLock)
}
