package controller

import (
	"context"
	"errors"
	"time"

	qBehavior "github.com/motojouya/ddd_go/pkg/queue/behavior"
	qCore "github.com/motojouya/ddd_go/pkg/queue/core"
)

func ExecuteWorker(qBhv qBehavior.QueueBehavior, route qCore.JobRouter, ctx context.Context, workerName string, keepWorking bool) error {
	worker, err := qBhv.GetWorker(workerName)
	if err != nil {
		return err
	}

	queues, err := qBhv.GetQueueByWorker(worker.Name)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return nil

		default:
			var j *qCore.Job = nil
			for _, queue := range queues {
				jobs, err := qBhv.GetJobByQueue(queue.Name, 1)
				if err != nil {
					return err
				}
				if len(jobs) == 1 {
					j = &jobs[0]
					break
				}
			}
			if j == nil {
				if keepWorking {
					time.Sleep(time.Second)
					continue
				} else {
					return nil
				}
			}

			job, err := qBhv.StartJob(*j)
			if err != nil {
				return err
			}

			exec, exist := route.GetExecuter(job.Procedure)
			if !exist {
				return errors.New("No Procedure.")
			}

			result, jobErr := exec(ctx, job)
			job, err = qBhv.FinishJob(job, result, jobErr)
			if err != nil {
				return err
			}
		}
	}
}
