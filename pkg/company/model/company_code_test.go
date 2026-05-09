package core_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/motojouya/ddd_go/pkg/company/core"
	"github.com/motojouya/ddd_go/pkg/local/repository"
)

func TestNewCompanyCode_Valid(t *testing.T) {
	validCodes := []string{
		"ABC12",
		"XYZ99",
		"AAA00",
		"ZZZ11",
	}

	for _, codeStr := range validCodes {
		code, err := core.NewCompanyCode(codeStr)
		if err != nil {
			t.Errorf("Valid code should not return error: %s, got error: %v", codeStr, err)
		}
		if diff := cmp.Diff(codeStr, code.String()); diff != "" {
			t.Errorf("CompanyCode mismatch (-want +got):\n%s", diff)
		}
	}
}

func TestNewCompanyCode_Invalid(t *testing.T) {
	invalidCodes := []string{
		"abc12",  // lowercase letters
		"AB12",   // only 2 letters
		"ABCD12", // 4 letters
		"ABC1",   // only 1 digit
		"ABC123", // 3 digits
		"123AB",  // digits first
		"AB-12",  // special character
		"",       // empty
		"ABCDE",  // no digits
		"12345",  // only digits
	}

	for _, codeStr := range invalidCodes {
		code, err := core.NewCompanyCode(codeStr)
		if err == nil {
			t.Errorf("Invalid code should return error: %s", codeStr)
		}
		if diff := cmp.Diff("", code.String()); diff != "" {
			t.Errorf("Empty code expected for invalid input (-want +got):\n%s", diff)
		}
		t.Logf("Invalid code '%s' error: %s", codeStr, err.Error())
	}
}

func TestGenerateCompanyCode(t *testing.T) {
	localer := repository.CreateLocal()

	// Generate multiple codes to ensure they follow the pattern
	for i := 0; i < 10; i++ {
		code := core.GenerateCompanyCode(localer)

		// Verify the generated code is valid
		_, err := core.NewCompanyCode(code.String())
		if err != nil {
			t.Errorf("Generated code should be valid: %s, got error: %v", code.String(), err)
		}
		if got := len(code.String()); got != 5 {
			t.Errorf("Generated code length = %d, want 5", got)
		}

		t.Logf("Generated code: %s", code.String())
	}
}

func TestCompanyCode_String(t *testing.T) {
	code, _ := core.NewCompanyCode("ABC12")
	if diff := cmp.Diff("ABC12", code.String()); diff != "" {
		t.Errorf("CompanyCode.String() mismatch (-want +got):\n%s", diff)
	}
}
