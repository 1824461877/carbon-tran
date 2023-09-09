package types

import (
	"github.com/dgrijalva/jwt-go"
)

type PayOrder struct {
	Initiator    string  `json:"initiator"`     // 支付的发起者
	Recipient    string  `json:"recipient"`     // 支付接受者
	CollectionID string  `json:"collection_id"` // collection_id
	PayAmount    float64 `json:"pay_amount"`    // 支付金额
	jwt.StandardClaims
}
