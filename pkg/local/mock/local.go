package mock

import (
	"github.com/google/uuid"
	"time"
)

type LocalerMock struct {
	FakeGenerateRamdomString func(length int, source string) string
	FakeGenerateUUID         func() (uuid.UUID, error)
	FakeGetNow               func() time.Time
}

func (mock LocalerMock) GenerateRamdomString(length int, source string) string {
	return mock.FakeGenerateRamdomString(length, source)
}

func (mock LocalerMock) GenerateUUID() (uuid.UUID, error) {
	return mock.FakeGenerateUUID()
}

func (mock LocalerMock) GetNow() time.Time {
	return mock.FakeGetNow()
}
