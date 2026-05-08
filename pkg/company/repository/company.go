package behavior

import (
	basic "github.com/motojouya/ddd_go/pkg/basic/core"
	"github.com/motojouya/ddd_go/pkg/company/core"
	dbCore "github.com/motojouya/ddd_go/pkg/database/core"
)

type Company interface {
	GetCompanyById(idGetter basic.IdGetter) ([]core.Company, error)
	GetCompanyByCode(codeGetter core.CompanyCodeGetter) ([]core.Company, error)
}

type CompanyBehavior struct {
	Executer dbCore.Executor
}

func NewCompanyBehavior(executer dbCore.Executor) *CompanyBehavior {
	return &CompanyBehavior{
		Executer: executer,
	}
}

func (b *CompanyBehavior) GetCompanyById(idGetter basic.IdGetter) ([]core.Company, error) {
	return GetCompanyById(b.Executer, idGetter)
}

func (b *CompanyBehavior) GetCompanyByCode(codeGetter core.CompanyCodeGetter) ([]core.Company, error) {
	return GetCompanyByCode(b.Executer, codeGetter)
}
