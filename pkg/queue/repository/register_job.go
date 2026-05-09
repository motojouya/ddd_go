package repository

import (
	"errors"

	local "github.com/motojouya/ddd_go/pkg/local/repository"
	queueModel "github.com/motojouya/ddd_go/pkg/queue/model"
	queueStore "github.com/motojouya/ddd_go/pkg/queue/store"
)

func RegisterJob(
	localer local.Localer,
	store queueStore.QueueStore,
	queueName string,
	source string,
	procedure string,
	jsonData any,
) (queueModel.Job, error) {

	queue, err := GetQueue(store, queueName, false)
	if err != nil {
		return queueModel.Job{}, err
	}
	if queue == nil {
		return queueModel.Job{}, errors.New("queue not found. name: " + queueName)
	}

	id, err := localer.GenerateID()
	if err != nil {
		return queueModel.Job{}, err
	}

	job, err := queueModel.NewJob(id, *queue, source, procedure, jsonData, localer.GetNow())
	if err != nil {
		return queueModel.Job{}, err
	}

	err = store.Insert(&job)
	if err != nil {
		return queueModel.Job{}, err
	}

	return job, nil
}
