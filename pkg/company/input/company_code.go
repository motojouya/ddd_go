package input

import (
	"github.com/motojouya/ddd_go/pkg/company/model"
)

type CompanyCode struct {
	Code string `param:"company_code" arg:"" name:"code" help:"Company Code"`
}

func (p CompanyCode) GetCompanyCode() ([]model.CompanyCode, error) {
	companyCode, err := model.NewCompanyCode(p.Code)
	if err != nil {
		return nil, err
	}

	return []model.CompanyCode{companyCode}, nil
}
