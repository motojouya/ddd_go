package behavior_test

import (
	"errors"
	"testing"

	basic "github.com/motojouya/ddd_go/pkg/basic/core"
	databaseCore "github.com/motojouya/ddd_go/pkg/database/core"
	databaseMock "github.com/motojouya/ddd_go/pkg/database/mock"
	behavior "github.com/motojouya/ddd_go/pkg/queue/behavior"
	queueCore "github.com/motojouya/ddd_go/pkg/queue/core"
)

func TestGetQueue(t *testing.T) {
	queue1 := queueCore.Queue{
		Name:       "queue-1",
		WorkerName: "worker-1",
	}
	queue2 := queueCore.Queue{
		Name:       "queue-2",
		WorkerName: "worker-1",
	}

	t.Run("正常系: Queueを取得できる", func(t *testing.T) {
		sqlMock := databaseMock.SqlExecutorMock{
			FakeGetIn: func(records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error) {
				if v, ok := records.(*[]queueCore.Queue); ok {
					*v = []queueCore.Queue{queue1, queue2}
				}
				return nil, nil
			},
		}

		result, err := behavior.GetQueue(sqlMock, queue1.Name, false)

		if err != nil {
			t.Errorf("エラーが発生しました: %v", err)
		}
		if result == nil {
			t.Fatal("nilが返されました")
		}
		if result.Name != queue1.Name {
			t.Errorf("Nameが不正です。期待値: %v, 実際: %v", queue1.Name, result.Name)
		}
	})

	t.Run("正常系: 存在しない場合はnilを返す", func(t *testing.T) {
		sqlMock := databaseMock.SqlExecutorMock{
			FakeGetIn: func(records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error) {
				return nil, nil
			},
		}

		result, err := behavior.GetQueue(sqlMock, "no-such-queue", false)

		if err != nil {
			t.Errorf("エラーが発生しました: %v", err)
		}
		if result != nil {
			t.Errorf("nilが期待されましたが、%v が返されました", result)
		}
	})

	t.Run("異常系: GetInでエラー", func(t *testing.T) {
		expectedErr := errors.New("store error")
		sqlMock := databaseMock.SqlExecutorMock{
			FakeGetIn: func(records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error) {
				return nil, expectedErr
			},
		}

		_, err := behavior.GetQueue(sqlMock, queue1.Name, false)

		if err != expectedErr {
			t.Errorf("期待したエラーが返されませんでした。期待値: %v, 実際: %v", expectedErr, err)
		}
	})
}

func TestGetWorker(t *testing.T) {
	worker1 := queueCore.Worker{
		Name: "worker-1",
	}

	t.Run("正常系: Workerを取得できる", func(t *testing.T) {
		sqlMock := databaseMock.SqlExecutorMock{
			FakeGetIn: func(records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error) {
				if v, ok := records.(*[]queueCore.Worker); ok {
					*v = []queueCore.Worker{worker1}
				}
				return nil, nil
			},
		}

		result, err := behavior.GetWorker(sqlMock, worker1.Name, false)

		if err != nil {
			t.Errorf("エラーが発生しました: %v", err)
		}
		if result == nil {
			t.Fatal("nilが返されました")
		}
		if result.Name != worker1.Name {
			t.Errorf("Nameが不正です。期待値: %v, 実際: %v", worker1.Name, result.Name)
		}
	})

	t.Run("正常系: 存在しない場合はnilを返す", func(t *testing.T) {
		sqlMock := databaseMock.SqlExecutorMock{
			FakeGetIn: func(records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error) {
				return nil, nil
			},
		}

		result, err := behavior.GetWorker(sqlMock, "no-such-worker", false)

		if err != nil {
			t.Errorf("エラーが発生しました: %v", err)
		}
		if result != nil {
			t.Errorf("nilが期待されましたが、%v が返されました", result)
		}
	})

	t.Run("異常系: GetInでエラー", func(t *testing.T) {
		expectedErr := errors.New("store error")
		sqlMock := databaseMock.SqlExecutorMock{
			FakeGetIn: func(records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error) {
				return nil, expectedErr
			},
		}

		_, err := behavior.GetWorker(sqlMock, worker1.Name, false)

		if err != expectedErr {
			t.Errorf("期待したエラーが返されませんでした。期待値: %v, 実際: %v", expectedErr, err)
		}
	})
}

func TestGetQueueByWorker(t *testing.T) {
	workerName := "worker-1"
	queue1 := queueCore.Queue{
		Name:         "queue-1",
		WorkerName:   workerName,
		ProcessOrder: 1,
	}
	queue2 := queueCore.Queue{
		Name:         "queue-2",
		WorkerName:   workerName,
		ProcessOrder: 2,
	}

	t.Run("正常系: 正しい条件・順序・ページャでGetPagingが呼ばれる", func(t *testing.T) {
		var capturedConditions map[string]interface{}
		var capturedOrders []databaseCore.Order
		var capturedPager basic.Pager

		sqlMock := databaseMock.SqlExecutorMock{
			FakeGetPaging: func(records interface{}, conditions map[string]interface{}, orders []databaseCore.Order, pager basic.Pager) ([]interface{}, error) {
				capturedConditions = conditions
				capturedOrders = orders
				capturedPager = pager
				if v, ok := records.(*[]queueCore.Queue); ok {
					*v = []queueCore.Queue{queue1, queue2}
				}
				return nil, nil
			},
		}

		result, err := behavior.GetQueueByWorker(sqlMock, workerName, false)

		if err != nil {
			t.Errorf("エラーが発生しました: %v", err)
		}
		if len(result) != 2 {
			t.Errorf("取得したQueue数が不正です。期待値: 2, 実際: %d", len(result))
		}
		if result[0].Name != queue1.Name {
			t.Errorf("Nameが不正です。期待値: %v, 実際: %v", queue1.Name, result[0].Name)
		}
		if result[1].Name != queue2.Name {
			t.Errorf("Nameが不正です。期待値: %v, 実際: %v", queue2.Name, result[1].Name)
		}

		// 条件の検証
		if v, ok := capturedConditions["worker_name"]; !ok || v != workerName {
			t.Errorf("worker_nameの条件が不正です。期待値: %v, 実際: %v", workerName, capturedConditions["worker_name"])
		}

		// 順序の検証
		if len(capturedOrders) != 1 {
			t.Errorf("Orderの数が不正です。期待値: 1, 実際: %d", len(capturedOrders))
		} else {
			if capturedOrders[0].Column != "process_order" {
				t.Errorf("Order列が不正です。期待値: process_order, 実際: %v", capturedOrders[0].Column)
			}
			if !capturedOrders[0].Ascending {
				t.Error("Orderが昇順でありません")
			}
		}

		// ページャの検証
		if capturedPager.Cursor != 1 {
			t.Errorf("Pager.Cursorが不正です。期待値: 1, 実際: %d", capturedPager.Cursor)
		}
		if capturedPager.Limit != 1000 {
			t.Errorf("Pager.Limitが不正です。期待値: 1000, 実際: %d", capturedPager.Limit)
		}
	})

	t.Run("正常系: 存在しない場合は空スライスを返す", func(t *testing.T) {
		sqlMock := databaseMock.SqlExecutorMock{
			FakeGetPaging: func(records interface{}, conditions map[string]interface{}, orders []databaseCore.Order, pager basic.Pager) ([]interface{}, error) {
				return nil, nil
			},
		}

		result, err := behavior.GetQueueByWorker(sqlMock, "no-such-worker", false)

		if err != nil {
			t.Errorf("エラーが発生しました: %v", err)
		}
		if len(result) != 0 {
			t.Errorf("取得したQueue数が不正です。期待値: 0, 実際: %d", len(result))
		}
	})

	t.Run("異常系: GetPagingでエラー", func(t *testing.T) {
		expectedErr := errors.New("store error")
		sqlMock := databaseMock.SqlExecutorMock{
			FakeGetPaging: func(records interface{}, conditions map[string]interface{}, orders []databaseCore.Order, pager basic.Pager) ([]interface{}, error) {
				return nil, expectedErr
			},
		}

		_, err := behavior.GetQueueByWorker(sqlMock, workerName, false)

		if err != expectedErr {
			t.Errorf("期待したエラーが返されませんでした。期待値: %v, 実際: %v", expectedErr, err)
		}
	})
}

func TestGetJobByQueue(t *testing.T) {
	queueName := "queue-1"
	job1 := queueCore.Job{
		Id:        basic.Identifier("job-1"),
		Queue:     queueName,
		Source:    "source-1",
		Procedure: string(queueCore.ProcedureAllocate),
	}
	job2 := queueCore.Job{
		Id:        basic.Identifier("job-2"),
		Queue:     queueName,
		Source:    "source-2",
		Procedure: string(queueCore.ProcedureAllocate),
	}

	t.Run("正常系: 正しい条件・順序・ページャでGetPagingが呼ばれる", func(t *testing.T) {
		limit := 5
		var capturedConditions map[string]interface{}
		var capturedOrders []databaseCore.Order
		var capturedPager basic.Pager

		sqlMock := databaseMock.SqlExecutorMock{
			FakeGetPaging: func(records interface{}, conditions map[string]interface{}, orders []databaseCore.Order, pager basic.Pager) ([]interface{}, error) {
				capturedConditions = conditions
				capturedOrders = orders
				capturedPager = pager
				if v, ok := records.(*[]queueCore.Job); ok {
					*v = []queueCore.Job{job1, job2}
				}
				return nil, nil
			},
		}

		result, err := behavior.GetJobByQueue(sqlMock, queueName, limit, false)

		if err != nil {
			t.Errorf("エラーが発生しました: %v", err)
		}
		if len(result) != 2 {
			t.Errorf("取得したJob数が不正です。期待値: 2, 実際: %d", len(result))
		}
		if result[0].Id != job1.Id {
			t.Errorf("IDが不正です。期待値: %v, 実際: %v", job1.Id, result[0].Id)
		}
		if result[1].Id != job2.Id {
			t.Errorf("IDが不正です。期待値: %v, 実際: %v", job2.Id, result[1].Id)
		}

		// 条件の検証
		if v, ok := capturedConditions["queue"]; !ok || v != queueName {
			t.Errorf("queueの条件が不正です。期待値: %v, 実際: %v", queueName, capturedConditions["queue"])
		}
		if v, ok := capturedConditions["start_date"]; !ok || v != nil {
			t.Errorf("start_dateの条件が不正です。期待値: nil, 実際: %v", capturedConditions["start_date"])
		}

		// 順序の検証
		if len(capturedOrders) != 1 {
			t.Errorf("Orderの数が不正です。期待値: 1, 実際: %d", len(capturedOrders))
		} else {
			if capturedOrders[0].Column != "register_date" {
				t.Errorf("Order列が不正です。期待値: register_date, 実際: %v", capturedOrders[0].Column)
			}
			if !capturedOrders[0].Ascending {
				t.Error("Orderが昇順でありません")
			}
		}

		// ページャの検証
		if capturedPager.Cursor != 1 {
			t.Errorf("Pager.Cursorが不正です。期待値: 1, 実際: %d", capturedPager.Cursor)
		}
		if capturedPager.Limit != uint(limit) {
			t.Errorf("Pager.Limitが不正です。期待値: %d, 実際: %d", limit, capturedPager.Limit)
		}
	})

	t.Run("正常系: 存在しない場合は空スライスを返す", func(t *testing.T) {
		sqlMock := databaseMock.SqlExecutorMock{
			FakeGetPaging: func(records interface{}, conditions map[string]interface{}, orders []databaseCore.Order, pager basic.Pager) ([]interface{}, error) {
				return nil, nil
			},
		}

		result, err := behavior.GetJobByQueue(sqlMock, queueName, 10, false)

		if err != nil {
			t.Errorf("エラーが発生しました: %v", err)
		}
		if len(result) != 0 {
			t.Errorf("取得したJob数が不正です。期待値: 0, 実際: %d", len(result))
		}
	})

	t.Run("異常系: GetPagingでエラー", func(t *testing.T) {
		expectedErr := errors.New("store error")
		sqlMock := databaseMock.SqlExecutorMock{
			FakeGetPaging: func(records interface{}, conditions map[string]interface{}, orders []databaseCore.Order, pager basic.Pager) ([]interface{}, error) {
				return nil, expectedErr
			},
		}

		_, err := behavior.GetJobByQueue(sqlMock, queueName, 10, false)

		if err != expectedErr {
			t.Errorf("期待したエラーが返されませんでした。期待値: %v, 実際: %v", expectedErr, err)
		}
	})
}
