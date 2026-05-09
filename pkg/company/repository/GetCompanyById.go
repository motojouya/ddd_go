package repository

import (
	basic "github.com/motojouya/ddd_go/pkg/basic/model"
	"github.com/motojouya/ddd_go/pkg/company/model"
	dbModel "github.com/motojouya/ddd_go/pkg/database/model"
)

func GetCompanyById(executer dbModel.Executor, idGetter basic.IdGetter) ([]model.Company, error) {
	identifiers, err := idGetter.GetId()
	if err != nil {
		return nil, err
	}

	if len(identifiers) == 0 {
		return []model.Company{}, nil
	}

	conditions := map[string][]interface{}{
		"id": basic.ToInterface(basic.ToStr(identifiers)),
	}

	var companies []model.Company
	if _, err := executer.GetIn(&companies, conditions, false); err != nil {
		return nil, err
	}
	return companies, nil
}
