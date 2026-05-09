package model_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	basic "github.com/motojouya/ddd_go/pkg/basic/model"
	"github.com/motojouya/ddd_go/pkg/company/model"
)

func TestNewCompany(t *testing.T) {
	id := basic.Identifier("test-id-123")
	code, _ := model.NewCompanyCode("ABC12")
	name := "Test Company Name"

	company := model.NewCompany(id, code, name)

	if diff := cmp.Diff(id, company.Id); diff != "" {
		t.Errorf("Company.Id mismatch (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff(code, company.Code); diff != "" {
		t.Errorf("Company.Code mismatch (-want +got):\n%s", diff)
	}
	if diff := cmp.Diff(name, company.Name); diff != "" {
		t.Errorf("Company.Name mismatch (-want +got):\n%s", diff)
	}
}
