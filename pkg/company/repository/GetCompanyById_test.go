package repository_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	basic "github.com/motojouya/ddd_go/pkg/basic/model"
	"github.com/motojouya/ddd_go/pkg/company/repository"
	"github.com/motojouya/ddd_go/pkg/company/model"
	dbmock "github.com/motojouya/ddd_go/pkg/database/mock"
)

type mockIdGetter struct {
	ids []basic.Identifier
	err error
}

func (m mockIdGetter) GetId() ([]basic.Identifier, error) {
	return m.ids, m.err
}

func TestGetCompanyById_Success(t *testing.T) {
	expectedCompanies := []model.Company{
		{Id: "id-1", Code: "ABC12", Name: "Company A"},
		{Id: "id-2", Code: "XYZ99", Name: "Company B"},
	}

	var capturedConditions map[string][]interface{}
	executer := dbmock.SqlExecutorMock{
		FakeGetIn: func(records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error) {
			capturedConditions = conditions
			companies := records.(*[]model.Company)
			*companies = expectedCompanies
			return nil, nil
		},
	}

	idGetter := mockIdGetter{ids: []basic.Identifier{"id-1", "id-2"}}

	result, err := repository.GetCompanyById(executer, idGetter)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if diff := cmp.Diff(expectedCompanies, result); diff != "" {
		t.Errorf("Companies mismatch (-want +got):\n%s", diff)
	}

	expectedConditions := map[string][]interface{}{
		"id": {"id-1", "id-2"},
	}
	if diff := cmp.Diff(expectedConditions, capturedConditions); diff != "" {
		t.Errorf("conditions mismatch (-want +got):\n%s", diff)
	}
}

func TestGetCompanyById_EmptyIds(t *testing.T) {
	getInCalled := false
	executer := dbmock.SqlExecutorMock{
		FakeGetIn: func(records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error) {
			getInCalled = true
			return nil, nil
		},
	}
	idGetter := mockIdGetter{ids: []basic.Identifier{}}

	result, err := repository.GetCompanyById(executer, idGetter)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if len(result) != 0 {
		t.Errorf("Expected empty result, got %d companies", len(result))
	}
	if getInCalled {
		t.Error("expected GetIn not to be called when ids is empty")
	}
}

func TestGetCompanyById_GetterError(t *testing.T) {
	executer := dbmock.SqlExecutorMock{}
	idGetter := mockIdGetter{err: basic.NewInvalidArgumentError("id", "", "invalid id")}

	result, err := repository.GetCompanyById(executer, idGetter)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	if result != nil {
		t.Errorf("Expected nil result, got %v", result)
	}
}
