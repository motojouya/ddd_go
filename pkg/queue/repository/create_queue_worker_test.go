package repository_test

import (
	"errors"
	"testing"

	databaseMock "github.com/motojouya/ddd_go/pkg/database/mock"
	repository "github.com/motojouya/ddd_go/pkg/queue/repository"
	queueCore "github.com/motojouya/ddd_go/pkg/queue/core"
	queueStore "github.com/motojouya/ddd_go/pkg/queue/store"
)

type ORPerMock struct {
	databaseMock.TransactionalDatabaseMock
	databaseMock.SqlExecutorMock
}

func makeORPerMock(sqlMock databaseMock.SqlExecutorMock) *ORPerMock {
	calledCount := &databaseMock.TransactionCalledCount{}
	return &ORPerMock{
		TransactionalDatabaseMock: databaseMock.GetTransactionalDatabaseMock(calledCount),
		SqlExecutorMock:           sqlMock,
	}
}

func TestCreateWorker(t *testing.T) {
	workerName := "worker-1"
	maxProcess := 5

	t.Run("正常系: Worker新規作成", func(t *testing.T) {
		var insertCalled bool
		sqlMock := databaseMock.SqlExecutorMock{
			FakeGet: func(i interface{}, keys ...interface{}) (interface{}, error) {
				return (*queueCore.Worker)(nil), nil
			},
			FakeInsert: func(list ...interface{}) error {
				insertCalled = true
				return nil
			},
		}

		storeMock := queueStore.NewQueueMock(makeORPerMock(sqlMock))
		result, err := repository.CreateWorker(storeMock, workerName, maxProcess)

		if err != nil {
			t.Errorf("エラーが発生しました: %v", err)
		}
		if result.Name != workerName {
			t.Errorf("Worker.Nameが不正です。期待値: %v, 実際: %v", workerName, result.Name)
		}
		if result.MaxProcess != maxProcess {
			t.Errorf("Worker.MaxProcessが不正です。期待値: %v, 実際: %v", maxProcess, result.MaxProcess)
		}
		if !insertCalled {
			t.Error("Insertが呼ばれませんでした")
		}
	})

	t.Run("正常系: Worker既存（Insertされない）", func(t *testing.T) {
		existing := queueCore.Worker{Name: workerName, MaxProcess: maxProcess}
		var insertCalled bool
		sqlMock := databaseMock.SqlExecutorMock{
			FakeGet: func(i interface{}, keys ...interface{}) (interface{}, error) {
				return &existing, nil
			},
			FakeInsert: func(list ...interface{}) error {
				insertCalled = true
				return nil
			},
		}

		storeMock := queueStore.NewQueueMock(makeORPerMock(sqlMock))
		result, err := repository.CreateWorker(storeMock, workerName, maxProcess)

		if err != nil {
			t.Errorf("エラーが発生しました: %v", err)
		}
		if result.Name != workerName {
			t.Errorf("Worker.Nameが不正です。期待値: %v, 実際: %v", workerName, result.Name)
		}
		if insertCalled {
			t.Error("既存Workerに対してInsertが呼ばれました")
		}
	})

	t.Run("異常系: workerName空文字（NewWorker失敗）", func(t *testing.T) {
		sqlMock := databaseMock.SqlExecutorMock{}
		storeMock := queueStore.NewQueueMock(makeORPerMock(sqlMock))

		_, err := repository.CreateWorker(storeMock, "", maxProcess)

		if err == nil {
			t.Error("エラーが期待されましたが、nilが返されました")
		}
	})

	t.Run("異常系: maxProcess不正（NewWorker失敗）", func(t *testing.T) {
		sqlMock := databaseMock.SqlExecutorMock{}
		storeMock := queueStore.NewQueueMock(makeORPerMock(sqlMock))

		_, err := repository.CreateWorker(storeMock, workerName, 0)

		if err == nil {
			t.Error("エラーが期待されましたが、nilが返されました")
		}
	})

	t.Run("異常系: GetOrCreate時のGetエラー", func(t *testing.T) {
		expectedErr := errors.New("get error")
		sqlMock := databaseMock.SqlExecutorMock{
			FakeGet: func(i interface{}, keys ...interface{}) (interface{}, error) {
				return nil, expectedErr
			},
		}

		storeMock := queueStore.NewQueueMock(makeORPerMock(sqlMock))
		_, err := repository.CreateWorker(storeMock, workerName, maxProcess)

		if err != expectedErr {
			t.Errorf("期待したエラーが返されませんでした。期待値: %v, 実際: %v", expectedErr, err)
		}
	})
}

func TestCreateQueue(t *testing.T) {
	workerName := "worker-1"
	queueName := "queue-1"
	processOrder := 1

	worker := queueCore.Worker{Name: workerName, MaxProcess: 5}

	t.Run("正常系: Worker存在 → Queue新規作成", func(t *testing.T) {
		var insertCalled bool
		sqlMock := databaseMock.SqlExecutorMock{
			FakeGetIn: func(records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error) {
				if v, ok := records.(*[]queueCore.Worker); ok {
					*v = []queueCore.Worker{worker}
				}
				return nil, nil
			},
			FakeGet: func(i interface{}, keys ...interface{}) (interface{}, error) {
				return (*queueCore.Queue)(nil), nil
			},
			FakeInsert: func(list ...interface{}) error {
				insertCalled = true
				return nil
			},
		}

		storeMock := queueStore.NewQueueMock(makeORPerMock(sqlMock))
		result, err := repository.CreateQueue(storeMock, workerName, queueName, processOrder)

		if err != nil {
			t.Errorf("エラーが発生しました: %v", err)
		}
		if result.Name != queueName {
			t.Errorf("Queue.Nameが不正です。期待値: %v, 実際: %v", queueName, result.Name)
		}
		if result.WorkerName != workerName {
			t.Errorf("Queue.WorkerNameが不正です。期待値: %v, 実際: %v", workerName, result.WorkerName)
		}
		if result.ProcessOrder != processOrder {
			t.Errorf("Queue.ProcessOrderが不正です。期待値: %v, 実際: %v", processOrder, result.ProcessOrder)
		}
		if !insertCalled {
			t.Error("Insertが呼ばれませんでした")
		}
	})

	t.Run("異常系: Worker存在しない", func(t *testing.T) {
		sqlMock := databaseMock.SqlExecutorMock{
			FakeGetIn: func(records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error) {
				return nil, nil
			},
		}

		storeMock := queueStore.NewQueueMock(makeORPerMock(sqlMock))
		_, err := repository.CreateQueue(storeMock, workerName, queueName, processOrder)

		if err == nil {
			t.Error("エラーが期待されましたが、nilが返されました")
		}
	})

	t.Run("異常系: GetWorkerエラー", func(t *testing.T) {
		expectedErr := errors.New("get worker error")
		sqlMock := databaseMock.SqlExecutorMock{
			FakeGetIn: func(records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error) {
				return nil, expectedErr
			},
		}

		storeMock := queueStore.NewQueueMock(makeORPerMock(sqlMock))
		_, err := repository.CreateQueue(storeMock, workerName, queueName, processOrder)

		if err != expectedErr {
			t.Errorf("期待したエラーが返されませんでした。期待値: %v, 実際: %v", expectedErr, err)
		}
	})

	t.Run("異常系: queueName空文字（NewQueue失敗）", func(t *testing.T) {
		sqlMock := databaseMock.SqlExecutorMock{
			FakeGetIn: func(records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error) {
				if v, ok := records.(*[]queueCore.Worker); ok {
					*v = []queueCore.Worker{worker}
				}
				return nil, nil
			},
		}

		storeMock := queueStore.NewQueueMock(makeORPerMock(sqlMock))
		_, err := repository.CreateQueue(storeMock, workerName, "", processOrder)

		if err == nil {
			t.Error("エラーが期待されましたが、nilが返されました")
		}
	})

	t.Run("異常系: processOrder不正（NewQueue失敗）", func(t *testing.T) {
		sqlMock := databaseMock.SqlExecutorMock{
			FakeGetIn: func(records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error) {
				if v, ok := records.(*[]queueCore.Worker); ok {
					*v = []queueCore.Worker{worker}
				}
				return nil, nil
			},
		}

		storeMock := queueStore.NewQueueMock(makeORPerMock(sqlMock))
		_, err := repository.CreateQueue(storeMock, workerName, queueName, 0)

		if err == nil {
			t.Error("エラーが期待されましたが、nilが返されました")
		}
	})

	t.Run("異常系: Queue GetOrCreate時のGetエラー", func(t *testing.T) {
		expectedErr := errors.New("queue get error")
		sqlMock := databaseMock.SqlExecutorMock{
			FakeGetIn: func(records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error) {
				if v, ok := records.(*[]queueCore.Worker); ok {
					*v = []queueCore.Worker{worker}
				}
				return nil, nil
			},
			FakeGet: func(i interface{}, keys ...interface{}) (interface{}, error) {
				return nil, expectedErr
			},
		}

		storeMock := queueStore.NewQueueMock(makeORPerMock(sqlMock))
		_, err := repository.CreateQueue(storeMock, workerName, queueName, processOrder)

		if err != expectedErr {
			t.Errorf("期待したエラーが返されませんでした。期待値: %v, 実際: %v", expectedErr, err)
		}
	})
}
