package messages

import (
	"crypto/sha256"
	"encoding/binary"
	"github.com/decred/dcrd/dcrec/secp256k1"
	"github.com/typticat/wisdom-server/util"
	"math/big"
	"time"
)

type Challenge struct {
	Complexity []byte `json:"complexity"`
	Timestamp  int64  `json:"timestamp"`
	Timeout    int64  `json:"timeout"`
	Entropy    []byte `json:"entropy"`
	Signature  []byte `json:"signature"`
	PublicKey  []byte `json:"public_key"`
	Solution   []byte `json:"solution"`
}

func NewChallenge(privateKeyBytes []byte, timeout int64) (*Challenge, error) {
	_, publicKey := secp256k1.PrivKeyFromBytes(privateKeyBytes)
	complexity, err := util.GetComplexity()
	if err != nil {
		return nil, err
	}

	entropy, err := util.GenerateEntropy()
	if err != nil {
		return nil, err
	}

	challenge := &Challenge{
		Complexity: complexity,
		Timestamp:  time.Now().Unix(),
		Timeout:    timeout,
		Entropy:    entropy,
		PublicKey:  publicKey.SerializeCompressed(),
	}

	return challenge, nil
}

func (c Challenge) CheckSolution() bool {
	solution := big.NewInt(0)
	solution.SetBytes(c.Solution)
	complexity := big.NewInt(0)
	complexity.SetBytes(c.Complexity)
	hash := big.NewInt(0)
	hash.SetBytes(c.Hash())
	res := hash.Cmp(complexity)

	if res == -1 {
		return true
	}

	return false
}

func (c Challenge) IsOverdue() bool {
	now := time.Now().Unix()
	deadline := c.Timestamp + c.Timeout

	return now >= deadline
}

func (c *Challenge) Solve() {
	solution := big.NewInt(0)
	solution.SetBytes(c.Solution)
	complexity := big.NewInt(0)
	complexity.SetBytes(c.Complexity)

	for {
		hash := big.NewInt(0)
		hash.SetBytes(c.Hash())
		res := hash.Cmp(complexity)

		if res == -1 {
			break
		}

		solution.Add(solution, big.NewInt(1))
		c.Solution = solution.Bytes()
	}
}

func (c Challenge) Hash() []byte {
	h := sha256.New()
	h.Write(c.Complexity)

	timestamp := make([]byte, 8)
	binary.LittleEndian.PutUint64(timestamp, uint64(c.Timestamp))
	h.Write(timestamp)

	timeout := make([]byte, 8)
	binary.LittleEndian.PutUint64(timeout, uint64(c.Timeout))
	h.Write(timeout)

	h.Write(c.Entropy)
	h.Write(c.PublicKey)
	h.Write(c.Solution)

	return h.Sum(nil)
}

func (c *Challenge) Sign(privateKeyBytes []byte) error {
	privateKey, publicKey := secp256k1.PrivKeyFromBytes(privateKeyBytes)
	sig, err := privateKey.Sign(c.Hash())
	if err != nil {
		return err
	}

	c.Signature = sig.Serialize()
	c.PublicKey = publicKey.SerializeCompressed()

	return nil
}

func (c Challenge) VerifySignature(privateKeyBytes []byte) (bool, error) {
	_, publicKey := secp256k1.PrivKeyFromBytes(privateKeyBytes)
	c.Solution = nil
	sig, err := secp256k1.ParseSignature(c.Signature)
	if err != nil {
		return false, err
	}

	return sig.Verify(c.Hash(), publicKey), nil
}
