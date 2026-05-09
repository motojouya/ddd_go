package repository

import (
	"math/rand"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/google/uuid"
	basicModel "github.com/motojouya/ddd_go/pkg/basic/model"
)

type Localer interface {
	GenerateRamdomString(length int, source string) string
	GenerateID() (basicModel.Identifier, error)
	GetNow() time.Time
}

type Local struct{}

func CreateLocal() *Local {
	return &Local{}
}

func (l Local) GenerateRamdomString(length int, source string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = source[rand.Intn(len(source))]
	}
	return string(b)
}

func (l Local) GenerateID() (basicModel.Identifier, error) {
	uuidValue, err := uuid.NewV7()
	if err != nil {
		return basicModel.Identifier(""), err
	}

	id, err := basicModel.NewIdentifier(uuidValue.String())
	if err != nil {
		return basicModel.Identifier(""), err
	}

	return id, nil
}

func (l Local) GetNow() time.Time {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	return time.Now().In(jst)
}

func GetEnv[T any]() (T, error) {
	return env.ParseAs[T]()
}
