/*
 * Copyright 2018 The OpenWallet Authors
 * This file is part of the OpenWallet library.
 *
 * The OpenWallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The OpenWallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package tncd

import (
	"github.com/blocktree/futurepia-adapter/futurepia"
	"github.com/blocktree/openwallet/v2/log"
)

const (
	Symbol = "TNC"
)

type WalletManager struct {
	*futurepia.WalletManager
}

func NewWalletManager() *WalletManager {
	wm := WalletManager{}
	wm.WalletManager = futurepia.NewWalletManager()
	wm.Config = futurepia.NewConfig(Symbol)
	wm.Log = log.NewOWLogger(wm.Symbol())
	log.Warn(wm.Symbol())
	wm.TxDecoder = NewTransactionDecoder(&wm)
	wm.DecoderV2 = NewAddressDecoder2(&wm)
	return &wm
}

//Decimal 小数位精度
func (wm *WalletManager) Decimal() int32 {
	return 6
}

//FullName 币种全名
func (wm *WalletManager) FullName() string {
	return "Tncd"
}
