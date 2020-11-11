package bitshares

import (
	"testing"

	"github.com/blocktree/bitshares-adapter/types"
	"github.com/blocktree/openwallet/v2/log"
)

func TestWalletClient_GetBlockchainInfo(t *testing.T) {
	b, err := tw.Api.GetBlockchainInfo()
	if err != nil {
		t.Errorf("GetBlockchainInfo failed unexpected error: %v\n", err)
	} else {
		log.Infof("GetBlockchainInfo info: %+v\n", b)
	}
}

func TestWalletClient_GetBlockByHeight(t *testing.T) {
	block, err := tw.Api.GetBlockByHeight(161025)
	if err != nil {
		t.Errorf("GetBlockByHeight failed unexpected error: %v\n", err)
	} else {
		log.Infof("GetBlockByHeight info: %+v", block)
	}
}

func TestWalletClient_GetTransaction(t *testing.T) {
	tx, err := tw.Api.GetTransaction(1545399, 0)
	if err != nil {
		t.Errorf("GetTransaction failed unexpected error: %v\n", err)
	} else {
		log.Infof("GetTransaction info: %+v", tx)
	}
}

func TestWalletClient_GetAssetsBalance(t *testing.T) {
	balances, err := tw.Api.GetAssetsBalance(types.MustParseObjectID("1.2.814225"), types.MustParseObjectID("1.3.0"))
	if err != nil {
		t.Errorf("Balances failed unexpected error: %v\n", err)
	} else {
		log.Infof("Balances info: %+v", balances)
	}
}

func TestWalletClient_GetAccountID(t *testing.T) {
	id, err := tw.Api.GetAccountID("fajjoweiwew1")
	if err != nil {
		t.Errorf("AccountID failed unexpected error: %v\n", err)
	} else {
		log.Infof("AccountID info: %+v", id)
	}
}

func TestWalletClient_GetAccounts(t *testing.T) {
	id, err := tw.Api.GetAccounts("zbalice111")
	if err != nil {
		t.Errorf("get Accounts failed unexpected error: %v\n", err)
	} else {
		log.Infof("get Accounts info: %+v", id)
	}
}
