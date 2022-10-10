package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"

	"triptych.labs/nexus"
	"triptych.labs/nexus/nex"
	"triptych.labs/nexus/nex/ops"
	"triptych.labs/utils"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func airdrop(publicKey string) {

	url := "https://api.devnet.solana.com"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`
  {"jsonrpc":"2.0","id":1, "method":"requestAirdrop", "params":["%s", 1000000000]}
`, publicKey))

	fmt.Println(payload)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func init() {
	nexus.SetProgramID(solana.MustPublicKeyFromBase58("5GxQd4ocs9Xs3osbaPqzHLRFaKRZHv8UNDzBQoDQyr9h"))
}

func main() {
	op := os.Args[1]
	switch op {
	case "fuzz":
		{
			Fuzz()
		}
	case "report":
		{
			Report()
		}
	case "moniker":
		{
			Moniker()
		}
	}
}

func Fuzz() {
	var wg sync.WaitGroup

	accounts := []solana.PrivateKey{
		solana.NewWallet().PrivateKey,
	}

	monikers := []string{
		RandStringRunes(10),
	}

	_monikers := []string{
		RandStringRunes(10),
	}

	log.Println(utils.NETWORK)
	// client := rpc.New("https://api.devnet.solana.com")

	log.Println("Airdropping...")
	for _, wallet := range accounts {
		func(_wallet solana.PrivateKey) {
			airdrop(_wallet.PublicKey().String())
			time.Sleep(10 * time.Second)
		}(wallet)
	}

	time.Sleep(5 * time.Second)

	log.Println("Onboarding...")
	for _, wallet := range accounts {
		go func(_wallet solana.PrivateKey) {
			wg.Add(1)
			instructions := make([]solana.Instruction, 0)

			onboardIx := ops.OnboardNex(_wallet.PublicKey())
			instructions = append(instructions, onboardIx)

			utils.SendTx(
				"create",
				instructions,
				append(make([]solana.PrivateKey, 0), _wallet),
				_wallet.PublicKey(),
			)
			wg.Done()
		}(wallet)
	}

	time.Sleep(5 * time.Second)
	wg.Wait()

	log.Println("Registering and Assign...")
	for i, wallet := range accounts {
		func(_i int, _wallet solana.PrivateKey) {
			wg.Add(1)
			instructions := make([]solana.Instruction, 0)

			registerIx := ops.RegisterMoniker(_wallet.PublicKey(), monikers[_i])
			instructions = append(instructions, registerIx)

			_registerIx := ops.RegisterMoniker(_wallet.PublicKey(), _monikers[_i])
			instructions = append(instructions, _registerIx)

			_assignIx := ops.AssignMoniker(_wallet.PublicKey(), _monikers[_i])
			instructions = append(instructions, _assignIx)

			utils.SendTx(
				"create",
				instructions,
				append(make([]solana.PrivateKey, 0), _wallet),
				_wallet.PublicKey(),
			)
			wg.Done()
		}(i, wallet)
	}

	wg.Wait()
}

func Report() {
	rpcClient := rpc.New(utils.NETWORK)
	nextumer := solana.MustPublicKeyFromBase58(os.Args[2])

  monikers := nex.GetMonikersData(rpcClient, nextumer)
  monikersJson, _ := json.MarshalIndent(monikers, "", "  ")

  fmt.Println(string(monikersJson))
}

func Moniker() {
  rpcClient := rpc.New(utils.NETWORK)
	monikerName := os.Args[2]

  nextumer := ops.GetNextumerFromMoniker(rpcClient, monikerName)

	log.Println(nextumer)
}
