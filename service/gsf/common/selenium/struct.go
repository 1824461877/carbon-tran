package selenium

type GsfToken struct {
	Data interface{} `json:"data"`
}

type Token struct {
	AccessToken string `json:"AccessToken"`
	IdToken     string `json:"idToken"`
}
