package core

import (
	"slices"

	"github.com/go-gorp/gorp/v3"
	basic "github.com/motojouya/ddd_go/pkg/basic/core"
)

type Keyed interface {
	Keys() []interface{}
}

type Transferable[Agr any, Dpd any, Nd Keyed] interface {
	Transfer(Agr, Dpd) (Agr, Nd, error)
}

type TransferableWithId[Agr any, Dpd any, Nd Keyed] interface {
	TransferWithId(Agr, Dpd, basic.Identifier) (Agr, Nd, error)
}

type TransferableWithIdAndNumber[Agr any, Dpd any, Nd Keyed] interface {
	TransferWithIdAndNumber(Agr, Dpd, basic.Identifier, uint) (Agr, Nd, error)
}

type TransferableRootAndSub[Agr any, Dpd any, Rt Keyed, Sb any] interface {
	TransferRootAndSub(Agr, Dpd, basic.Identifier, uint) (Agr, Rt, Sb, error)
}

type TransferableRootAndMulti[Agr any, Dpd any, Rt Keyed, Sb any] interface {
	TransferRootAndMulti(Agr, Dpd, basic.Identifier, uint) (Agr, Rt, []Sb, error)
}

type TransferableList[Agr any, Dpd any, Nd any] interface {
	TransferList(Agr, Dpd) (Agr, MutationList[Nd], error)
}

type Transfer[Agr any, Dpd any, Nd Keyed, In any] func(Agr, Dpd, In) (Agr, Nd, error)

func EmptyTransfer[Agr any, Dpd any, Nd Keyed, In any](aggregate Agr, _ Dpd, _ In) (Agr, Nd, error) {
	var zeroNode Nd
	return aggregate, zeroNode, nil
}

type TransferWithId[Agr any, Dpd any, Nd Keyed, In any] func(Agr, Dpd, In, basic.Identifier) (Agr, Nd, error)

type TransferWithIdAndNumber[Agr any, Dpd any, Nd Keyed, In any] func(Agr, Dpd, In, basic.Identifier, uint) (Agr, Nd, error)

type GenerateNumber[Dpd any] func(gorp.SqlExecutor, Dpd) (uint, error)

type MutationList[Nd any] struct {
	Dels []Nd
	Upds []Nd
	Inss []Nd
}

func (ml MutationList[Nd]) Merge(additional MutationList[Nd]) MutationList[Nd] {
	return MutationList[Nd]{
		Dels: slices.Concat(ml.Dels, additional.Dels),
		Upds: slices.Concat(ml.Upds, additional.Upds),
		Inss: slices.Concat(ml.Inss, additional.Inss),
	}
}

type Replace[Agr any, Dpd any, Nd any, In any] func(Agr, Dpd, In) (Agr, MutationList[Nd], error)

func EmptyReplace[Agr any, Dpd any, Nd any, In any](aggregate Agr, _ Dpd, _ In) (Agr, MutationList[Nd], error) {
	var zeroMutationList MutationList[Nd]
	return aggregate, zeroMutationList, nil
}

type Order struct {
	Column    string
	Ascending bool
}

func NewOrder(column string, ascending bool) Order {
	return Order{
		Column:    column,
		Ascending: ascending,
	}
}
