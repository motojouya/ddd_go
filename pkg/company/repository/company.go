package repository

import (
	basic "github.com/motojouya/ddd_go/pkg/basic/core"
	"github.com/motojouya/ddd_go/pkg/company/core"
	dbCore "github.com/motojouya/ddd_go/pkg/database/core"
)

type Company interface {
	GetCompanyById(idGetter basic.IdGetter) ([]core.Company, error)
	GetCompanyByCode(codeGetter core.CompanyCodeGetter) ([]core.Company, error)
}

type CompanyRepository struct {
	Executer dbCore.Executor
}

func NewCompanyRepository(executer dbCore.Executor) *CompanyRepository {
	return &CompanyRepository{
		Executer: executer,
	}
}

func (b *CompanyRepository) GetCompanyById(idGetter basic.IdGetter) ([]core.Company, error) {
	return GetCompanyById(b.Executer, idGetter)
}

func (b *CompanyRepository) GetCompanyByCode(codeGetter core.CompanyCodeGetter) ([]core.Company, error) {
	return GetCompanyByCode(b.Executer, codeGetter)
}
