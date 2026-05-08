package store

import (
	"github.com/doug-martin/goqu/v9"
	database "github.com/motojouya/ddd_go/pkg/database/core"
	queueCore "github.com/motojouya/ddd_go/pkg/queue/core"
)

func SelectWorkerByQueue(db database.Executor, queueName string, forUpdate bool) (*queueCore.Worker, error) {
	query := database.Dialect.Select(
		goqu.I(queueCore.WorkerAlias+".name"),
		goqu.I(queueCore.WorkerAlias+".max_process"),
	).From(goqu.T(queueCore.WorkerTable).As(queueCore.WorkerAlias)).
		InnerJoin(
			goqu.T(queueCore.QueueTable).As(queueCore.QueueAlias),
			goqu.On(goqu.I(queueCore.WorkerAlias+".name").Eq(goqu.I(queueCore.QueueAlias+".worker_name"))),
		).
		Where(goqu.I(queueCore.QueueAlias + ".name").Eq(queueName))

	if forUpdate {
		query = query.ForUpdate(goqu.Wait)
	}

	workers, err := database.Select[queueCore.Worker](db, query)
	if err != nil {
		return nil, err
	}

	if len(workers) == 0 {
		return nil, nil
	}

	return &workers[0], nil
}

func (s *queueStore) SelectWorkerByQueue(queueName string, forUpdate bool) (*queueCore.Worker, error) {
	return SelectWorkerByQueue(s.ORPer, queueName, forUpdate)
}
