package entry

import (
	"github.com/motojouya/ddd_go/pkg/company/core"
)

type CompanyCode struct {
	Code string `param:"company_code" arg:"" name:"code" help:"Company Code"`
}

func (p CompanyCode) GetCompanyCode() ([]core.CompanyCode, error) {
	companyCode, err := core.NewCompanyCode(p.Code)
	if err != nil {
		return nil, err
	}

	return []core.CompanyCode{companyCode}, nil
}
