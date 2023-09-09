package types

import (
	"github.com/dgrijalva/jwt-go"
)

type TradeOrder struct {
	Uid             string  `json:"uid"`
	PayAmount       float64 `json:"pay_amount"`
	Number          int64   `json:"number"`
	CollectionID    string  `json:"collection_id"`
	ExchangeAssetID string  `json:"exchange_asset_id"`
	Initiator       string  `json:"initiator"` // 交易的发起者
	Recipient       string  `json:"recipient"` // 交易接受者
	jwt.StandardClaims
}

type PayOrder struct {
	Initiator    string  `json:"initiator"`     // 支付的发起者
	Recipient    string  `json:"recipient"`     // 支付接受者
	CollectionID string  `json:"collection_id"` // collection_id
	PayAmount    float64 `json:"pay_amount"`    // 支付金额
	jwt.StandardClaims
}
