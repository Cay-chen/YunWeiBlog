package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(s string) string {
	//"7BN#SDQa!+_%123"
	s = s + "7BN#SDQa!+_%123"
	m := md5.New()
	m.Write([]byte(s))
	return hex.EncodeToString(m.Sum(nil))
}
