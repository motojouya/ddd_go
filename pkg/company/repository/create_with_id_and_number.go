package repository

import (
	"errors"

	"github.com/motojouya/ddd_go/pkg/company/model"
	database "github.com/motojouya/ddd_go/pkg/database/model"
	local "github.com/motojouya/ddd_go/pkg/local/repository"
)

func GetMaxNum[Nd database.Keyed](db database.Executor, node Nd, company model.Company) (uint, error) {
	conditions := map[string]interface{}{
		"company_id": company.Id,
	}

	max, err := db.GetMax(node, "num", conditions)
	if err != nil {
		return 0, err
	}

	return uint(max + 1), nil
}

func CreateWithIdAndNumber[Agr any, Dpd model.Companied, Nd database.Keyed, In database.TransferableWithIdAndNumber[Agr, Dpd, Nd]](
	db database.Executor,
	lcl local.Localer,
	aggregate Agr,
	depend Dpd,
	input In,
) (Agr, error) {

	var zeroAgr Agr
	var err error

	id, err := lcl.GenerateID()
	if err != nil {
		return zeroAgr, err
	}

	var zeroNd Nd
	num, err := GetMaxNum(db, zeroNd, depend.GetCompany())
	if err != nil {
		return zeroAgr, err
	}

	newAgr, node, err := input.TransferWithIdAndNumber(aggregate, depend, id, num)
	if err != nil {
		return zeroAgr, err
	}

	keys := node.Keys()
	rowNode, err := db.Get(node, keys...)
	if err != nil {
		return zeroAgr, err
	}

	if existNode, ok := rowNode.(*Nd); ok && existNode != nil {
		// TODO errorはもう少しわかりやすいのにする
		return zeroAgr, errors.New("record already exists")
	}

	err = db.Insert(&node)
	if err != nil {
		return zeroAgr, err
	}

	return newAgr, nil
}
