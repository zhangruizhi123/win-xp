package common

import (
	"crypto/md5"
	"encoding/hex"
)

//获取字符串md5值
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
