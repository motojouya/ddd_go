package behavior

import (
	"github.com/google/uuid"
	"math/rand"
	"time"
	"github.com/caarlos0/env/v11"
)

type Localer interface {
	GenerateRamdomString(length int, source string) string
	GenerateUUID() (uuid.UUID, error)
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

func (l Local) GenerateUUID() (uuid.UUID, error) {
	var uuidValue, err = uuid.NewV7()
	if err != nil {
		var zero = uuid.UUID{}
		return zero, err
	}

	return uuidValue, nil
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

