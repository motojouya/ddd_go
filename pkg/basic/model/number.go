package model

type NumberGetter interface {
	GetNumber() ([]uint, error)
}

type NumberList struct {
	NumList []uint
}

func (obj *NumberList) GetNumber() ([]uint, error) {
	return obj.NumList, nil
}

func NewNumberList(numberList []uint) NumberList {
	return NumberList{NumList: numberList}
}

func ToNumberList[T any](sourceList []T, mapper func (T) uint) NumberList {
	return NewNumberList(Map(sourceList, mapper))
}
