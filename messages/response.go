package messages

import (
	"time"
)

type ResponseMessageType uint

const (
	ResponseChallenge ResponseMessageType = iota
	ResponseQuote
)

type ResponseMessage struct {
	Id        string              `json:"id"`
	Type      ResponseMessageType `json:"type"`
	Error     string              `json:"error"`
	Timestamp int64               `json:"timestamp"`
	Body      interface{}         `json:"body"`
}

func NewResponseMessage(Id string, responseType ResponseMessageType, body interface{}) *ResponseMessage {
	responseMsg := &ResponseMessage{
		Id:        Id,
		Type:      responseType,
		Timestamp: time.Now().Unix(),
		Body:      body,
	}

	return responseMsg
}
