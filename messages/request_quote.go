package messages

type RequestQuoteMessage struct {
	Challenge Challenge `json:"challenge"`
}

func NewRequestQuoteMessage(challenge *Challenge) *RequestQuoteMessage {
	return &RequestQuoteMessage{
		Challenge: *challenge,
	}
}
