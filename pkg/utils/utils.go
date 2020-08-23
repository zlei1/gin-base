package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	tnet "github.com/toolkits/net"
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

var (
	once     sync.Once
	clientIP = "127.0.0.1"
)

func GetLocalIP() string {
	once.Do(func() {
		ips, _ := tnet.IntranetIP()
		if len(ips) > 0 {
			clientIP = ips[0]
		} else {
			clientIP = "127.0.0.1"
		}
	})
	return clientIP
}

func GenUuid() string {
	u, _ := uuid.NewRandom()
	return u.String()
}
