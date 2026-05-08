package store_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	util "github.com/motojouya/ddd_go/pkg/database/test"
	queueCore "github.com/motojouya/ddd_go/pkg/queue/core"
	"github.com/motojouya/ddd_go/pkg/queue/store"
)

// truncateの順序を決めているので、順番が重要。依存される側があとに来るようにする。
var tables = []string{
	"job",
	"queue",
	"worker",
}

func TestSelectWorkerByQueue_Found(t *testing.T) {
	util.Truncate(t, orp, tables)

	workerRecords := []queueCore.Worker{
		{Name: "worker1"},
	}
	util.Ready(t, orp, workerRecords)

	queueRecords := []queueCore.Queue{
		{Name: "queue1", WorkerName: "worker1"},
	}
	util.Ready(t, orp, queueRecords)

	result, err := store.SelectWorkerByQueue(orp, "queue1", false)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	expected := &queueCore.Worker{Name: "worker1"}

	if diff := cmp.Diff(expected, result); diff != "" {
		t.Errorf("SelectWorkerByQueue result mismatch (-want +got):\n%s", diff)
	}
}

func TestSelectWorkerByQueue_NotFound(t *testing.T) {
	util.Truncate(t, orp, tables)

	result, err := store.SelectWorkerByQueue(orp, "nonexistent_queue", false)
	if err != nil {
		t.Fatalf("Unexpected error: %s", err)
	}

	if result != nil {
		t.Errorf("Expected nil, got %+v", result)
	}
}
