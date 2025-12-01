package core

import (
	"os"
)

func init() {
	os.Setenv("TZ", "Asia/Tokyo")
}
