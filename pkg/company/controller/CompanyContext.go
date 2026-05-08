package controller

import (
	"github.com/motojouya/geezer_auth/pkg/shelter/user"
	basic "github.com/motojouya/ddd_go/pkg/basic/core"
	"github.com/motojouya/ddd_go/pkg/company/repository"
	"github.com/motojouya/ddd_go/pkg/company/core"
)

func CompanyContext[C repository.Company, E core.CompanyCodeGetter, R any](callback func(C, E, *user.Authentic, core.Company) (R, error)) func(C, E, *user.Authentic) (R, error) {
	return func(control C, entry E, authentic *user.Authentic) (R, error) {

		companies, err := control.GetCompanyByCode(entry)
		if err != nil {
			var zero R
			return zero, err
		}

		if len(companies) == 0 {
			var zero R
			return zero, basic.NewNotFoundError("Company", map[string]string{}, "company not found")
		}

		// FIXME ここで認可制御を入れる。

		return callback(control, entry, authentic, companies[0])
	}
}
