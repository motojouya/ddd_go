package core_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	basic "github.com/motojouya/ddd_go/pkg/basic/core"
	core "github.com/motojouya/ddd_go/pkg/queue/core"
)

func TestNewWorker_Success(t *testing.T) {
	worker, err := core.NewWorker("worker-1", 5)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := core.Worker{
		Name:       "worker-1",
		MaxProcess: 5,
	}

	if diff := cmp.Diff(expected, worker); diff != "" {
		t.Errorf("Worker mismatch (-want +got):\n%s", diff)
	}
}

func TestNewWorker_EmptyName(t *testing.T) {
	_, err := core.NewWorker("", 5)

	if err == nil {
		t.Fatal("expected error for empty name, got nil")
	}

	if _, ok := err.(basic.InvalidArgumentError); !ok {
		t.Errorf("expected InvalidArgumentError, got %T", err)
	}
}

func TestNewWorker_InvalidMaxProcess_Zero(t *testing.T) {
	_, err := core.NewWorker("worker-1", 0)

	if err == nil {
		t.Fatal("expected error for zero maxProcess, got nil")
	}

	if _, ok := err.(basic.InvalidArgumentError); !ok {
		t.Errorf("expected InvalidArgumentError, got %T", err)
	}
}

func TestNewWorker_InvalidMaxProcess_Over1000(t *testing.T) {
	_, err := core.NewWorker("worker-1", 1001)

	if err == nil {
		t.Fatal("expected error for maxProcess over 1000, got nil")
	}

	if _, ok := err.(basic.InvalidArgumentError); !ok {
		t.Errorf("expected InvalidArgumentError, got %T", err)
	}
}
