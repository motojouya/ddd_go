package store_test

import (
	"testing"

	"github.com/motojouya/ddd_go/pkg/database/model"
	util "github.com/motojouya/ddd_go/pkg/database/test"
)

var orp model.ORPer

func TestMain(m *testing.M) {
	util.ExecuteDatabaseTest("../../../", func(orpArg model.ORPer) int {
		orp = orpArg
		return m.Run()
	})
	orp = nil // il?
}
