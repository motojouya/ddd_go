package repository

import (
	"errors"

	database "github.com/motojouya/ddd_go/pkg/database/repository"
	queueCore "github.com/motojouya/ddd_go/pkg/queue/core"
	queueStore "github.com/motojouya/ddd_go/pkg/queue/store"
)

func CreateQueue(store queueStore.QueueStore, workerName string, queueName string, processOrder int) (queueCore.Queue, error) {
	worker, err := GetWorker(store, workerName, false)
	if err != nil {
		return queueCore.Queue{}, err
	}

	if worker == nil {
		return queueCore.Queue{}, errors.New("worker does not exists. worker_name: " + workerName)
	}

	queue, err := queueCore.NewQueue(queueName, processOrder, *worker)
	if err != nil {
		return queueCore.Queue{}, err
	}

	return database.GetOrCreate(store, queue)
}
