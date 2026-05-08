package behavior

import (
	queueCore "github.com/motojouya/ddd_go/pkg/queue/core"
)

type QueueMock struct {
	CreateWorkerFunc     func(workerName string, maxProcess int) (queueCore.Worker, error)
	CreateQueueFunc      func(workerName string, queueName string, processOrder int) (queueCore.Queue, error)
	RegisterJobFunc      func(queueName string, source string, procedure string, jsonData any) (queueCore.Job, error)
	StartJobFunc         func(job queueCore.Job) (queueCore.Job, error)
	FinishJobFunc        func(job queueCore.Job, result string, errCause error) (queueCore.Job, error)
	GetQueueFunc         func(name string) (*queueCore.Queue, error)
	GetWorkerFunc        func(name string) (*queueCore.Worker, error)
	GetWorkerByQueueFunc func(queueName string) (*queueCore.Worker, error)
	GetQueueByWorkerFunc func(workerName string) ([]queueCore.Queue, error)
	GetJobByQueueFunc    func(queueName string, limit int) ([]queueCore.Job, error)
}

func NewQueueMock() *QueueMock {
	return &QueueMock{
		CreateWorkerFunc: func(workerName string, maxProcess int) (queueCore.Worker, error) {
			return queueCore.Worker{}, nil
		},
		CreateQueueFunc: func(workerName string, queueName string, processOrder int) (queueCore.Queue, error) {
			return queueCore.Queue{}, nil
		},
		RegisterJobFunc: func(queueName string, source string, procedure string, jsonData any) (queueCore.Job, error) {
			return queueCore.Job{}, nil
		},
		StartJobFunc: func(job queueCore.Job) (queueCore.Job, error) {
			return queueCore.Job{}, nil
		},
		FinishJobFunc: func(job queueCore.Job, result string, errCause error) (queueCore.Job, error) {
			return queueCore.Job{}, nil
		},
		GetQueueFunc: func(name string) (*queueCore.Queue, error) {
			return nil, nil
		},
		GetWorkerFunc: func(name string) (*queueCore.Worker, error) {
			return nil, nil
		},
		GetWorkerByQueueFunc: func(queueName string) (*queueCore.Worker, error) {
			return nil, nil
		},
		GetQueueByWorkerFunc: func(workerName string) ([]queueCore.Queue, error) {
			return []queueCore.Queue{}, nil
		},
		GetJobByQueueFunc: func(queueName string, limit int) ([]queueCore.Job, error) {
			return []queueCore.Job{}, nil
		},
	}
}

func (mock *QueueMock) CreateWorker(workerName string, maxProcess int) (queueCore.Worker, error) {
	return mock.CreateWorkerFunc(workerName, maxProcess)
}

func (mock *QueueMock) CreateQueue(workerName string, queueName string, processOrder int) (queueCore.Queue, error) {
	return mock.CreateQueueFunc(workerName, queueName, processOrder)
}

func (mock *QueueMock) RegisterJob(queueName string, source string, procedure string, jsonData any) (queueCore.Job, error) {
	return mock.RegisterJobFunc(queueName, source, procedure, jsonData)
}

func (mock *QueueMock) StartJob(job queueCore.Job) (queueCore.Job, error) {
	return mock.StartJobFunc(job)
}

func (mock *QueueMock) FinishJob(job queueCore.Job, result string, errCause error) (queueCore.Job, error) {
	return mock.FinishJobFunc(job, result, errCause)
}

func (mock *QueueMock) GetQueue(name string) (*queueCore.Queue, error) {
	return mock.GetQueueFunc(name)
}

func (mock *QueueMock) GetWorker(name string) (*queueCore.Worker, error) {
	return mock.GetWorkerFunc(name)
}

func (mock *QueueMock) GetWorkerByQueue(queueName string) (*queueCore.Worker, error) {
	return mock.GetWorkerByQueueFunc(queueName)
}

func (mock *QueueMock) GetQueueByWorker(workerName string) ([]queueCore.Queue, error) {
	return mock.GetQueueByWorkerFunc(workerName)
}

func (mock *QueueMock) GetJobByQueue(queueName string, limit int) ([]queueCore.Job, error) {
	return mock.GetJobByQueueFunc(queueName, limit)
}
