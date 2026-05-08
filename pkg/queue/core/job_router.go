package core

import (
	"context"
	"errors"
)

// FIXME 返り値がstringなのかbyte[]なのか要検討
type ExecuteJob = func(context.Context, Job) (string, error)

// LocationStock.Allocate use basicController.HandleJob
type JobRouter struct {
	jobMap map[string]ExecuteJob
}

func NewJobRouter() JobRouter {
	return JobRouter{jobMap: make(map[string]ExecuteJob)}
}

func (jr *JobRouter) GetExecuter(key string) (ExecuteJob, bool) {
	exe, exists := jr.jobMap[key]
	return exe, exists
}

func (jr *JobRouter) AddExecuter(key string, procedure ExecuteJob) error {
	if jr.jobMap == nil {
		jr.jobMap = make(map[string]ExecuteJob)
	}
	if _, exists := jr.jobMap[key]; exists {
		return errors.New("Duplicate Procedure Key")
	}

	jr.jobMap[key] = procedure
	return nil
}
