package store

import (
	database "github.com/motojouya/ddd_go/pkg/database/core"
	queueCore "github.com/motojouya/ddd_go/pkg/queue/core"
)

type QueueStore interface {
	database.ORPer
	SelectWorkerByQueue(queueName string, forUpdate bool) (*queueCore.Worker, error)
}

type queueStore struct {
	database.ORPer
}

func NewQueueStore(db database.ORPer) QueueStore {
	return &queueStore{
		ORPer: db,
	}
}
