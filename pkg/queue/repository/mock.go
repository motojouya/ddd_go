package repository

import (
	queueModel "github.com/motojouya/ddd_go/pkg/queue/model"
)

type QueueMock struct {
	CreateWorkerFunc     func(workerName string, maxProcess int) (queueModel.Worker, error)
	CreateQueueFunc      func(workerName string, queueName string, processOrder int) (queueModel.Queue, error)
	RegisterJobFunc      func(queueName string, source string, procedure string, jsonData any) (queueModel.Job, error)
	StartJobFunc         func(job queueModel.Job) (queueModel.Job, error)
	FinishJobFunc        func(job queueModel.Job, result string, errCause error) (queueModel.Job, error)
	GetQueueFunc         func(name string) (*queueModel.Queue, error)
	GetWorkerFunc        func(name string) (*queueModel.Worker, error)
	GetWorkerByQueueFunc func(queueName string) (*queueModel.Worker, error)
	GetQueueByWorkerFunc func(workerName string) ([]queueModel.Queue, error)
	GetJobByQueueFunc    func(queueName string, limit int) ([]queueModel.Job, error)
}

func NewQueueMock() *QueueMock {
	return &QueueMock{
		CreateWorkerFunc: func(workerName string, maxProcess int) (queueModel.Worker, error) {
			return queueModel.Worker{}, nil
		},
		CreateQueueFunc: func(workerName string, queueName string, processOrder int) (queueModel.Queue, error) {
			return queueModel.Queue{}, nil
		},
		RegisterJobFunc: func(queueName string, source string, procedure string, jsonData any) (queueModel.Job, error) {
			return queueModel.Job{}, nil
		},
		StartJobFunc: func(job queueModel.Job) (queueModel.Job, error) {
			return queueModel.Job{}, nil
		},
		FinishJobFunc: func(job queueModel.Job, result string, errCause error) (queueModel.Job, error) {
			return queueModel.Job{}, nil
		},
		GetQueueFunc: func(name string) (*queueModel.Queue, error) {
			return nil, nil
		},
		GetWorkerFunc: func(name string) (*queueModel.Worker, error) {
			return nil, nil
		},
		GetWorkerByQueueFunc: func(queueName string) (*queueModel.Worker, error) {
			return nil, nil
		},
		GetQueueByWorkerFunc: func(workerName string) ([]queueModel.Queue, error) {
			return []queueModel.Queue{}, nil
		},
		GetJobByQueueFunc: func(queueName string, limit int) ([]queueModel.Job, error) {
			return []queueModel.Job{}, nil
		},
	}
}

func (mock *QueueMock) CreateWorker(workerName string, maxProcess int) (queueModel.Worker, error) {
	return mock.CreateWorkerFunc(workerName, maxProcess)
}

func (mock *QueueMock) CreateQueue(workerName string, queueName string, processOrder int) (queueModel.Queue, error) {
	return mock.CreateQueueFunc(workerName, queueName, processOrder)
}

func (mock *QueueMock) RegisterJob(queueName string, source string, procedure string, jsonData any) (queueModel.Job, error) {
	return mock.RegisterJobFunc(queueName, source, procedure, jsonData)
}

func (mock *QueueMock) StartJob(job queueModel.Job) (queueModel.Job, error) {
	return mock.StartJobFunc(job)
}

func (mock *QueueMock) FinishJob(job queueModel.Job, result string, errCause error) (queueModel.Job, error) {
	return mock.FinishJobFunc(job, result, errCause)
}

func (mock *QueueMock) GetQueue(name string) (*queueModel.Queue, error) {
	return mock.GetQueueFunc(name)
}

func (mock *QueueMock) GetWorker(name string) (*queueModel.Worker, error) {
	return mock.GetWorkerFunc(name)
}

func (mock *QueueMock) GetWorkerByQueue(queueName string) (*queueModel.Worker, error) {
	return mock.GetWorkerByQueueFunc(queueName)
}

func (mock *QueueMock) GetQueueByWorker(workerName string) ([]queueModel.Queue, error) {
	return mock.GetQueueByWorkerFunc(workerName)
}

func (mock *QueueMock) GetJobByQueue(queueName string, limit int) ([]queueModel.Job, error) {
	return mock.GetJobByQueueFunc(queueName, limit)
}
