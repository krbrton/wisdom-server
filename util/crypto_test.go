package util

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateEntropy(t *testing.T) {
	entropy1, err := GenerateEntropy()
	assert.NoError(t, err)

	entropy2, err := GenerateEntropy()
	assert.NoError(t, err)

	res := bytes.Compare(entropy1, entropy2)
	assert.Equal(t, 1, res)
}
