package behavior

import (
	"github.com/go-gorp/gorp/v3"
	basic "github.com/motojouya/ddd_go/pkg/basic/core"
	"github.com/motojouya/ddd_go/pkg/database/core"
)

func Mutate[Agr any, Dpd any, Nd any, In core.TransferableList[Agr, Dpd, Nd]](
	db gorp.SqlExecutor,
	aggregate Agr,
	depend Dpd,
	input In,
) (Agr, error) {

	var zeroAgr Agr

	newAgr, mutationList, err := input.TransferList(aggregate, depend)
	if err != nil {
		return zeroAgr, err
	}

	if len(mutationList.Dels) > 0 {
		dPtrs := basic.ToPtr(mutationList.Dels)
		d := basic.ToInterface(dPtrs)
		_, err := db.Delete(d...)
		if err != nil {
			return zeroAgr, err
		}
	}

	if len(mutationList.Upds) > 0 {
		uPtrs := basic.ToPtr(mutationList.Upds)
		u := basic.ToInterface(uPtrs)
		_, err := db.Update(u...)
		if err != nil {
			return zeroAgr, err
		}
	}

	if len(mutationList.Inss) > 0 {
		iPtrs := basic.ToPtr(mutationList.Inss)
		i := basic.ToInterface(iPtrs)
		err := db.Insert(i...)
		if err != nil {
			return zeroAgr, err
		}
	}

	return newAgr, nil
}
