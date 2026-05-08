package repository

import (
	local "github.com/motojouya/ddd_go/pkg/local/repository"
	queueCore "github.com/motojouya/ddd_go/pkg/queue/core"
	queueStore "github.com/motojouya/ddd_go/pkg/queue/store"
)

type QueueRepository interface {
	// Mutation
	CreateWorker(workerName string, maxProcess int) (queueCore.Worker, error)
	CreateQueue(workerName string, queueName string, processOrder int) (queueCore.Queue, error)
	RegisterJob(queueName string, source string, procedure string, jsonData any) (queueCore.Job, error)
	StartJob(job queueCore.Job) (queueCore.Job, error)
	FinishJob(job queueCore.Job, result string, errCause error) (queueCore.Job, error)
	// Query
	GetQueue(name string) (*queueCore.Queue, error)
	GetWorker(name string) (*queueCore.Worker, error)
	GetWorkerByQueue(queueName string) (*queueCore.Worker, error)
	GetQueueByWorker(workerName string) ([]queueCore.Queue, error)
	GetJobByQueue(queueName string, limit int) ([]queueCore.Job, error)
}

type queueBehave struct {
	store   queueStore.QueueStore
	localer local.Localer
}

func NewQueueRepository(store queueStore.QueueStore, localer local.Localer) QueueRepository {
	return &queueBehave{store: store, localer: localer}
}

func (b *queueBehave) CreateWorker(workerName string, maxProcess int) (queueCore.Worker, error) {
	return CreateWorker(b.store, workerName, maxProcess)
}

func (b *queueBehave) CreateQueue(workerName string, queueName string, processOrder int) (queueCore.Queue, error) {
	return CreateQueue(b.store, workerName, queueName, processOrder)
}

func (b *queueBehave) RegisterJob(queueName string, source string, procedure string, jsonData any) (queueCore.Job, error) {
	return RegisterJob(b.localer, b.store, queueName, source, procedure, jsonData)
}

func (b *queueBehave) StartJob(job queueCore.Job) (queueCore.Job, error) {
	return StartJob(b.localer, b.store, job)
}

func (b *queueBehave) FinishJob(job queueCore.Job, result string, errCause error) (queueCore.Job, error) {
	return FinishJob(b.localer, b.store, job, result, errCause)
}

func (b *queueBehave) GetQueue(name string) (*queueCore.Queue, error) {
	return GetQueue(b.store, name, false)
}

func (b *queueBehave) GetWorker(name string) (*queueCore.Worker, error) {
	return GetWorker(b.store, name, false)
}

func (b *queueBehave) GetWorkerByQueue(queueName string) (*queueCore.Worker, error) {
	return b.store.SelectWorkerByQueue(queueName, false)
}

func (b *queueBehave) GetQueueByWorker(workerName string) ([]queueCore.Queue, error) {
	return GetQueueByWorker(b.store, workerName, false)
}

func (b *queueBehave) GetJobByQueue(queueName string, limit int) ([]queueCore.Job, error) {
	return GetJobByQueue(b.store, queueName, limit, false)
}
