package openwtester

import (
	"path/filepath"
	"testing"

	"github.com/blocktree/openwallet/v2/log"
	"github.com/blocktree/openwallet/v2/openw"
	"github.com/blocktree/openwallet/v2/openwallet"
)

var (
	testApp        = "assets-adapter"
	configFilePath = filepath.Join("conf")
	dbFilePath     = filepath.Join("data", "db")
	dbFileName     = "blockchain-TNC.db"
)

func testInitWalletManager() *openw.WalletManager {
	log.SetLogFuncCall(true)
	tc := openw.NewConfig()

	tc.ConfigDir = configFilePath
	tc.EnableBlockScan = false
	tc.SupportAssets = []string{
		"TNC",
	}

	return openw.NewWalletManager(tc)
	//tm.Init()
}

func TestWalletManager_CreateWallet(t *testing.T) {
	tm := testInitWalletManager()
	w := &openwallet.Wallet{Alias: "TEST", IsTrust: true, Password: "1234567"}
	nw, key, err := tm.CreateWallet(testApp, w)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("wallet:", nw)
	log.Info("key:", key)

}

func TestWalletManager_GetWalletInfo(t *testing.T) {

	tm := testInitWalletManager()

	wallet, err := tm.GetWalletInfo(testApp, "Vz2oTGQwm8qLwLfKxTKPJHiunukqiCbtey")
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	log.Info("wallet:", wallet)
}

func TestWalletManager_GetWalletList(t *testing.T) {

	tm := testInitWalletManager()

	list, err := tm.GetWalletList(testApp, 0, 10000000)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Info("wallet[", i, "] :", w)
	}
	log.Info("wallet count:", len(list))

	tm.CloseDB(testApp)
}

//blocklinktest2  blocklinktest3
func TestWalletManager_CreateAssetsAccount(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "Vz2oTGQwm8qLwLfKxTKPJHiunukqiCbtey"
	//walletID = "WFn2M6aUjG7FZQKjPeArAzmddNJxcawKQh"
	//accountID="DUoAZE1eTF1ZVCDkN2mzJqaKvMQW8PsUviuX5yVzsSZG"
	account := &openwallet.AssetsAccount{Alias: "blocklinktest2", WalletID: walletID, Required: 1, Symbol: "TNC", IsTrust: true}
	account, address, err := tm.CreateAssetsAccount(testApp, walletID, "1234567", account, nil)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("account:", account)
	log.Info("address:", address)

	tm.CloseDB(testApp)
}

func TestWalletManager_GetAssetsAccountList(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "Vz2oTGQwm8qLwLfKxTKPJHiunukqiCbtey"
	list, err := tm.GetAssetsAccountList(testApp, walletID, 0, 10000000)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Info("account[", i, "] :", w)
	}
	log.Info("account count:", len(list))

	tm.CloseDB(testApp)

}

func TestWalletManager_CreateAddress(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "Vz2oTGQwm8qLwLfKxTKPJHiunukqiCbtey"
	accountID := "CbnmpvJNsUjtEMRoy5Nf5FGTyfjLbke8FuKjKtEUc7fs"
	address, err := tm.CreateAddress(testApp, walletID, accountID, 1)
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("address:", address)

	tm.CloseDB(testApp)
}

func TestWalletManager_GetAddressList(t *testing.T) {

	tm := testInitWalletManager()

	walletID := "Vz2oTGQwm8qLwLfKxTKPJHiunukqiCbtey"
	accountID := "BSVgqfUhUPm3rHr8kWjTwMAHF1gXEZZSqyYnwvuSDBHw"
	//walletID = "WFn2M6aUjG7FZQKjPeArAzmddNJxcawKQh"
	//accountID="DUoAZE1eTF1ZVCDkN2mzJqaKvMQW8PsUviuX5yVzsSZG"
	list, err := tm.GetAddressList(testApp, walletID, accountID, 0, -1, false)
	if err != nil {
		log.Error("unexpected error:", err)
		return
	}
	for i, w := range list {
		log.Info("address[", i, "] :", w.Address)
	}
	log.Info("address count:", len(list))

	tm.CloseDB(testApp)
}
