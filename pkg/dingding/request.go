package dingding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const httpTimoutSecond = time.Duration(30) * time.Second

type Client struct {
	AccessToken string
	Secret      string
}

type Response struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode int64  `json:"errcode"`
}

func (c *Client) Send(msg Message) (Response, error) {
	res := Response{}

	reqBytes, err := msg.ToJsonByte()
	if err != nil {
		return res, err
	}

	dingUrl, err := GetDingUrl(c.AccessToken, c.Secret)
	if err != nil {
		return res, err
	}

	req, err := http.NewRequest("POST", dingUrl, bytes.NewReader(reqBytes))
	if err != nil {
		return res, err
	}
	req.Header.Add("Accept-Charset", "utf8")
	req.Header.Add("Content-Type", "application/json")

	client := new(http.Client)
	client.Timeout = httpTimoutSecond
	resp, err := client.Do(req)
	if err != nil {
		return res, err
	}

	result, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return res, err
	}

	err = json.Unmarshal(result, &res)
	if err != nil {
		return res, fmt.Errorf("unmarshal http response body from json error = %v", err)
	}

	if res.ErrCode != 0 {
		return res, fmt.Errorf("send message to dingtalk error = %s", res.ErrMsg)
	}

	return res, nil
}
