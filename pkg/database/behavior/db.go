package behavior

import (
	"errors"
	"sync"

	"github.com/go-gorp/gorp/v3"
	company "github.com/motojouya/ddd_go/pkg/company/core"
	"github.com/motojouya/ddd_go/pkg/database/core"
	localBehavior "github.com/motojouya/ddd_go/pkg/local/behavior"
	queue "github.com/motojouya/ddd_go/pkg/queue/core"
)

type DatabaseGetter interface {
	GetDatabase() (core.ORPer, error)
}

type DatabaseGet struct{}

func NewDatabaseGet() *DatabaseGet {
	return &DatabaseGet{}
}

var orp core.ORPer
var once sync.Once

func (getter DatabaseGet) GetDatabase() (core.ORPer, error) {
	var err error
	once.Do(func() {
		dbAccess, err := localBehavior.GetEnv[core.DBAccess]()
		if err != nil {
			return
		}

		connection, err := dbAccess.CreateConnection()
		if err != nil {
			return
		}

		orp = core.CreateDatabase(connection, RegisterTable)
		if err != nil {
			return
		}
	})

	if err != nil {
		return nil, err
	}

	if orp == nil {
		return nil, errors.New("Database Connection Disabled.")
	}

	return orp, nil
}

/**
 * ここにテーブル登録を追加していく
 */
func RegisterTable(dbMap *gorp.DbMap) {
	company.AddCompanyTable(dbMap)
	queue.AddWorkerTable(dbMap)
	queue.AddQueueTable(dbMap)
	queue.AddJobTable(dbMap)
}
