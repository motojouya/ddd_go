package repository

import (
	database "github.com/motojouya/ddd_go/pkg/database/repository"
	queueModel "github.com/motojouya/ddd_go/pkg/queue/model"
	queueStore "github.com/motojouya/ddd_go/pkg/queue/store"
)

func CreateWorker(store queueStore.QueueStore, workerName string, maxProcess int) (queueModel.Worker, error) {
	worker, err := queueModel.NewWorker(workerName, maxProcess)
	if err != nil {
		return queueModel.Worker{}, err
	}

	return database.GetOrCreate(store, worker)
}
