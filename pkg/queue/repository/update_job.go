package repository

import (
	"encoding/json"
	"errors"

	basic "github.com/motojouya/ddd_go/pkg/basic/model"
	local "github.com/motojouya/ddd_go/pkg/local/repository"
	queueModel "github.com/motojouya/ddd_go/pkg/queue/model"
	queueStore "github.com/motojouya/ddd_go/pkg/queue/store"
)

type UpdateCallback = func(queueModel.Job) queueModel.Job

func UpdateJob(
	localer local.Localer,
	store queueStore.QueueStore,
	jobId basic.Identifier,
	update UpdateCallback,
) (queueModel.Job, error) {
	err := store.Begin()
	if err != nil {
		return queueModel.Job{}, err
	}

	conditions := map[string][]interface{}{
		"id": {jobId},
	}

	var jobs []queueModel.Job
	_, err = store.GetIn(&jobs, conditions, true)
	if err != nil {
		store.Rollback()
		return queueModel.Job{}, err
	}

	if len(jobs) == 0 {
		return queueModel.Job{}, errors.New("Job Is Gone.")
	}

	job := update(jobs[0])

	_, err = store.Update(&job)
	if err != nil {
		store.Rollback()
		return queueModel.Job{}, err
	}

	err = store.Commit()
	if err != nil {
		return queueModel.Job{}, err
	}

	return job, nil
}

func FinishJob(
	localer local.Localer,
	store queueStore.QueueStore,
	job queueModel.Job,
	result string,
	errCause error,
) (queueModel.Job, error) {

	errSerial := ""
	status := true
	if errCause != nil {
		errJson := queueModel.ErrorJson{Err: errCause.Error()}
		serial, err := json.Marshal(errJson)
		if err != nil {
			return queueModel.Job{}, err
		}
		errSerial = string(serial)
		status = false
	}

	return UpdateJob(localer, store, job.Id, func(job queueModel.Job) queueModel.Job {
		return queueModel.FinishJob(job, result, errSerial, localer.GetNow(), status)
	})
}

func StartJob(
	localer local.Localer,
	store queueStore.QueueStore,
	job queueModel.Job,
) (queueModel.Job, error) {
	return UpdateJob(localer, store, job.Id, func(job queueModel.Job) queueModel.Job {
		return queueModel.StartJob(job, localer.GetNow())
	})
}
