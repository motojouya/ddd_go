package repository

import (
	basic "github.com/motojouya/ddd_go/pkg/basic/model"
	"github.com/motojouya/ddd_go/pkg/company/model"
	dbModel "github.com/motojouya/ddd_go/pkg/database/model"
)

type Company interface {
	GetCompanyById(idGetter basic.IdGetter) ([]model.Company, error)
	GetCompanyByCode(codeGetter model.CompanyCodeGetter) ([]model.Company, error)
}

type CompanyRepository struct {
	Executer dbModel.Executor
}

func NewCompanyRepository(executer dbModel.Executor) *CompanyRepository {
	return &CompanyRepository{
		Executer: executer,
	}
}

func (b *CompanyRepository) GetCompanyById(idGetter basic.IdGetter) ([]model.Company, error) {
	return GetCompanyById(b.Executer, idGetter)
}

func (b *CompanyRepository) GetCompanyByCode(codeGetter model.CompanyCodeGetter) ([]model.Company, error) {
	return GetCompanyByCode(b.Executer, codeGetter)
}
