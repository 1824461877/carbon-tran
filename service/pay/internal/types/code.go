package types

const (
	SuccessCode             int = 200
	ValidErrorCode          int = 400
	InvalidRequestErrorCode int = 401
	InternalErrorCode       int = 500

	UserNotRegistered int = 9992
	RemoteLogin       int = 9993
	TokenErrorCode    int = 99997
)

const (
	Unpaid int32 = 1001
	Paid   int32 = 1002
	Expire int32 = 1003
)
