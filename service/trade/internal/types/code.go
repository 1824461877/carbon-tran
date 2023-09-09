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
	Completed  int64 = 1002
	Incomplete int64 = 2002
	Expire     int64 = 1003
)
