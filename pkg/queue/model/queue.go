package model

import (
	"strconv"

	"github.com/go-gorp/gorp/v3"
	basic "github.com/motojouya/ddd_go/pkg/basic/model"
)

const QueueTable = "queue"
const QueueAlias = "q"

type Queue struct {
	Name         string `db:"name"`
	WorkerName   string `db:"worker_name"`
	ProcessOrder int    `db:"process_order"`
}

func (q Queue) Keys() []interface{} {
	return []interface{}{q.Name}
}

func AddQueueTable(dbMap *gorp.DbMap) {
	dbMap.AddTableWithName(Queue{}, QueueTable).SetKeys(false, "Name")
}

func NewQueue(name string, processOrder int, worker Worker) (Queue, error) {
	if name == "" {
		return Queue{}, basic.NewInvalidArgumentError("name", name, "name is required")
	}
	if worker.Name == "" {
		return Queue{}, basic.NewInvalidArgumentError("workerName", worker.Name, "workerName is required")
	}
	if processOrder <= 0 {
		return Queue{}, basic.NewInvalidArgumentError("processOrder", strconv.Itoa(processOrder), "processOrder should be over 0")
	}
	return Queue{
		Name:         name,
		WorkerName:   worker.Name,
		ProcessOrder: processOrder,
	}, nil
}
