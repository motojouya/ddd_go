package repository

import (
	basic "github.com/motojouya/ddd_go/pkg/basic/core"
	"github.com/motojouya/ddd_go/pkg/company/core"
	dbCore "github.com/motojouya/ddd_go/pkg/database/core"
)

func GetCompanyById(executer dbCore.Executor, idGetter basic.IdGetter) ([]core.Company, error) {
	identifiers, err := idGetter.GetId()
	if err != nil {
		return nil, err
	}

	if len(identifiers) == 0 {
		return []core.Company{}, nil
	}

	conditions := map[string][]interface{}{
		"id": basic.ToInterface(basic.ToStr(identifiers)),
	}

	var companies []core.Company
	if _, err := executer.GetIn(&companies, conditions, false); err != nil {
		return nil, err
	}
	return companies, nil
}
