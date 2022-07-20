package main

import (
	"github.com/sirupsen/logrus"
	wisdom_server "github.com/typticat/wisdom-server"
	"github.com/typticat/wisdom-server/server"
)

func main() {
	cfg, err := wisdom_server.ReadConfig()
	if err != nil {
		panic(err)
	}

	srv := server.NewServer(cfg)

	if err := srv.Run(); err != nil {
		logrus.Fatalln(err)
	}
}
