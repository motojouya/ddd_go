package repository

import (
	"time"

	basicModel "github.com/motojouya/ddd_go/pkg/basic/model"
)

type LocalerMock struct {
	FakeGenerateRamdomString func(length int, source string) string
	FakeGenerateID           func() (basicModel.Identifier, error)
	FakeGetNow               func() time.Time
}

func (mock LocalerMock) GenerateRamdomString(length int, source string) string {
	return mock.FakeGenerateRamdomString(length, source)
}

func (mock LocalerMock) GenerateID() (basicModel.Identifier, error) {
	return mock.FakeGenerateID()
}

func (mock LocalerMock) GetNow() time.Time {
	return mock.FakeGetNow()
}
