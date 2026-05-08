package repository

import (
	"errors"

	"github.com/go-gorp/gorp/v3"
	"github.com/motojouya/ddd_go/pkg/database/core"
	local "github.com/motojouya/ddd_go/pkg/local/repository"
)

func CreateWithId[Agr any, Dpd any, Nd core.Keyed, In core.TransferableWithId[Agr, Dpd, Nd]](
	db gorp.SqlExecutor,
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

	newAgr, node, err := input.TransferWithId(aggregate, depend, id)
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
