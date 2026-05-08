package behavior_test

import (
	"testing"

	"github.com/motojouya/ddd_go/pkg/local/behavior"
	"github.com/stretchr/testify/assert"
)

func TestGenerateRamdomString(t *testing.T) {
	l := behavior.CreateLocal()
	randomString := l.GenerateRamdomString(10, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	assert.Len(t, randomString, 10, "Random string should be of length 10")
	assert.Regexp(t, "^[a-zA-Z0-9]+$", randomString, "Random string should only contain alphanumeric characters")
}
