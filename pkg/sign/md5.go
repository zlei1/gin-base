package sign

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(str []byte) string {
	m := md5.New()
	m.Write(str)
	return hex.EncodeToString(m.Sum(nil))
}
