package repository

import (
	basic "github.com/motojouya/ddd_go/pkg/basic/core"
	database "github.com/motojouya/ddd_go/pkg/database/core"
	queueCore "github.com/motojouya/ddd_go/pkg/queue/core"
)

func GetQueue(db database.Executor, name string, forUpdate bool) (*queueCore.Queue, error) {
	conditions := map[string][]interface{}{
		"name": {name},
	}

	var result []queueCore.Queue
	_, err := db.GetIn(&result, conditions, forUpdate)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, nil
	}

	return &result[0], nil
}

func GetWorker(db database.Executor, name string, forUpdate bool) (*queueCore.Worker, error) {
	conditions := map[string][]interface{}{
		"name": {name},
	}

	var result []queueCore.Worker
	_, err := db.GetIn(&result, conditions, forUpdate)
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, nil
	}

	return &result[0], nil
}

func GetQueueByWorker(db database.Executor, workerName string, forUpdate bool) ([]queueCore.Queue, error) {
	conditions := map[string]interface{}{
		"worker_name": workerName,
	}

	orders := []database.Order{
		{Column: "process_order", Ascending: true},
	}

	pager := basic.Pager{Cursor: 1, Limit: 1000}

	var result []queueCore.Queue
	_, err := db.GetPaging(&result, conditions, orders, pager)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetJobByQueue(db database.Executor, queueName string, limit int, forUpdate bool) ([]queueCore.Job, error) {
	conditions := map[string]interface{}{
		"queue":      queueName,
		"start_date": nil,
	}

	orders := []database.Order{
		{Column: "register_date", Ascending: true},
	}

	pager := basic.Pager{Cursor: 1, Limit: uint(limit)}

	var result []queueCore.Job
	_, err := db.GetPaging(&result, conditions, orders, pager)
	if err != nil {
		return nil, err
	}

	return result, nil
}
