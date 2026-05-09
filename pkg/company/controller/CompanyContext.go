package controller

import (
	"github.com/motojouya/geezer_auth/pkg/shelter/user"
	basic "github.com/motojouya/ddd_go/pkg/basic/model"
	"github.com/motojouya/ddd_go/pkg/company/repository"
	"github.com/motojouya/ddd_go/pkg/company/model"
)

func CompanyContext[C repository.Company, E model.CompanyCodeGetter, R any](callback func(C, E, *user.Authentic, model.Company) (R, error)) func(C, E, *user.Authentic) (R, error) {
	return func(control C, input E, authentic *user.Authentic) (R, error) {

		companies, err := control.GetCompanyByCode(input)
		if err != nil {
			var zero R
			return zero, err
		}

		if len(companies) == 0 {
			var zero R
			return zero, basic.NewNotFoundError("Company", map[string]string{}, "company not found")
		}

		// FIXME ここで認可制御を入れる。

		return callback(control, input, authentic, companies[0])
	}
}
