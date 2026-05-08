package repository

import (
	"github.com/go-gorp/gorp/v3"
	"github.com/motojouya/ddd_go/pkg/database/core"
)

func Update[Agr any, Dpd any, Nd core.Keyed, In core.Transferable[Agr, Dpd, Nd]](
	db gorp.SqlExecutor,
	aggregate Agr,
	depend Dpd,
	input In,
) (Agr, error) {

	var zeroAgr Agr
	var err error

	newAgr, node, err := input.Transfer(aggregate, depend)
	if err != nil {
		return zeroAgr, err
	}

	keys := node.Keys()
	rowNode, err := db.Get(node, keys...)
	if err != nil {
		return zeroAgr, err
	}

	if existNode, ok := rowNode.(*Nd); ok && existNode != nil {
		_, err = db.Update(&node)
		if err != nil {
			return zeroAgr, err
		}

	} else {
		err = db.Insert(&node)
		if err != nil {
			return zeroAgr, err
		}
	}

	return newAgr, nil
}
