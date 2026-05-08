package repository_test

import (
	"errors"
	"testing"
	"time"

	basicCore "github.com/motojouya/ddd_go/pkg/basic/core"
	databaseMock "github.com/motojouya/ddd_go/pkg/database/mock"
	localMock "github.com/motojouya/ddd_go/pkg/local/mock"
	repository "github.com/motojouya/ddd_go/pkg/queue/repository"
	queueCore "github.com/motojouya/ddd_go/pkg/queue/core"
	queueStore "github.com/motojouya/ddd_go/pkg/queue/store"
)

func TestRegisterJob(t *testing.T) {
	queueName := "queue-1"
	workerName := "worker-1"
	source := "source-1"
	procedure := string(queueCore.ProcedureAllocate)
	jsonData := map[string]string{"key": "value"}
	jobID := basicCore.Identifier("job-id-001")
	now := time.Date(2026, 3, 13, 12, 0, 0, 0, time.UTC)

	queue := queueCore.Queue{
		Name:         queueName,
		WorkerName:   workerName,
		ProcessOrder: 1,
	}

	makeStoreMock := func(sqlMock databaseMock.SqlExecutorMock) *queueStore.QueueMock {
		return queueStore.NewQueueMock(makeORPerMock(sqlMock))
	}

	makeLocaler := func() localMock.LocalerMock {
		return localMock.LocalerMock{
			FakeGenerateID: func() (basicCore.Identifier, error) {
				return jobID, nil
			},
			FakeGetNow: func() time.Time {
				return now
			},
		}
	}

	t.Run("正常系: Queue存在 → ID生成 → Job作成 → Insert成功", func(t *testing.T) {
		var insertedJob interface{}
		sqlMock := databaseMock.SqlExecutorMock{
			FakeGetIn: func(records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error) {
				if v, ok := records.(*[]queueCore.Queue); ok {
					*v = []queueCore.Queue{queue}
				}
				return nil, nil
			},
			FakeInsert: func(list ...interface{}) error {
				if len(list) > 0 {
					insertedJob = list[0]
				}
				return nil
			},
		}

		storeMock := makeStoreMock(sqlMock)
		localer := makeLocaler()

		result, err := repository.RegisterJob(localer, storeMock, queueName, source, procedure, jsonData)

		if err != nil {
			t.Errorf("エラーが発生しました: %v", err)
		}
		if result.Id != jobID {
			t.Errorf("Job.Idが不正です。期待値: %v, 実際: %v", jobID, result.Id)
		}
		if result.Queue != queueName {
			t.Errorf("Job.Queueが不正です。期待値: %v, 実際: %v", queueName, result.Queue)
		}
		if result.Source != source {
			t.Errorf("Job.Sourceが不正です。期待値: %v, 実際: %v", source, result.Source)
		}
		if result.Procedure != procedure {
			t.Errorf("Job.Procedureが不正です。期待値: %v, 実際: %v", procedure, result.Procedure)
		}
		if insertedJob == nil {
			t.Error("Insertが呼ばれませんでした")
		}
	})

	t.Run("異常系: Queue存在しない（GetQueue returns nil）", func(t *testing.T) {
		sqlMock := databaseMock.SqlExecutorMock{
			FakeGetIn: func(records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error) {
				return nil, nil
			},
		}

		storeMock := makeStoreMock(sqlMock)
		localer := makeLocaler()

		_, err := repository.RegisterJob(localer, storeMock, queueName, source, procedure, jsonData)

		if err == nil {
			t.Error("エラーが期待されましたが、nilが返されました")
		}
	})

	t.Run("異常系: GetQueueエラー", func(t *testing.T) {
		expectedErr := errors.New("get queue error")
		sqlMock := databaseMock.SqlExecutorMock{
			FakeGetIn: func(records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error) {
				return nil, expectedErr
			},
		}

		storeMock := makeStoreMock(sqlMock)
		localer := makeLocaler()

		_, err := repository.RegisterJob(localer, storeMock, queueName, source, procedure, jsonData)

		if err != expectedErr {
			t.Errorf("期待したエラーが返されませんでした。期待値: %v, 実際: %v", expectedErr, err)
		}
	})

	t.Run("異常系: GenerateIDエラー", func(t *testing.T) {
		expectedErr := errors.New("generate id error")
		sqlMock := databaseMock.SqlExecutorMock{
			FakeGetIn: func(records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error) {
				if v, ok := records.(*[]queueCore.Queue); ok {
					*v = []queueCore.Queue{queue}
				}
				return nil, nil
			},
		}

		storeMock := makeStoreMock(sqlMock)
		localer := localMock.LocalerMock{
			FakeGenerateID: func() (basicCore.Identifier, error) {
				return basicCore.Identifier(""), expectedErr
			},
			FakeGetNow: func() time.Time {
				return now
			},
		}

		_, err := repository.RegisterJob(localer, storeMock, queueName, source, procedure, jsonData)

		if err != expectedErr {
			t.Errorf("期待したエラーが返されませんでした。期待値: %v, 実際: %v", expectedErr, err)
		}
	})

	t.Run("異常系: NewJobエラー（procedure不正）", func(t *testing.T) {
		sqlMock := databaseMock.SqlExecutorMock{
			FakeGetIn: func(records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error) {
				if v, ok := records.(*[]queueCore.Queue); ok {
					*v = []queueCore.Queue{queue}
				}
				return nil, nil
			},
		}

		storeMock := makeStoreMock(sqlMock)
		localer := makeLocaler()

		_, err := repository.RegisterJob(localer, storeMock, queueName, source, "INVALID_PROCEDURE", jsonData)

		if err == nil {
			t.Error("エラーが期待されましたが、nilが返されました")
		}
	})

	t.Run("異常系: Insertエラー", func(t *testing.T) {
		expectedErr := errors.New("insert error")
		sqlMock := databaseMock.SqlExecutorMock{
			FakeGetIn: func(records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error) {
				if v, ok := records.(*[]queueCore.Queue); ok {
					*v = []queueCore.Queue{queue}
				}
				return nil, nil
			},
			FakeInsert: func(list ...interface{}) error {
				return expectedErr
			},
		}

		storeMock := makeStoreMock(sqlMock)
		localer := makeLocaler()

		_, err := repository.RegisterJob(localer, storeMock, queueName, source, procedure, jsonData)

		if err != expectedErr {
			t.Errorf("期待したエラーが返されませんでした。期待値: %v, 実際: %v", expectedErr, err)
		}
	})
}
