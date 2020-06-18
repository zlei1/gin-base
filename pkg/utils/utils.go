package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// 生成随机数
func randomString(l int, charsets string) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = charsets[rand.Intn(len(charsets))]
	}
	return string(bytes)
}

func GenCommonCode(prefix string, num int) string {
	var charsets string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	return strings.Join([]string{
		prefix,
		strconv.FormatInt(time.Now().UnixNano(), 10),
		randomString(num, charsets),
	}, "")
}

func GenPhoneCode() string {
	var charsets string = "1234567890"
	return randomString(6, charsets)
}
