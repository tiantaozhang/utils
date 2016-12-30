package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5Hash md5 crypt
func MD5Hash(text string) string{
	h:= md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}
