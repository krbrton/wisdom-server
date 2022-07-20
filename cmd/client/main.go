package main

import (
	"github.com/sirupsen/logrus"
	wisdom_server "github.com/typticat/wisdom-server"
	"github.com/typticat/wisdom-server/client"
	"math/big"
)

func main() {
	cfg, err := wisdom_server.ReadConfig()
	if err != nil {
		panic(err)
	}

	cli := client.NewClient(cfg)
	challenge, err := cli.RequestChallenge()
	if err != nil {
		logrus.Panic(err)
		return
	}

	logrus.Info("Got challenge")

	challenge.Solve()
	solution := big.NewInt(0)
	solution.SetBytes(challenge.Solution)

	logrus.
		WithField("solution", solution.String()).
		Info("Challenge solved")

	quote, err := cli.RequestQuote(challenge)
	if err != nil {
		logrus.Panic(err)
		return
	}

	logrus.
		WithField("quote", quote).
		Info("Received quote")
}
