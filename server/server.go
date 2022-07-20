package server

import (
	"bufio"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	wisdom_server "github.com/typticat/wisdom-server"
	"github.com/typticat/wisdom-server/messages"
	"net"
)

type Server struct {
	config *wisdom_server.Config
}

func NewServer(config *wisdom_server.Config) *Server {
	return &Server{
		config: config,
	}
}

func (s Server) readMessage(conn net.Conn) (*messages.RequestMessage, error) {
	var (
		requestMsgBase64 string
		requestMsgData   []byte
		requestMsg       messages.RequestMessage
	)

	reader := bufio.NewReader(conn)
	requestMsgBase64, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	requestMsgData, err = base64.StdEncoding.DecodeString(requestMsgBase64)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(requestMsgData, &requestMsg); err != nil {
		return nil, err
	}

	return &requestMsg, nil
}

func (s Server) writeMessage(conn net.Conn, responseMsg *messages.ResponseMessage) error {
	responseMsgData, err := json.Marshal(responseMsg)
	if err != nil {
		return err
	}

	responseMsgBase64 := base64.StdEncoding.EncodeToString(responseMsgData)
	_, err = conn.Write([]byte(responseMsgBase64 + "\n"))

	return err
}

func (s Server) Run() error {
	address := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	ln, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}

	defer ln.Close()

	logrus.WithField("address", address).Info("Wisdom Server started")

	for {
		conn, err := ln.Accept()
		if err != nil {
			logrus.Error(err)
			continue
		}

		go s.HandleConnection(conn)
	}
}
