package messages

type RequestChallengeMessage struct{}

func NewRequestChallengeMessage() *RequestMessage {
	return NewRequestMessage(RequestChallenge, RequestChallengeMessage{})
}
