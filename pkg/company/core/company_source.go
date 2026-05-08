package core

import (
	"errors"

	basic "github.com/motojouya/ddd_go/pkg/basic/core"
)

type SourceType string

const (
	SourceTypeCompany        SourceType = "COMPANY"
	SourceTypeImage          SourceType = "IMAGE"
	SourceTypeItem           SourceType = "ITEM"
)

var sourceTypes = map[SourceType]bool{
	SourceTypeCompany:        true,
	SourceTypeImage:          true,
	SourceTypeItem:           true,
}

func GetSourceType(code SourceType) error {
	if !sourceTypes[code] {
		return errors.New("invalid source type: " + string(code))
	}
	return nil
}

type CompanySource struct {
	SourceType  SourceType       `db:"source_type"`
	Id          basic.Identifier `db:"id"`
	CompanyCode CompanyCode      `db:"-"`
	Num         uint             `db:"-"`
}

type CompanySourceGetter interface {
	GetCompanySource() CompanySource
}

func NewCompanySource(sourceType SourceType, id basic.Identifier, companyCode CompanyCode, num uint) (CompanySource, error) {
	if err := GetSourceType(sourceType); err != nil {
		return CompanySource{}, basic.NewInvalidArgumentError("sourceType", string(sourceType), "invalid source type")
	}
	return CompanySource{
		SourceType:  sourceType,
		Id:          id,
		CompanyCode: companyCode,
		Num:         num,
	}, nil
}
