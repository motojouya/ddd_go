package repository

import (
	local "github.com/motojouya/ddd_go/pkg/local/repository"
	queueModel "github.com/motojouya/ddd_go/pkg/queue/model"
	queueStore "github.com/motojouya/ddd_go/pkg/queue/store"
)

type QueueRepository interface {
	// Mutation
	CreateWorker(workerName string, maxProcess int) (queueModel.Worker, error)
	CreateQueue(workerName string, queueName string, processOrder int) (queueModel.Queue, error)
	RegisterJob(queueName string, source string, procedure string, jsonData any) (queueModel.Job, error)
	StartJob(job queueModel.Job) (queueModel.Job, error)
	FinishJob(job queueModel.Job, result string, errCause error) (queueModel.Job, error)
	// Query
	GetQueue(name string) (*queueModel.Queue, error)
	GetWorker(name string) (*queueModel.Worker, error)
	GetWorkerByQueue(queueName string) (*queueModel.Worker, error)
	GetQueueByWorker(workerName string) ([]queueModel.Queue, error)
	GetJobByQueue(queueName string, limit int) ([]queueModel.Job, error)
}

type queueBehave struct {
	store   queueStore.QueueStore
	localer local.Localer
}

func NewQueueRepository(store queueStore.QueueStore, localer local.Localer) QueueRepository {
	return &queueBehave{store: store, localer: localer}
}

func (b *queueBehave) CreateWorker(workerName string, maxProcess int) (queueModel.Worker, error) {
	return CreateWorker(b.store, workerName, maxProcess)
}

func (b *queueBehave) CreateQueue(workerName string, queueName string, processOrder int) (queueModel.Queue, error) {
	return CreateQueue(b.store, workerName, queueName, processOrder)
}

func (b *queueBehave) RegisterJob(queueName string, source string, procedure string, jsonData any) (queueModel.Job, error) {
	return RegisterJob(b.localer, b.store, queueName, source, procedure, jsonData)
}

func (b *queueBehave) StartJob(job queueModel.Job) (queueModel.Job, error) {
	return StartJob(b.localer, b.store, job)
}

func (b *queueBehave) FinishJob(job queueModel.Job, result string, errCause error) (queueModel.Job, error) {
	return FinishJob(b.localer, b.store, job, result, errCause)
}

func (b *queueBehave) GetQueue(name string) (*queueModel.Queue, error) {
	return GetQueue(b.store, name, false)
}

func (b *queueBehave) GetWorker(name string) (*queueModel.Worker, error) {
	return GetWorker(b.store, name, false)
}

func (b *queueBehave) GetWorkerByQueue(queueName string) (*queueModel.Worker, error) {
	return b.store.SelectWorkerByQueue(queueName, false)
}

func (b *queueBehave) GetQueueByWorker(workerName string) ([]queueModel.Queue, error) {
	return GetQueueByWorker(b.store, workerName, false)
}

func (b *queueBehave) GetJobByQueue(queueName string, limit int) ([]queueModel.Job, error) {
	return GetJobByQueue(b.store, queueName, limit, false)
}
