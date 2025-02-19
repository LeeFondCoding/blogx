package util

import (
	"crypto/md5"
	"encoding/hex"
)

func InList[T comparable](key T, list []T) bool {
	for _, s := range list {
		if key == s {
			return true
		}
	}
	return false
}

func MD5(data []byte) string {
	md5Hash := md5.New()
	md5Hash.Write(data)
	return hex.EncodeToString(md5Hash.Sum(nil))
}