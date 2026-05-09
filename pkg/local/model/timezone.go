package model

import (
	"os"
)

func init() {
	os.Setenv("TZ", "Asia/Tokyo")
}
