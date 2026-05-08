package repository

import (
	basic "github.com/motojouya/ddd_go/pkg/basic/core"
	"github.com/motojouya/ddd_go/pkg/company/core"
)

type CompanyMock struct {
	GetCompanyByIdFunc   func(idGetter basic.IdGetter) ([]core.Company, error)
	GetCompanyByCodeFunc func(codeGetter core.CompanyCodeGetter) ([]core.Company, error)
}

func NewCompanyMock() *CompanyMock {
	return &CompanyMock{
		GetCompanyByIdFunc: func(idGetter basic.IdGetter) ([]core.Company, error) {
			return []core.Company{}, nil
		},
		GetCompanyByCodeFunc: func(codeGetter core.CompanyCodeGetter) ([]core.Company, error) {
			return []core.Company{}, nil
		},
	}
}

func (mock *CompanyMock) GetCompanyById(idGetter basic.IdGetter) ([]core.Company, error) {
	return mock.GetCompanyByIdFunc(idGetter)
}

func (mock *CompanyMock) GetCompanyByCode(codeGetter core.CompanyCodeGetter) ([]core.Company, error) {
	return mock.GetCompanyByCodeFunc(codeGetter)
}
