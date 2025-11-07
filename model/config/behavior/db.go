package behavior

import (
	"github.com/motojouya/mvc_go/valve/database"
	"github.com/motojouya/mvc_go/valve/local"
	"github.com/motojouya/mvc_go/model/config/core"
)

type DatabaseGetter interface {
	GetDatabase() (*database.ORP, error)
}

type DatabaseGet struct {}

func NewDatabaseGet() *DatabaseGet {
	return &DatabaseGet{}
}

var dbAccess *core.DBAccess

func (getter DatabaseGet) GetDatabase() (*database.ORP, error) {
	// access 情報はcacheするが、connectionはcacheしない
	if dbAccess == nil {
		var dbAccessData, err = local.GetEnv[*core.DBAccess]()
		if err != nil {
			return nil, err
		}
		dbAccess = &dbAccessData
	}

	var connection, err = dbAccess.CreateConnection()
	if err != nil {
		return nil, err
	}

	var db = database.CreateDatabase(connection)
	if err != nil {
		return db, err
	}

	return db, nil
}
