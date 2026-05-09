package store

import (
	database "github.com/motojouya/ddd_go/pkg/database/model"
	queueModel "github.com/motojouya/ddd_go/pkg/queue/model"
)

type QueueMock struct {
	database.ORPer
	SelectWorkerByQueueFunc func(queueName string, forUpdate bool) (*queueModel.Worker, error)
}

func NewQueueMock(executorMock database.ORPer) *QueueMock {
	return &QueueMock{
		ORPer: executorMock,
		SelectWorkerByQueueFunc: func(queueName string, forUpdate bool) (*queueModel.Worker, error) {
			return nil, nil
		},
	}
}

func (mock *QueueMock) SelectWorkerByQueue(queueName string, forUpdate bool) (*queueModel.Worker, error) {
	return mock.SelectWorkerByQueueFunc(queueName, forUpdate)
}
