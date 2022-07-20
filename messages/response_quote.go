package messages

type ResponseQuoteMessage struct {
	Quote string `json:"quote"`
}

func NewResponseQuoteMessage(quote string) *ResponseQuoteMessage {
	return &ResponseQuoteMessage{Quote: quote}
}
