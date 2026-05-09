package store

import (
	"github.com/doug-martin/goqu/v9"
	database "github.com/motojouya/ddd_go/pkg/database/model"
	queueModel "github.com/motojouya/ddd_go/pkg/queue/model"
)

func SelectWorkerByQueue(db database.Executor, queueName string, forUpdate bool) (*queueModel.Worker, error) {
	query := database.Dialect.Select(
		goqu.I(queueModel.WorkerAlias+".name"),
		goqu.I(queueModel.WorkerAlias+".max_process"),
	).From(goqu.T(queueModel.WorkerTable).As(queueModel.WorkerAlias)).
		InnerJoin(
			goqu.T(queueModel.QueueTable).As(queueModel.QueueAlias),
			goqu.On(goqu.I(queueModel.WorkerAlias+".name").Eq(goqu.I(queueModel.QueueAlias+".worker_name"))),
		).
		Where(goqu.I(queueModel.QueueAlias + ".name").Eq(queueName))

	if forUpdate {
		query = query.ForUpdate(goqu.Wait)
	}

	workers, err := database.Select[queueModel.Worker](db, query)
	if err != nil {
		return nil, err
	}

	if len(workers) == 0 {
		return nil, nil
	}

	return &workers[0], nil
}

func (s *queueStore) SelectWorkerByQueue(queueName string, forUpdate bool) (*queueModel.Worker, error) {
	return SelectWorkerByQueue(s.ORPer, queueName, forUpdate)
}
