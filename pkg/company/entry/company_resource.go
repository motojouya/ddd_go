package entry

import (
	basic "github.com/motojouya/ddd_go/pkg/basic/entry"
)

type CompanyResource struct {
	CompanyCode
	basic.Number
}
