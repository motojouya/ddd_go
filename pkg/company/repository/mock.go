package repository

import (
	basic "github.com/motojouya/ddd_go/pkg/basic/model"
	"github.com/motojouya/ddd_go/pkg/company/model"
)

type CompanyMock struct {
	GetCompanyByIdFunc   func(idGetter basic.IdGetter) ([]model.Company, error)
	GetCompanyByCodeFunc func(codeGetter model.CompanyCodeGetter) ([]model.Company, error)
}

func NewCompanyMock() *CompanyMock {
	return &CompanyMock{
		GetCompanyByIdFunc: func(idGetter basic.IdGetter) ([]model.Company, error) {
			return []model.Company{}, nil
		},
		GetCompanyByCodeFunc: func(codeGetter model.CompanyCodeGetter) ([]model.Company, error) {
			return []model.Company{}, nil
		},
	}
}

func (mock *CompanyMock) GetCompanyById(idGetter basic.IdGetter) ([]model.Company, error) {
	return mock.GetCompanyByIdFunc(idGetter)
}

func (mock *CompanyMock) GetCompanyByCode(codeGetter model.CompanyCodeGetter) ([]model.Company, error) {
	return mock.GetCompanyByCodeFunc(codeGetter)
}
