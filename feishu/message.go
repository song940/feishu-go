package feishu

import (
	"encoding/json"
)

type FeishuMessage struct {
	ReceiveId string `json:"receive_id"`
	Type      string `json:"msg_type"`
	Content   string `json:"content"`
}

type FeishuMessageResponseData struct {
}

type FeishuMessageResponse struct {
	FeishuResponseBase
	Data FeishuMessageResponseData `json:"data"`
}

func NewTextMessage(content string) (message FeishuMessage) {
	message.Type = "text"
	var data, _ = json.Marshal(map[string]string{
		"text": content,
	})
	message.Content = string(data)
	return
}

// https://open.feishu.cn/document/server-docs/im-v1/message/create

func (c *FeishuClient) SendMessage(message FeishuMessage) (out FeishuMessageResponse, err error) {
	data, err := c.RequestWithAccessToken("/im/v1/messages?receive_id_type=open_id", message)
	json.Unmarshal(data, &out)
	return
}
