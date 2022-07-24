package token

import (
	"context"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/saenyakorn/cleverse-intern-knowledge-sharing/pkg/contracts"
)

type TokenFetcher struct {
	ETHClient *ethclient.Client
}

func NewTokenFetcehr(client *ethclient.Client) (*TokenFetcher, error) {
	return &TokenFetcher{
		ETHClient: client,
	}, nil
}

func (tf *TokenFetcher) getTokenSymbol(ctx context.Context, addr common.Address) (string, error) {
	erc20Client, err := contracts.NewERC20Caller(addr, tf.ETHClient)
	if err != nil {
		return "", errors.WithStack(err)
	}

	symbol, err := erc20Client.Symbol(&bind.CallOpts{Context: ctx})
	if err != nil {
		return "", errors.WithStack(err)
	}

	return symbol, nil
}

func (tf *TokenFetcher) GetTokenSymbolsV1(ctx context.Context, tokenAddrs []common.Address) ([]string, error) {
	symbols := make([]string, len(tokenAddrs))
	for i, addr := range tokenAddrs {
		symbol, err := tf.getTokenSymbol(ctx, addr)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		symbols[i] = symbol
	}
	return symbols, nil
}

func (tf *TokenFetcher) GetTokenSymbolsV2(ctx context.Context, tokenAddrs []common.Address) ([]string, error) {
	var wg sync.WaitGroup
	symbols := make([]string, len(tokenAddrs))
	errChan := make(chan error, len(tokenAddrs))

	for i, addr := range tokenAddrs {
		wg.Add(1)
		go func(i int, addr common.Address) {
			defer wg.Done()
			symbol, err := tf.getTokenSymbol(ctx, addr)
			if err != nil {
				errChan <- errors.WithStack(err)
				return
			}
			symbols[i] = symbol
		}(i, addr)
	}
	wg.Wait()
	close(errChan)

	var errs error
	for err := range errChan {
		errs = errors.Wrap(errs, err.Error())
	}
	if errs != nil {
		return nil, errs
	}

	return symbols, nil
}
