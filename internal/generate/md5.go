package generate

import (
	"crypto/md5"
)

func MD5Hash(payload string) string {
	bytes := md5.Sum([]byte(payload))
	return string(bytes[:])
}
