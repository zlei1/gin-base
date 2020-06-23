package dingding

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"time"
)

const dingHost = "oapi.dingtalk.com"

var dingUrl url.URL = url.URL{
	Scheme: "https",
	Host:   dingHost,
	Path:   "robot/send",
}

func GetDingUrl(accessToken string, secret string) (string, error) {
	dtu := dingUrl
	value := url.Values{}
	value.Set("access_token", accessToken)

	if secret == "" {
		dtu.RawQuery = value.Encode()
		return dtu.String(), nil
	}

	timestamp := strconv.FormatInt(time.Now().Unix()*1000, 10)
	sign, err := sign(timestamp, secret)
	if err != nil {
		dtu.RawQuery = value.Encode()
		return dtu.String(), err
	}

	value.Set("timestamp", timestamp)
	value.Set("sign", sign)
	dtu.RawQuery = value.Encode()
	return dtu.String(), nil
}

func sign(timestamp string, secret string) (string, error) {
	stringToSign := fmt.Sprintf("%s\n%s", timestamp, secret)
	h := hmac.New(sha256.New, []byte(secret))
	if _, err := io.WriteString(h, stringToSign); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}
