package model_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	basic "github.com/motojouya/ddd_go/pkg/basic/model"
	model "github.com/motojouya/ddd_go/pkg/queue/model"
)

func TestNewQueue_Success(t *testing.T) {
	worker := model.Worker{Name: "worker-1", MaxProcess: 5}
	queue, err := model.NewQueue("queue-1", 1, worker)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := model.Queue{
		Name:         "queue-1",
		WorkerName:   "worker-1",
		ProcessOrder: 1,
	}

	if diff := cmp.Diff(expected, queue); diff != "" {
		t.Errorf("Queue mismatch (-want +got):\n%s", diff)
	}
}

func TestNewQueue_EmptyName(t *testing.T) {
	worker := model.Worker{Name: "worker-1", MaxProcess: 5}
	_, err := model.NewQueue("", 1, worker)

	if err == nil {
		t.Fatal("expected error for empty name, got nil")
	}

	if _, ok := err.(basic.InvalidArgumentError); !ok {
		t.Errorf("expected InvalidArgumentError, got %T", err)
	}
}

func TestNewQueue_EmptyWorkerName(t *testing.T) {
	worker := model.Worker{Name: ""}
	_, err := model.NewQueue("queue-1", 1, worker)

	if err == nil {
		t.Fatal("expected error for empty workerName, got nil")
	}

	if _, ok := err.(basic.InvalidArgumentError); !ok {
		t.Errorf("expected InvalidArgumentError, got %T", err)
	}
}

func TestNewQueue_InvalidProcessOrder(t *testing.T) {
	worker := model.Worker{Name: "worker-1", MaxProcess: 5}
	_, err := model.NewQueue("queue-1", 0, worker)

	if err == nil {
		t.Fatal("expected error for zero processOrder, got nil")
	}

	if _, ok := err.(basic.InvalidArgumentError); !ok {
		t.Errorf("expected InvalidArgumentError, got %T", err)
	}
}
