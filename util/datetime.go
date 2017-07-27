package util

import (
	"time"
)

func NowRFC3339() string {
	return time.Now().Format(time.RFC3339)
}
