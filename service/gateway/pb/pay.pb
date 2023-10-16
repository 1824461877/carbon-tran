
ù
pb/proto/pay.protopb"@
PayReq
	TaskToken (	R	TaskToken
ReqTime (	RReqTime",

PayOnceReq

PayOrderId (	R
PayOrderId"Î
PayOnceResp

PayOrderId (	R
PayOrderId
	Initiator (	R	Initiator
	Recipient (	R	Recipient
	PayStatus (R	PayStatus
	PayAmount (R	PayAmount$
InitiatorTime (	RInitiatorTime

FinishTime (	R
FinishTime"@
PayListResp1
PayListResp (2.pb.PayOnceRespRPayListResp"p
PayExecutionResp

PayOrderId (	R
PayOrderId
	PayStatus (R	PayStatus

CreateTime (	R
CreateTime"/
PayApproveReq

PayOrderId (	R
PayOrderId".
PayApproveResp
	PayStatus (	R	PayStatus"3
PayOrderStatusReq

PayOrderId (	R
PayOrderId"h
PayOrderStatusResp
Code (RCode
TTL (RTTL
PayID (	RPayID
Status (	RStatus"A
PayAllListResp/

PayAllList (2.pb.PayOnceListR
PayAllList"Å
PayOnceList

PayOrderId (	R
PayOrderId
PayId (	RPayId
	Initiator (	R	Initiator
	Recipient (	R	Recipient
	PayStatus (R	PayStatus
	PayAmount (R	PayAmount$
InitiatorTime (	RInitiatorTime

FinishTime (	R
FinishTime2ñ
	PayServer2
GetBasicPayOnce.pb.PayOnceReq.pb.PayOnceList-

GetPayList.pb.PayOnceReq.pb.PayListResp0
PayExecution
.pb.PayReq.pb.PayExecutionResp?
PayOrderStatus.pb.PayOrderStatusReq.pb.PayOrderStatusResp3

PayApprove.pb.PayApproveReq.pb.PayApproveRespB	Z./protobproto3