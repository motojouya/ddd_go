package repository

import (
	"errors"

	basic "github.com/motojouya/ddd_go/pkg/basic/core"
	"github.com/motojouya/ddd_go/pkg/company/core"
	database "github.com/motojouya/ddd_go/pkg/database/core"
	local "github.com/motojouya/ddd_go/pkg/local/repository"
)

func CreateRootAndMulti[Agr any, Dpd core.Companied, Nd database.Keyed, Sb any, In database.TransferableRootAndMulti[Agr, Dpd, Nd, Sb]](
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

	newAgr, root, subList, err := input.TransferRootAndMulti(aggregate, depend, id, num)
	if err != nil {
		return zeroAgr, err
	}

	keys := root.Keys()
	rowNode, err := db.Get(root, keys...)
	if err != nil {
		return zeroAgr, err
	}

	if existNode, ok := rowNode.(*Nd); ok && existNode != nil {
		// TODO errorはもう少しわかりやすいのにする
		return zeroAgr, errors.New("record already exists")
	}

	err = db.Insert(&root)
	if err != nil {
		return zeroAgr, err
	}

	if len(subList) > 0 {
		iPtrs := basic.ToPtr(subList)
		i := basic.ToInterface(iPtrs)
		err := db.Insert(i...)
		if err != nil {
			return zeroAgr, err
		}
	}

	return newAgr, nil
}
