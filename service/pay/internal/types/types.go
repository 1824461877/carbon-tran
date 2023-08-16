package types

import (
	"github.com/dgrijalva/jwt-go"
)

type PayOrder struct {
	Initiator         string  `json:"initiator"` // 支付的发起者
	Recipient         string  `json:"recipient"` // 支付接受者
	InitiatorWalletId string  `json:"initiator_wallet_id"`
	RecipientWalletId string  `json:"recipient_wallet_id"`
	PayAmount         float64 `json:"pay_amount"` // 支付金额
	jwt.StandardClaims
}
