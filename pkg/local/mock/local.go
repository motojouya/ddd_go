package mock

import (
	"time"

	basicCore "github.com/motojouya/ddd_go/pkg/basic/core"
)

type LocalerMock struct {
	FakeGenerateRamdomString func(length int, source string) string
	FakeGenerateID           func() (basicCore.Identifier, error)
	FakeGetNow               func() time.Time
}

func (mock LocalerMock) GenerateRamdomString(length int, source string) string {
	return mock.FakeGenerateRamdomString(length, source)
}

func (mock LocalerMock) GenerateID() (basicCore.Identifier, error) {
	return mock.FakeGenerateID()
}

func (mock LocalerMock) GetNow() time.Time {
	return mock.FakeGetNow()
}
