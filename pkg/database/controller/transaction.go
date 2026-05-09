package controller

import (
	"github.com/motojouya/geezer_auth/pkg/shelter/user"
	dbModel "github.com/motojouya/ddd_go/pkg/database/model"
)

func RollbackWithError(database dbModel.Transactional, err error) error {
	if rollbackErr := database.Rollback(); rollbackErr != nil {
		return rollbackErr
	}
	return err
}

// control処理の頭から最後までトランザクションとする場合に有効な関数。例えば、DBアクセスもするし、APIアクセスもして、トランザクションの粒度を操作したい場合は、control処理内でbegin/commitすべき
func Transact[C dbModel.TransactionalDatabase, E any, R any](callback func(C, E, *user.Authentic) (R, error)) func(C, E, *user.Authentic) (R, error) {
	return func(control C, input E, authentic *user.Authentic) (R, error) {

		var zero R
		if err := control.Begin(); err != nil {
			return zero, err
		}

		// callback で panic が発生した場合に transaction を必ず Rollback する。
		// これがないと、orp が singleton のため次リクエストの Begin が
		// 「transaction is already started」で全 5xx になる。
		defer func() {
			if r := recover(); r != nil {
				_ = control.Rollback()
				panic(r)
			}
		}()

		result, err := callback(control, input, authentic)

		if err != nil {
			if rollbackErr := control.Rollback(); rollbackErr != nil {
				return zero, rollbackErr
			}
		} else {
			if commitErr := control.Commit(); commitErr != nil {
				return zero, commitErr
			}
		}

		return result, err
	}
}
