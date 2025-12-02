package core

import (
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

type IdGetter interface {
  GetId() ([]string, error)
}
