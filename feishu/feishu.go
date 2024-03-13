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
	*Config
	*http.Client
}

type ResponseBase struct {
	Code int    `json:"code,omitempty"`
	Msg  string `json:"msg,omitempty"`
}

func NewClient(config *Config) *Client {
	return &Client{
		Config: config,
		Client: http.DefaultClient,
	}
}

func (c *Client) Request(path string, headers map[string]string, data interface{}) (out []byte, err error) {
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
	res, err := c.Client.Do(req)
	if err != nil {
		return
	}
	out, err = io.ReadAll(res.Body)
	return
}

func (c *Client) RequestWithAppSecret(path string) (out []byte, err error) {
	return c.Request(path, nil, map[string]string{
		"app_id":     c.Config.AppID,
		"app_secret": c.Config.AppSecret,
	})
}

func (c *Client) SetAccessToken(accessToken string) {
	// log.Println("accessToken:", accessToken)
	c.Config.AccessToken = accessToken
}

func (c *Client) RequestWithAccessToken(path string, data interface{}) (out []byte, err error) {
	headers := map[string]string{
		"Authorization": "Bearer " + c.Config.AccessToken,
	}
	return c.Request(path, headers, data)
}
