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
