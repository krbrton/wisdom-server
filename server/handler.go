package server

import (
	"encoding/base64"
	"encoding/json"
	"github.com/sirupsen/logrus"
	wisdom_server "github.com/typticat/wisdom-server"
	"github.com/typticat/wisdom-server/messages"
	"net"
)

func (s Server) HandleConnection(conn net.Conn) {
	var (
		responseBody        interface{}
		responseError       string
		responseMessageType messages.ResponseMessageType
	)

	defer conn.Close()

	requestMsg, err := s.readMessage(conn)
	if err != nil {
		logrus.Error(err)
		return
	}

	logrus.WithField("address", conn.RemoteAddr()).
		WithField("id", requestMsg.Id).
		WithField("type", requestMsg.Type).
		Info("New request")

	privateKey, err := base64.StdEncoding.DecodeString(s.config.SecretKey)
	if err != nil {
		logrus.Error(err)
		return
	}

	switch requestMsg.Type {
	case messages.RequestChallenge:
		newChallenge, err := messages.NewChallenge(privateKey, s.config.Timeout)
		if err != nil {
			logrus.Error(err)
			return
		}

		if err = newChallenge.Sign(privateKey); err != nil {
			logrus.Error(err)
			return
		}

		responseMessageType = messages.ResponseChallenge
		responseBody = messages.NewResponseChallengeMessage(newChallenge)

	case messages.RequestQuote:
		responseMessageType = messages.ResponseQuote
		responseBody = messages.NewResponseQuoteMessage("")

		reqBodyJson, err := json.Marshal(requestMsg.Body)
		if err != nil {
			logrus.Error(err)
			return
		}

		var body messages.RequestQuoteMessage

		if err = json.Unmarshal(reqBodyJson, &body); err != nil {
			logrus.Error(err)
			return
		}

		ok, err := body.Challenge.VerifySignature(privateKey)
		if err != nil {
			responseError = err.Error()
			goto sendResponse
		}

		if !ok {
			responseError = "invalid signature"
			goto sendResponse
		}

		if !body.Challenge.CheckSolution() {
			responseError = "invalid solution"
			goto sendResponse
		}

		if body.Challenge.IsOverdue() {
			responseError = "challenge deadline exceeded"
			goto sendResponse
		}

		quote, err := wisdom_server.GetQuote()
		if err != nil {
			responseError = err.Error()
			goto sendResponse
		}

		responseBody = messages.NewResponseQuoteMessage(quote)

	default:
		panic("not implemented")
	}

sendResponse:
	responseMsg := messages.NewResponseMessage(requestMsg.Id, responseMessageType, responseBody)

	if responseError != "" {
		responseMsg.Error = responseError
	}

	if err = s.writeMessage(conn, responseMsg); err != nil {
		logrus.Error(err)
	}
}
