package core

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/go-gorp/gorp/v3"
)

func Select[R any](executer gorp.SqlExecutor, query *goqu.SelectDataset) ([]R, error) {

	sql, args, err := query.Prepared(true).ToSQL()
	if err != nil {
		return nil, err
	}

	var records []R
	_, err = executer.Select(&records, sql, args...)
	if err != nil {
		return nil, err
	}

	return records, nil
}
