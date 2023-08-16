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
	Completed  int32 = 2001
	Incomplete int32 = 2002
)
