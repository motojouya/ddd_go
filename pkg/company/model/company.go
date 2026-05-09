package core

import (
	"github.com/doug-martin/goqu/v9"
	"github.com/go-gorp/gorp/v3"
	basic "github.com/motojouya/ddd_go/pkg/basic/core"
)

const CompanyTable = "company"
const CompanyAlias = "c"

var CompanySelect = goqu.Select(
	goqu.I(CompanyAlias+".id"),
	goqu.I(CompanyAlias+".code"),
	goqu.I(CompanyAlias+".name"),
).From(goqu.T(CompanyTable).As(CompanyAlias))

type Company struct {
	Id   basic.Identifier `db:"id" json:"_"`
	Code CompanyCode      `db:"code" json:"code"`
	Name string           `db:"name" json:"name"`
}

func AddCompanyTable(dbMap *gorp.DbMap) {
	dbMap.AddTableWithName(Company{}, CompanyTable).SetKeys(false, "Id")
}

func NewCompany(id basic.Identifier, code CompanyCode, name string) Company {
	return Company{
		Id:   id,
		Code: code,
		Name: name,
	}
}
