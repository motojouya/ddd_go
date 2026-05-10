package model

import (
	"database/sql/driver"

	"github.com/google/uuid"
)

type Identifier string

func NewIdentifier(source string) (Identifier, error) {
	identifier, err := uuid.Parse(source)
	if err != nil {
		return "", err
	}
	return Identifier(identifier.String()), nil
}

func (id Identifier) String() string {
	return string(id)
}

func (id Identifier) Value() (driver.Value, error) {
	return driver.Value(string(id)), nil
}

func (id *Identifier) Scan(value interface{}) error {
	// v := string(value.([]uint8)) ?
	v, ok := value.(string)
	if !ok {
		return NewFormatError("id", "string", "", "invalid type for Identifier scan")
	}
	identifier, error := NewIdentifier(v)
	if error != nil {
		return error
	}
	*id = identifier
	return nil
}

type IdGetter interface {
	GetId() ([]Identifier, error)
}

type IdList struct {
	Ids []Identifier
}

func (obj *IdList) GetId() ([]Identifier, error) {
	return obj.Ids, nil
}

func NewIdList(ids []Identifier) IdList {
	return IdList{Ids: ids}
}

func EqId(left Identifier, right Identifier) bool {
	return left == right
}

func ToIdList[T any](sourceList []T, mapper func (T) Identifier) IdList {
	return NewIdList(Map(sourceList, mapper))
}
