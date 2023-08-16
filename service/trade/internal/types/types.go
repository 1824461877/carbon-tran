package types

import (
	"github.com/dgrijalva/jwt-go"
)

type TradeOrder struct {
	PayOrderId    string `json:"pay_order_id"`    // 支付订单id
	CarbonAssetId string `json:"carbon_asset_id"` // 用户密码
	Initiator     string `json:"initiator"`       // 交易的发起者
	Recipient     string `json:"recipient"`       // 交易接受者
	jwt.StandardClaims
}
