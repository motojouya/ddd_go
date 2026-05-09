package core

import (
	"github.com/go-gorp/gorp/v3"
	basic "github.com/motojouya/ddd_go/pkg/basic/core"
)

/*
 * 集約中のリスト構造をまとめてDBに反映させる
 */
func Mutate[R any](executer gorp.SqlExecutor, dels []R, upds []R, inss []R) error {
	if len(dels) > 0 {
		dPtrs := basic.ToPtr(dels)
		d := basic.ToInterface(dPtrs)
		_, err := executer.Delete(d...)
		if err != nil {
			return err
		}
	}

	if len(upds) > 0 {
		uPtrs := basic.ToPtr(upds)
		u := basic.ToInterface(uPtrs)
		_, err := executer.Update(u...)
		if err != nil {
			return err
		}
	}

	if len(inss) > 0 {
		iPtrs := basic.ToPtr(inss)
		i := basic.ToInterface(iPtrs)
		err := executer.Insert(i...)
		if err != nil {
			return err
		}
	}

	return nil
}
