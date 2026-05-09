package repository

import (
	"errors"
	"sync"

	"github.com/go-gorp/gorp/v3"
	company "github.com/motojouya/ddd_go/pkg/company/model"
	"github.com/motojouya/ddd_go/pkg/database/model"
	localRepository "github.com/motojouya/ddd_go/pkg/local/repository"
	queue "github.com/motojouya/ddd_go/pkg/queue/model"
)

type DatabaseGetter interface {
	GetDatabase() (model.ORPer, error)
}

type DatabaseGet struct{}

func NewDatabaseGet() *DatabaseGet {
	return &DatabaseGet{}
}

var orp model.ORPer
var once sync.Once

func (getter DatabaseGet) GetDatabase() (model.ORPer, error) {
	var err error
	once.Do(func() {
		dbAccess, err := localRepository.GetEnv[model.DBAccess]()
		if err != nil {
			return
		}

		connection, err := dbAccess.CreateConnection()
		if err != nil {
			return
		}

		orp = model.CreateDatabase(connection, RegisterTable)
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
