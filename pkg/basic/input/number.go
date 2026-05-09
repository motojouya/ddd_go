package input

type Number struct {
	Num uint `param:"num"`
}

func (obj *Number) GetNumber() ([]uint, error) {
	return []uint{obj.Num}, nil
}
