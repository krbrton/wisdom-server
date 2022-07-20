package wisdom_server

import (
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"strings"
)

const (
	DefaultQuotePath = ".wisdom-server"
	DefaultQuoteFile = "quote.txt"
)

func GetQuote() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	quotePath := path.Join(homeDir, DefaultQuotePath, DefaultQuoteFile)
	file, err := os.Open(quotePath)
	if err != nil {
		return "", err
	}

	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	quotes := strings.Split(string(data), "\n")
	index := rand.Intn(len(quotes) - 1)

	return quotes[index], nil
}
