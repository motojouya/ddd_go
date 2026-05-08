package repository

import (
	"github.com/go-gorp/gorp/v3"
	"github.com/motojouya/ddd_go/pkg/database/core"
)

func Delete[Agr any, Dpd any, Nd core.Keyed, In core.Transferable[Agr, Dpd, Nd]](
	db gorp.SqlExecutor,
	aggregate Agr,
	depend Dpd,
	input In,
) (Agr, bool, error) {

	var zeroAgr Agr
	var err error

	newAgr, node, err := input.Transfer(aggregate, depend)
	if err != nil {
		return zeroAgr, false, err
	}

	keys := node.Keys()
	rowNode, err := db.Get(node, keys...)
	if err != nil {
		return zeroAgr, false, err
	}

	if existNode, ok := rowNode.(*Nd); !ok || existNode == nil {
		return zeroAgr, false, nil
	}

	_, err = db.Delete(&node)
	if err != nil {
		return zeroAgr, false, err
	}

	return newAgr, true, nil
}
