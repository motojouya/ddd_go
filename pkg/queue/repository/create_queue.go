package repository

import (
	"errors"

	database "github.com/motojouya/ddd_go/pkg/database/repository"
	queueModel "github.com/motojouya/ddd_go/pkg/queue/model"
	queueStore "github.com/motojouya/ddd_go/pkg/queue/store"
)

func CreateQueue(store queueStore.QueueStore, workerName string, queueName string, processOrder int) (queueModel.Queue, error) {
	worker, err := GetWorker(store, workerName, false)
	if err != nil {
		return queueModel.Queue{}, err
	}

	if worker == nil {
		return queueModel.Queue{}, errors.New("worker does not exists. worker_name: " + workerName)
	}

	queue, err := queueModel.NewQueue(queueName, processOrder, *worker)
	if err != nil {
		return queueModel.Queue{}, err
	}

	return database.GetOrCreate(store, queue)
}
