package messages

type ResponseChallengeMessage struct {
	Challenge Challenge `json:"challenge"`
}

func NewResponseChallengeMessage(challenge *Challenge) *ResponseChallengeMessage {
	return &ResponseChallengeMessage{
		Challenge: *challenge,
	}
}
