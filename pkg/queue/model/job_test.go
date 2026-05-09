package model_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	basic "github.com/motojouya/ddd_go/pkg/basic/model"
	model "github.com/motojouya/ddd_go/pkg/queue/model"
)

func TestNewJob_Success(t *testing.T) {
	id := basic.Identifier("job-123")
	queue := model.Queue{Name: "queue-1", WorkerName: "worker-1"}
	source := "source-1"
	procedure := "LOCATION_STOCK_PROCEDURE_ALLOCATE"
	jsonData := map[string]string{"key": "value"}
	registerDate := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)

	job, err := model.NewJob(id, queue, source, procedure, jsonData, registerDate)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := model.Job{
		Id:           basic.Identifier("job-123"),
		Queue:        "queue-1",
		Source:       "source-1",
		Procedure:    "LOCATION_STOCK_PROCEDURE_ALLOCATE",
		JsonParams:   `{"key":"value"}`,
		JsonResult:   "",
		ErrorJson:    "",
		RegisterDate: registerDate,
		StartDate:    nil,
		FinishDate:   nil,
		StatusCode:   false,
	}

	if diff := cmp.Diff(expected, job); diff != "" {
		t.Errorf("Job mismatch (-want +got):\n%s", diff)
	}
}

func TestNewJob_InvalidProcedure(t *testing.T) {
	id := basic.Identifier("job-123")
	queue := model.Queue{Name: "queue-1", WorkerName: "worker-1"}
	registerDate := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)

	_, err := model.NewJob(id, queue, "source-1", "INVALID", nil, registerDate)

	if err == nil {
		t.Fatal("expected error for invalid procedure, got nil")
	}

	if _, ok := err.(basic.InvalidArgumentError); !ok {
		t.Errorf("expected InvalidArgumentError, got %T", err)
	}
}

func TestNewJob_EmptyQueue(t *testing.T) {
	id := basic.Identifier("job-123")
	queue := model.Queue{Name: "", WorkerName: "worker-1"}
	registerDate := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)

	_, err := model.NewJob(id, queue, "source-1", "LOCATION_STOCK_PROCEDURE_ALLOCATE", nil, registerDate)

	if err == nil {
		t.Fatal("expected error for empty queue, got nil")
	}

	if _, ok := err.(basic.InvalidArgumentError); !ok {
		t.Errorf("expected InvalidArgumentError, got %T", err)
	}
}

func TestNewJob_EmptySource(t *testing.T) {
	id := basic.Identifier("job-123")
	queue := model.Queue{Name: "queue-1", WorkerName: "worker-1"}
	registerDate := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)

	_, err := model.NewJob(id, queue, "", "LOCATION_STOCK_PROCEDURE_ALLOCATE", nil, registerDate)

	if err == nil {
		t.Fatal("expected error for empty source, got nil")
	}

	if _, ok := err.(basic.InvalidArgumentError); !ok {
		t.Errorf("expected InvalidArgumentError, got %T", err)
	}
}

func TestGetProcedure_Valid(t *testing.T) {
	err := model.GetProcedure("LOCATION_STOCK_PROCEDURE_ALLOCATE")

	if err != nil {
		t.Errorf("unexpected error for valid procedure: %v", err)
	}
}

func TestGetProcedure_Invalid(t *testing.T) {
	err := model.GetProcedure("INVALID")

	if err == nil {
		t.Fatal("expected error for invalid procedure, got nil")
	}
}

func TestJob_Keys(t *testing.T) {
	job := model.Job{Id: basic.Identifier("job-123")}

	keys := job.Keys()

	expected := []interface{}{basic.Identifier("job-123")}

	if diff := cmp.Diff(expected, keys); diff != "" {
		t.Errorf("Keys mismatch (-want +got):\n%s", diff)
	}
}

func TestJob_GetId(t *testing.T) {
	job := model.Job{Id: basic.Identifier("job-123")}

	ids, err := job.GetId()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []basic.Identifier{"job-123"}

	if diff := cmp.Diff(expected, ids); diff != "" {
		t.Errorf("GetId mismatch (-want +got):\n%s", diff)
	}
}
