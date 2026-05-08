package core

import (
	"strconv"

	"github.com/go-gorp/gorp/v3"
	basic "github.com/motojouya/ddd_go/pkg/basic/core"
)

const WorkerTable = "worker"
const WorkerAlias = "w"

type Worker struct {
	Name       string `db:"name"`
	MaxProcess int    `db:"max_process"`
}

func (w Worker) Keys() []interface{} {
	return []interface{}{w.Name}
}

func AddWorkerTable(dbMap *gorp.DbMap) {
	dbMap.AddTableWithName(Worker{}, WorkerTable).SetKeys(false, "Name")
}

func NewWorker(name string, maxPrcess int) (Worker, error) {
	if name == "" {
		return Worker{}, basic.NewInvalidArgumentError("name", name, "name is required")
	}
	if maxPrcess <= 0 || maxPrcess > 1000 {
		return Worker{}, basic.NewInvalidArgumentError("maxPrcess", strconv.Itoa(maxPrcess), "maxPrcess should be between 1 to 1000")
	}
	return Worker{
		Name:       name,
		MaxProcess: maxPrcess,
	}, nil
}
