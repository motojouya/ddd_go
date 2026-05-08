package store_test

import (
	"testing"

	"github.com/motojouya/ddd_go/pkg/database/core"
	util "github.com/motojouya/ddd_go/pkg/database/test"
)

var orp core.ORPer

func TestMain(m *testing.M) {
	util.ExecuteDatabaseTest("../../../", func(orpArg core.ORPer) int {
		orp = orpArg
		return m.Run()
	})
	orp = nil // il?
}
