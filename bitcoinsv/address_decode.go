/*
 * Copyright 2018 The openwallet Authors
 * This file is part of the openwallet library.
 *
 * The openwallet library is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The openwallet library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Lesser General Public License for more details.
 */

package bitcoinsv

import (
	"github.com/blocktree/go-owcdrivers/addressEncoder"
	"github.com/blocktree/go-owcrypt"
)

func init() {

}

//var (
//	AddressDecoder = &openwallet.AddressDecoder{
//		PrivateKeyToWIF:    PrivateKeyToWIF,
//		PublicKeyToAddress: PublicKeyToAddress,
//		WIFToPrivateKey:    WIFToPrivateKey,
//	}
//)

type addressDecoder struct {
	wm *WalletManager //钱包管理者
}

//NewAddressDecoder 地址解析器
func NewAddressDecoder(wm *WalletManager) *addressDecoder {
	decoder := addressDecoder{}
	decoder.wm = wm
	return &decoder
}

//PrivateKeyToWIF 私钥转WIF
func (decoder *addressDecoder) PrivateKeyToWIF(priv []byte, isTestnet bool) (string, error) {

	cfg := addressEncoder.BTC_mainnetPrivateWIFCompressed
	if decoder.wm.Config.IsTestNet {
		cfg = addressEncoder.BTC_testnetPrivateWIFCompressed
	}

	//privateKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), priv)
	//wif, err := btcutil.NewWIF(privateKey, &cfg, true)
	//if err != nil {
	//	return "", err
	//}

	wif := addressEncoder.AddressEncode(priv, cfg)

	return wif, nil

}

//PublicKeyToAddress 公钥转地址
func (decoder *addressDecoder) PublicKeyToAddress(pub []byte, isTestnet bool) (string, error) {

	cfg := addressEncoder.BTC_mainnetAddressP2PKH

	if decoder.wm.Config.IsTestNet {
		cfg = addressEncoder.BTC_testnetAddressP2PKH
	}

	//pkHash := btcutil.Hash160(pub)
	//address, err :=  btcutil.NewAddressPubKeyHash(pkHash, &cfg)
	//if err != nil {
	//	return "", err
	//}

	pkHash := owcrypt.Hash(pub, 0, owcrypt.HASH_ALG_HASH160)

	address := addressEncoder.AddressEncode(pkHash, cfg)

	return address, nil

}

//RedeemScriptToAddress 多重签名赎回脚本转地址
func (decoder *addressDecoder) RedeemScriptToAddress(pubs [][]byte, required uint64, isTestnet bool) (string, error) {

	cfg := addressEncoder.BTC_mainnetAddressP2SH
	if decoder.wm.Config.IsTestNet {
		cfg = addressEncoder.BTC_testnetAddressP2SH
	}

	redeemScript := make([]byte, 0)

	for _, pub := range pubs {
		redeemScript = append(redeemScript, pub...)
	}

	pkHash := owcrypt.Hash(redeemScript, 0, owcrypt.HASH_ALG_HASH160)

	address := addressEncoder.AddressEncode(pkHash, cfg)

	return address, nil

}

//WIFToPrivateKey WIF转私钥
func (decoder *addressDecoder) WIFToPrivateKey(wif string, isTestnet bool) ([]byte, error) {

	cfg := addressEncoder.BTC_mainnetPrivateWIFCompressed
	if decoder.wm.Config.IsTestNet {
		cfg = addressEncoder.BTC_testnetPrivateWIFCompressed
	}

	priv, err := addressEncoder.AddressDecode(wif, cfg)
	if err != nil {
		return nil, err
	}

	return priv, err

}

//ScriptPubKeyToBech32Address scriptPubKey转Bech32地址
func (decoder *addressDecoder) ScriptPubKeyToBech32Address(scriptPubKey []byte) (string, error) {
	return "", nil

}