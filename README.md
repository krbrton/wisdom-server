## Wisdom Server

### Getting started

To run wisdom server and client all you need:

```bash
cd docker
docker-compose up --build wisdom-server
# open another terminal tab to see logs from server & client
docker-compose build wisdom-client
docker-compose run --rm wisdom-client
```

Wisdom server example logs:
```
Starting docker_wisdom-server_1 ... done
Attaching to docker_wisdom-server_1
wisdom-server_1  | INFO[0000] Config loaded                                 path=/root/.wisdom-server/config.yaml
wisdom-server_1  | INFO[0000] Wisdom Server started                         address="0.0.0.0:7777"
wisdom-server_1  | INFO[0009] New request                                   address="172.21.0.3:44128" id=17a312ab-f42c-4551-bc28-a5356c4ca447 type=0
wisdom-server_1  | INFO[0009] New request                                   address="172.21.0.3:44130" id=73490f3c-2b52-47c8-a9e8-d04d125f3ee5 type=1
```

Wisom client example logs:
```
Creating docker_wisdom-client_run ... done
INFO[0000] Config loaded                                 path=/root/.wisdom-server/config.yaml
INFO[0000] Got challenge
INFO[0000] Challenge solved                              solution=16130
INFO[0000] Received quote                                quote="You don’t like something change it; if you can’t change it, change the way you think about it."
```

---

### Tests

```
go test ./... -v                                                            0 (3.573s) < 14:13:50
?   	github.com/typticat/wisdom-server	[no test files]
?   	github.com/typticat/wisdom-server/client	[no test files]
?   	github.com/typticat/wisdom-server/cmd/client	[no test files]
?   	github.com/typticat/wisdom-server/cmd/server	[no test files]
=== RUN   TestChallenge_Solve
--- PASS: TestChallenge_Solve (0.01s)
=== RUN   TestChallenge_IsOverdue
--- PASS: TestChallenge_IsOverdue (2.00s)
=== RUN   TestChallenge_Sign
--- PASS: TestChallenge_Sign (0.09s)
PASS
ok  	github.com/typticat/wisdom-server/messages	2.476s
?   	github.com/typticat/wisdom-server/server	[no test files]
=== RUN   TestGenerateEntropy
--- PASS: TestGenerateEntropy (0.00s)
PASS
ok  	github.com/typticat/wisdom-server/util	0.192s
```

### Configuration
Config's default path: `$HOME/.wisdom-server/config.yaml`

So if you want to run this locally you need to create such config. Default one can be taken from `docker` directory.

 - `host` to listen to
 - `port` to listen on
 - `timeout` for the challenge
 - `secret_key` of a server, shoule be 32-byte sized, store as a hex string

---

### Cryptography

Under the hood there are two main concepts:
1) **Proof Of Work**

Used to ensure client's solution is valid and we can send him a really wisdom quote or not. 
2) **ECDSA**

Used to ensure user didn't changed received challenge and solved it honestly.

---

### How does this works

To get a wisdom quote from the server client needs first to request challenge from the server.
This challenge should be solved by client and sent back to the server with RequestQuote.

```go
type Challenge struct {
	Complexity []byte
	Timestamp  int64
	Timeout    int64
	Entropy    []byte
	Signature  []byte
	PublicKey  []byte
	Solution   []byte
}
```

 - `Complexity` is the number which we comparing to. The smaller complexity value = higher complexity and computation time.
 - `Timestamp` points to challenge creation time.
 - `Timeout` is taken from the config and used to check challenge is overdue.
 - `Entropy` is some random value, like a nonce - used only once in each challenge.
 - `Signature` Secp256K1 signature. Signed by server's private key from the config.
 - `PublicKey` Secp256K1 public key, associated with private key.
 - `Solution` is a big number witch client should pick up and proove he solved this primitive mathematical puzzle.

### Roadmap

 - [ ] Store and check already solved and approved by server challenges so that no one could reuse challenge even during timeout
 - [ ] Make dynamic complexity. more load on server = more complexity. more user rate = more complexity.
 - [ ] Write more tests on complexity