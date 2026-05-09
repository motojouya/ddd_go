package model

import (
	"database/sql/driver"
	"regexp"

	basic "github.com/motojouya/ddd_go/pkg/basic/model"
	"github.com/motojouya/ddd_go/pkg/local/repository"
)

type CompanyCode string

var companyCodePattern = regexp.MustCompile(`^[A-Z]{3}\d{2}$`)

func NewCompanyCode(code string) (CompanyCode, error) {
	if !companyCodePattern.MatchString(code) {
		return "", basic.NewFormatError(
			"CompanyCode",
			"[A-Z]{3}\\d{2}",
			code,
			"CompanyCode must be 3 uppercase letters followed by 2 digits",
		)
	}
	return CompanyCode(code), nil
}

func GenerateCompanyCode(localer repository.Localer) CompanyCode {
	letters := localer.GenerateRamdomString(3, "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	numbers := localer.GenerateRamdomString(2, "0123456789")
	return CompanyCode(letters + numbers)
}

func (c CompanyCode) String() string {
	return string(c)
}

func (c CompanyCode) Value() (driver.Value, error) {
	return driver.Value(c.String()), nil
}

func (c *CompanyCode) Scan(value interface{}) error {
	v, ok := value.(string)
	if !ok {
		return basic.NewFormatError("CompanyCode", "string", "", "invalid type for CompanyCode scan")
	}
	code, err := NewCompanyCode(v)
	if err != nil {
		return err
	}
	*c = code
	return nil
}

type CompanyCodeGetter interface {
	GetCompanyCode() ([]CompanyCode, error)
}
