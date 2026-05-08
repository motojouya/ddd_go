package behavior

import (
	database "github.com/motojouya/ddd_go/pkg/database/core"
)

func GetOrCreate[T database.Keyed](executor database.Executor, obj T) (T, error) {
	var zeroObj T

	keys := obj.Keys()
	rowObj, err := executor.Get(obj, keys...)
	if err != nil {
		return zeroObj, err
	}

	existObj := rowObj.(*T)
	if existObj == nil {
		err = executor.Insert(&obj)
		if err != nil {
			return zeroObj, err
		}
	}

	return obj, nil
}
