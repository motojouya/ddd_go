package behavior

import (
	database "github.com/motojouya/ddd_go/pkg/database/behavior"
	queueCore "github.com/motojouya/ddd_go/pkg/queue/core"
	queueStore "github.com/motojouya/ddd_go/pkg/queue/store"
)

func CreateWorker(store queueStore.QueueStore, workerName string, maxProcess int) (queueCore.Worker, error) {
	worker, err := queueCore.NewWorker(workerName, maxProcess)
	if err != nil {
		return queueCore.Worker{}, err
	}

	return database.GetOrCreate(store, worker)
}
