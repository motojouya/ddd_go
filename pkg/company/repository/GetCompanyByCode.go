package repository

import (
	basic "github.com/motojouya/ddd_go/pkg/basic/model"
	"github.com/motojouya/ddd_go/pkg/company/model"
	dbModel "github.com/motojouya/ddd_go/pkg/database/model"
)

// GetCompanyByCode は CompanyCodeGetter から取得した複数の CompanyCode に一致する Company を一括取得する。
// 内部では database.Executor.GetIn を利用して `WHERE code IN (...)` で取得する。
func GetCompanyByCode(executer dbModel.Executor, codeGetter model.CompanyCodeGetter) ([]model.Company, error) {
	codes, err := codeGetter.GetCompanyCode()
	if err != nil {
		return nil, err
	}
	if len(codes) == 0 {
		return []model.Company{}, nil
	}

	conditions := map[string][]interface{}{
		"code": basic.ToInterface(basic.ToStr(codes)),
	}

	var result []model.Company
	if _, err := executer.GetIn(&result, conditions, false); err != nil {
		return nil, err
	}
	return result, nil
}
