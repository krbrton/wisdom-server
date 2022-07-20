package messages

import (
	"crypto/rand"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const challengeTestTimeout = 2

func TestChallenge_Solve(t *testing.T) {
	challenge, err := NewChallenge(nil, challengeTestTimeout)
	assert.NoError(t, err)

	challenge.Solve()
	assert.True(t, challenge.CheckSolution())
}

func TestChallenge_IsOverdue(t *testing.T) {
	challenge, err := NewChallenge(nil, challengeTestTimeout)
	assert.NoError(t, err)

	assert.False(t, challenge.IsOverdue())
	time.Sleep(time.Second*challengeTestTimeout + 1)
	assert.True(t, challenge.IsOverdue())
}

func TestChallenge_Sign(t *testing.T) {
	privateKey := make([]byte, 32)
	_, err := rand.Read(privateKey)
	assert.NoError(t, err)

	challenge, err := NewChallenge(privateKey, challengeTestTimeout)
	assert.NoError(t, err)

	assert.Len(t, challenge.Signature, 0)

	// Ensure signature is not valid before signing
	_, err = challenge.VerifySignature(privateKey)
	assert.Error(t, err)

	// Sign challenge by server's secret key
	assert.NoError(t, challenge.Sign(privateKey))

	assert.True(t, len(challenge.Signature) > 0)
	assert.Len(t, challenge.PublicKey, 33)

	// Signature check is done properly
	ok, err := challenge.VerifySignature(privateKey)
	assert.NoError(t, err)
	assert.True(t, ok)

	// Signature remains valid after Solve()
	challenge.Solve()
	ok, err = challenge.VerifySignature(privateKey)
	assert.NoError(t, err)
	assert.True(t, ok)

	// Signature becomes invalid with changed PublicKey
	_, err = rand.Read(challenge.PublicKey)
	assert.NoError(t, err)

	ok, err = challenge.VerifySignature(privateKey)
	assert.NoError(t, err)
	assert.False(t, ok)
}
