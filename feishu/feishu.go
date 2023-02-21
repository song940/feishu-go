package feishu

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type Config struct {
	AppID       string `json:"app_id,omitempty"`
	AppSecret   string `json:"app_secret,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
}

type Client struct {
	config  *Config
	request *http.Client
}

type ResponseBase struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

type FetchResponse = []byte

func NewClient(config *Config) (client *Client, err error) {
	request := http.DefaultClient
	client = &Client{
		config,
		request,
	}
	return
}

func (c *Client) Request(path string, headers map[string]string, data interface{}) (out FetchResponse, err error) {
	payload, _ := json.Marshal(data)
	// log.Println("payload", string(payload))
	api := "https://open.feishu.cn/open-apis"
	req, err := http.NewRequest(http.MethodPost, api+path, bytes.NewBuffer(payload))
	if err != nil {
		return
	}
	req.Header.Add("content-type", "application/json; charset=utf-8")
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	res, err := c.request.Do(req)
	if err != nil {
		return
	}
	out, err = io.ReadAll(res.Body)
	return
}

func (c *Client) RequestWithAppSecret(path string) (out FetchResponse, err error) {
	return c.Request(path, nil, map[string]string{
		"app_id":     c.config.AppID,
		"app_secret": c.config.AppSecret,
	})
}

func (c *Client) SetAccessToken(accessToken string) {
	// log.Println("accessToken:", accessToken)
	c.config.AccessToken = accessToken
}

func (c *Client) RequestWithAccessToken(path string, data interface{}) (out FetchResponse, err error) {
	headers := map[string]string{
		"Authorization": "Bearer " + c.config.AccessToken,
	}
	return c.Request(path, headers, data)
}
