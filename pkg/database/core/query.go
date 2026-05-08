package core

import (
	"github.com/go-gorp/gorp/v3"
	basic "github.com/motojouya/ddd_go/pkg/basic/core"
)

type Executor interface {
	gorp.SqlExecutor
	GetIn(records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error)
	GetMax(records interface{}, maxColName string, conditions map[string]interface{}) (int, error)
	GetPaging(records interface{}, conditions map[string]interface{}, orders []Order, pager basic.Pager) ([]interface{}, error)
}
