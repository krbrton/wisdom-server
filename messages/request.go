package messages

import "github.com/google/uuid"

type RequestMessageType uint

const (
	RequestChallenge RequestMessageType = iota
	RequestQuote
)

type RequestMessage struct {
	Id   string             `json:"id"`
	Type RequestMessageType `json:"type"`
	Body interface{}        `json:"body"`
}

func NewRequestMessage(requestType RequestMessageType, body interface{}) *RequestMessage {
	return &RequestMessage{
		Id:   uuid.New().String(),
		Type: requestType,
		Body: body,
	}
}
