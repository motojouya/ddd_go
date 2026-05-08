package behavior_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"
	basic "github.com/motojouya/ddd_go/pkg/basic/core"
	"github.com/motojouya/ddd_go/pkg/company/behavior"
	"github.com/motojouya/ddd_go/pkg/company/core"
	dbmock "github.com/motojouya/ddd_go/pkg/database/mock"
)

type mockCompanyCodeGetter struct {
	codes []core.CompanyCode
	err   error
}

func (m mockCompanyCodeGetter) GetCompanyCode() ([]core.CompanyCode, error) {
	return m.codes, m.err
}

func TestGetCompanyByCode_Success(t *testing.T) {
	expectedCompany := core.Company{Id: "id-1", Code: "ABC12", Name: "Company A"}

	var capturedConditions map[string][]interface{}
	executer := dbmock.SqlExecutorMock{
		FakeGetIn: func(records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error) {
			capturedConditions = conditions
			companies := records.(*[]core.Company)
			*companies = []core.Company{expectedCompany}
			return nil, nil
		},
	}

	code, _ := core.NewCompanyCode("ABC12")
	codeGetter := mockCompanyCodeGetter{codes: []core.CompanyCode{code}}

	result, err := behavior.GetCompanyByCode(executer, codeGetter)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	expected := []core.Company{expectedCompany}
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Errorf("Company mismatch (-want +got):\n%s", diff)
	}

	// conditions に code IN ("ABC12") が含まれていることを確認
	if codeConds, ok := capturedConditions["code"]; !ok {
		t.Error("expected 'code' key in conditions")
	} else if len(codeConds) != 1 || codeConds[0] != "ABC12" {
		t.Errorf("conditions[\"code\"] mismatch: got %v", codeConds)
	}
}

func TestGetCompanyByCode_MultipleCodes(t *testing.T) {
	companyA := core.Company{Id: "id-1", Code: "ABC12", Name: "Company A"}
	companyB := core.Company{Id: "id-2", Code: "DEF34", Name: "Company B"}

	var capturedConditions map[string][]interface{}
	executer := dbmock.SqlExecutorMock{
		FakeGetIn: func(records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error) {
			capturedConditions = conditions
			companies := records.(*[]core.Company)
			*companies = []core.Company{companyA, companyB}
			return nil, nil
		},
	}

	codeA, _ := core.NewCompanyCode("ABC12")
	codeB, _ := core.NewCompanyCode("DEF34")
	codeGetter := mockCompanyCodeGetter{codes: []core.CompanyCode{codeA, codeB}}

	result, err := behavior.GetCompanyByCode(executer, codeGetter)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	expected := []core.Company{companyA, companyB}
	if diff := cmp.Diff(expected, result); diff != "" {
		t.Errorf("Company mismatch (-want +got):\n%s", diff)
	}

	if codeConds, ok := capturedConditions["code"]; !ok {
		t.Error("expected 'code' key in conditions")
	} else if len(codeConds) != 2 {
		t.Errorf("conditions[\"code\"] expected 2 elements, got %d: %v", len(codeConds), codeConds)
	}
}

func TestGetCompanyByCode_EmptyCodes(t *testing.T) {
	getInCalled := false
	executer := dbmock.SqlExecutorMock{
		FakeGetIn: func(records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error) {
			getInCalled = true
			return nil, nil
		},
	}
	codeGetter := mockCompanyCodeGetter{codes: nil}

	result, err := behavior.GetCompanyByCode(executer, codeGetter)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	if len(result) != 0 {
		t.Errorf("Expected empty result, got %v", result)
	}
	if getInCalled {
		t.Error("expected GetIn not to be called when codes is empty")
	}
}

func TestGetCompanyByCode_GetterError(t *testing.T) {
	executer := dbmock.SqlExecutorMock{}
	getterErr := basic.NewInvalidArgumentError("code", "", "invalid code")
	codeGetter := mockCompanyCodeGetter{err: getterErr}

	result, err := behavior.GetCompanyByCode(executer, codeGetter)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
	var target basic.InvalidArgumentError
	if !errors.As(err, &target) {
		t.Errorf("expected InvalidArgumentError, got %T: %v", err, err)
	}
	if result != nil {
		t.Errorf("Expected nil result, got %v", result)
	}
}

func TestGetCompanyByCode_GetInError(t *testing.T) {
	getInErr := errors.New("get in failed")
	executer := dbmock.SqlExecutorMock{
		FakeGetIn: func(records interface{}, conditions map[string][]interface{}, forLock bool) ([]interface{}, error) {
			return nil, getInErr
		},
	}

	code, _ := core.NewCompanyCode("ABC12")
	codeGetter := mockCompanyCodeGetter{codes: []core.CompanyCode{code}}

	result, err := behavior.GetCompanyByCode(executer, codeGetter)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !errors.Is(err, getInErr) {
		t.Errorf("expected GetIn error to be returned, got %v", err)
	}
	if result != nil {
		t.Errorf("Expected nil result, got %v", result)
	}
}
