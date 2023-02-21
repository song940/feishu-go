package feishu

import (
	"encoding/json"
)

type Message struct {
	ReceiveId string `json:"receive_id"`
	Type      string `json:"msg_type"`
	Content   string `json:"content"`
}

type MessageResponseData struct {
}

type MessageResponse struct {
	FeishuResponseBase
	Data MessageResponseData `json:"data"`
}

func NewTextMessage(content string) (message Message) {
	message.Type = "text"
	var data, _ = json.Marshal(map[string]string{
		"text": content,
	})
	message.Content = string(data)
	return
}

// https://open.feishu.cn/document/server-docs/im-v1/message/create
func (c *Client) SendMessage(message Message) (out MessageResponse, err error) {
	data, err := c.RequestWithAccessToken("/im/v1/messages?receive_id_type=open_id", message)
	json.Unmarshal(data, &out)
	return
}
