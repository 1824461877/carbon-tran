package types

type Err struct {
	Code    int    `json:"code"`
	Err     string `json:"err"`
	Message string `json:"message"`
}
