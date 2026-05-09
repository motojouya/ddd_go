package model_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	basic "github.com/motojouya/ddd_go/pkg/basic/model"
	model "github.com/motojouya/ddd_go/pkg/company/model"
)

func TestNewCompanySource_Success(t *testing.T) {
	id := basic.Identifier("source-123")
	companyCode := model.CompanyCode("ABC12")
	num := uint(1)

	source, err := model.NewCompanySource(model.SourceTypeItem, id, companyCode, num)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := model.CompanySource{
		SourceType:  model.SourceTypeItem,
		Id:          "source-123",
		CompanyCode: "ABC12",
		Num:         1,
	}

	if diff := cmp.Diff(expected, source); diff != "" {
		t.Errorf("CompanySource mismatch (-want +got):\n%s", diff)
	}
}

func TestNewCompanySource_InvalidSourceType(t *testing.T) {
	id := basic.Identifier("source-123")
	companyCode := model.CompanyCode("ABC12")
	num := uint(1)

	_, err := model.NewCompanySource(model.SourceType("INVALID"), id, companyCode, num)

	if err == nil {
		t.Fatal("expected error for invalid source type, got nil")
	}

	if _, ok := err.(basic.InvalidArgumentError); !ok {
		t.Errorf("expected InvalidArgumentError, got %T", err)
	}
}

func TestGetSourceType_Valid(t *testing.T) {
	validTypes := []model.SourceType{
		model.SourceTypeCompany,
		model.SourceTypeImage,
		model.SourceTypeItem,
	}

	for _, st := range validTypes {
		if err := model.GetSourceType(st); err != nil {
			t.Errorf("expected no error for %s, got %v", st, err)
		}
	}
}

func TestGetSourceType_Invalid(t *testing.T) {
	err := model.GetSourceType(model.SourceType("INVALID"))
	if err == nil {
		t.Fatal("expected error for invalid source type, got nil")
	}
}
