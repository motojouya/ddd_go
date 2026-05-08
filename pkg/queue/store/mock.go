package store

import (
	database "github.com/motojouya/ddd_go/pkg/database/core"
	queueCore "github.com/motojouya/ddd_go/pkg/queue/core"
)

type QueueMock struct {
	database.ORPer
	SelectWorkerByQueueFunc func(queueName string, forUpdate bool) (*queueCore.Worker, error)
}

func NewQueueMock(executorMock database.ORPer) *QueueMock {
	return &QueueMock{
		ORPer: executorMock,
		SelectWorkerByQueueFunc: func(queueName string, forUpdate bool) (*queueCore.Worker, error) {
			return nil, nil
		},
	}
}

func (mock *QueueMock) SelectWorkerByQueue(queueName string, forUpdate bool) (*queueCore.Worker, error) {
	return mock.SelectWorkerByQueueFunc(queueName, forUpdate)
}
