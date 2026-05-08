package behavior

import (
	"errors"

	local "github.com/motojouya/ddd_go/pkg/local/behavior"
	queueCore "github.com/motojouya/ddd_go/pkg/queue/core"
	queueStore "github.com/motojouya/ddd_go/pkg/queue/store"
)

func RegisterJob(
	localer local.Localer,
	store queueStore.QueueStore,
	queueName string,
	source string,
	procedure string,
	jsonData any,
) (queueCore.Job, error) {

	queue, err := GetQueue(store, queueName, false)
	if err != nil {
		return queueCore.Job{}, err
	}
	if queue == nil {
		return queueCore.Job{}, errors.New("queue not found. name: " + queueName)
	}

	id, err := localer.GenerateID()
	if err != nil {
		return queueCore.Job{}, err
	}

	job, err := queueCore.NewJob(id, *queue, source, procedure, jsonData, localer.GetNow())
	if err != nil {
		return queueCore.Job{}, err
	}

	err = store.Insert(&job)
	if err != nil {
		return queueCore.Job{}, err
	}

	return job, nil
}
