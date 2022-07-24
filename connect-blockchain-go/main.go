package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/saenyakorn/cleverse-intern-knowledge-sharing/pkg/fibo"
	"github.com/saenyakorn/cleverse-intern-knowledge-sharing/pkg/token"
)

func main() {

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	nodeUrl := "https://rpc.ankr.com/bsc"

	client, err := ethclient.DialContext(ctx, nodeUrl)
	if err != nil {
		log.Panic("Can not connect to the Ethereum client", err)
	}

	tf, err := token.NewTokenFetcehr(client)
	if err != nil {
		log.Panic("Can not create token fetcher", err)
	}

	var (
		WBNDAddr = common.HexToAddress("0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c")
		BUSDAddr = common.HexToAddress("0xe9e7cea3dedca5984780bafc599bd69add087d56")
	)

	var amount = 2
	var mockAddrs []common.Address
	for i := 1; i < amount; i++ {
		mockAddrs = append(mockAddrs, WBNDAddr)
		mockAddrs = append(mockAddrs, BUSDAddr)
	}

	// Version 1
	start := time.Now()
	symbols, err := tf.GetTokenSymbolsV1(ctx, mockAddrs)
	if err != nil {
		log.Panic("Can not get token symbol", err)
	}
	elapsed := time.Since(start)
	fmt.Println("Version 1, Took", elapsed, ", Symbols:", symbols)

	// Version 2
	start = time.Now()
	symbols, err = tf.GetTokenSymbolsV2(ctx, mockAddrs)
	if err != nil {
		log.Panic("Can not get token symbol", err)
	}
	elapsed = time.Since(start)
	fmt.Println("Version 2, Took", elapsed, ", Symbols:", symbols)

	// Fibonacci
	resultChan := make(chan int)
	go func() {
		resultChan <- fibo.Fibonacci(ctx, 42)
	}()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context Done")
			return
		case result := <-resultChan:
			fmt.Println("Fibonacci:", result)
			return
		default:
			fmt.Println("Sleep for 1 second")
			time.Sleep(time.Second)
		}
	}
}
