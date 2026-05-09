package model

import (
	"errors"

	"github.com/doug-martin/goqu/v9"
	"github.com/go-gorp/gorp/v3"
	basic "github.com/motojouya/ddd_go/pkg/basic/model"
)

func GetMax(executer gorp.SqlExecutor, records interface{}, maxColName string, conditions map[string]interface{}) (int, error) {
	tableMap, err := tableFor(executer, records)
	if err != nil {
		// TODO errorはもう少しわかりやすいのにする
		return 0, errors.New("record is not table")
	}

	existColName := basic.Some(tableMap.Columns, func(colName *gorp.ColumnMap) bool {
		return maxColName == colName.ColumnName
	})
	if !existColName {
		// TODO errorはもう少しわかりやすいのにする
		return 0, errors.New(maxColName + " is not exist column in table")
	}

	counter := 0
	whereExpressions := make([]goqu.Expression, len(conditions))
	for key, condition := range conditions {
		whereExpressions[counter] = goqu.C(key).Eq(condition)
		counter++
	}

	sql, args, err := Dialect.
		Select(goqu.COALESCE(goqu.MAX(maxColName), 0)).
		From(goqu.T(tableMap.TableName)).
		Where(whereExpressions...).
		Prepared(true).
		ToSQL()
	if err != nil {
		return 0, err
	}

	max, err := executer.SelectInt(sql, args...)
	if err != nil {
		return 0, err
	}

	return int(max), nil
}

func (o *ORP) GetMax(records interface{}, maxColName string, conditions map[string]interface{}) (int, error) {
	return GetMax(o, records, maxColName, conditions)
}
