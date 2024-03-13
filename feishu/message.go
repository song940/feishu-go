package feishu

import (
	"encoding/json"
	"fmt"
)

type ReceiveIdType string

const (
	OPEN_ID  ReceiveIdType = "open_id"
	USER_ID  ReceiveIdType = "user_id"
	UNION_ID ReceiveIdType = "union_id"
	EMAIL    ReceiveIdType = "email"
	CHAT_ID  ReceiveIdType = "chat_id"
)

type Message struct {
	ReceiveIdType ReceiveIdType
	ReceiveId     string `json:"receive_id"`
	Type          string `json:"msg_type"`
	Content       string `json:"content"`
	UUID          string `json:"uuid,omitempty"`
}

type MessageResponseData struct {
}

type MessageResponse struct {
	ResponseBase
	Data MessageResponseData `json:"data"`
}

func NewTextMessage(content string) (message Message) {
	var data, _ = json.Marshal(map[string]string{
		"text": content,
	})
	message.Type = "text"
	message.Content = string(data)
	return
}

// https://open.feishu.cn/document/server-docs/im-v1/message/create
func (c *Client) SendMessage(message *Message) (out *MessageResponse, err error) {
	if message.ReceiveIdType == "" {
		message.ReceiveIdType = OPEN_ID
	}
	path := fmt.Sprintf("/im/v1/messages?receive_id_type=%s", message.ReceiveIdType)
	data, err := c.RequestWithAccessToken(path, message)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &out)
	return
}

func (c *Client) SendTextMessage(receiveId string, content string) (out *MessageResponse, err error) {
	message := NewTextMessage(content)
	message.ReceiveId = receiveId
	return c.SendMessage(&message)
}
