package store

import (
	database "github.com/motojouya/ddd_go/pkg/database/model"
	queueModel "github.com/motojouya/ddd_go/pkg/queue/model"
)

type QueueStore interface {
	database.ORPer
	SelectWorkerByQueue(queueName string, forUpdate bool) (*queueModel.Worker, error)
}

type queueStore struct {
	database.ORPer
}

func NewQueueStore(db database.ORPer) QueueStore {
	return &queueStore{
		ORPer: db,
	}
}
