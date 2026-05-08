package repository

import (
	"encoding/json"
	"errors"

	basic "github.com/motojouya/ddd_go/pkg/basic/core"
	local "github.com/motojouya/ddd_go/pkg/local/repository"
	queueCore "github.com/motojouya/ddd_go/pkg/queue/core"
	queueStore "github.com/motojouya/ddd_go/pkg/queue/store"
)

type UpdateCallback = func(queueCore.Job) queueCore.Job

func UpdateJob(
	localer local.Localer,
	store queueStore.QueueStore,
	jobId basic.Identifier,
	update UpdateCallback,
) (queueCore.Job, error) {
	err := store.Begin()
	if err != nil {
		return queueCore.Job{}, err
	}

	conditions := map[string][]interface{}{
		"id": {jobId},
	}

	var jobs []queueCore.Job
	_, err = store.GetIn(&jobs, conditions, true)
	if err != nil {
		store.Rollback()
		return queueCore.Job{}, err
	}

	if len(jobs) == 0 {
		return queueCore.Job{}, errors.New("Job Is Gone.")
	}

	job := update(jobs[0])

	_, err = store.Update(&job)
	if err != nil {
		store.Rollback()
		return queueCore.Job{}, err
	}

	err = store.Commit()
	if err != nil {
		return queueCore.Job{}, err
	}

	return job, nil
}

func FinishJob(
	localer local.Localer,
	store queueStore.QueueStore,
	job queueCore.Job,
	result string,
	errCause error,
) (queueCore.Job, error) {

	errSerial := ""
	status := true
	if errCause != nil {
		errJson := queueCore.ErrorJson{Err: errCause.Error()}
		serial, err := json.Marshal(errJson)
		if err != nil {
			return queueCore.Job{}, err
		}
		errSerial = string(serial)
		status = false
	}

	return UpdateJob(localer, store, job.Id, func(job queueCore.Job) queueCore.Job {
		return queueCore.FinishJob(job, result, errSerial, localer.GetNow(), status)
	})
}

func StartJob(
	localer local.Localer,
	store queueStore.QueueStore,
	job queueCore.Job,
) (queueCore.Job, error) {
	return UpdateJob(localer, store, job.Id, func(job queueCore.Job) queueCore.Job {
		return queueCore.StartJob(job, localer.GetNow())
	})
}
