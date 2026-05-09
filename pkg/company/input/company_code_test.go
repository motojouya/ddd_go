package entry_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/motojouya/ddd_go/pkg/company/core"
	"github.com/motojouya/ddd_go/pkg/company/entry"
)

func TestCompanyCode_GetCompanyCode_Valid(t *testing.T) {
	param := entry.CompanyCode{Code: "ABC12"}

	codes, err := param.GetCompanyCode()
	if err != nil {
		t.Errorf("Valid code should not return error: %v", err)
	}

	expected := []core.CompanyCode{core.CompanyCode("ABC12")}
	if diff := cmp.Diff(expected, codes); diff != "" {
		t.Errorf("CompanyCode mismatch (-want +got):\n%s", diff)
	}
}

func TestCompanyCode_GetCompanyCode_Empty(t *testing.T) {
	param := entry.CompanyCode{Code: ""}

	codes, err := param.GetCompanyCode()
	if err == nil {
		t.Errorf("Empty code should return error")
	}

	if codes != nil {
		t.Errorf("Expected nil codes for invalid input, got %v", codes)
	}
}

func TestCompanyCode_GetCompanyCode_Invalid(t *testing.T) {
	invalidCodes := []string{
		"abc12",
		"AB12",
		"ABCD12",
		"ABC1",
		"ABC123",
	}

	for _, code := range invalidCodes {
		param := entry.CompanyCode{Code: code}

		resultCodes, err := param.GetCompanyCode()
		if err == nil {
			t.Errorf("Invalid code should return error: %s", code)
		}
		if resultCodes != nil {
			t.Errorf("Expected nil codes for invalid code, got %v", resultCodes)
		}
		t.Logf("Invalid code '%s' error: %s", code, err.Error())
	}
}
