package retire_cert

import "github.com/fogleman/gg"

type RetireCertificate struct {
	RetiredBy      string
	ProjectName    string
	ByFrom         string
	ByTo           string
	OnTime         string
	VerifiedNumber int64
}

type RetireConfig struct {
	Path            string
	ImageBackground string
	FontTTF         string
}

type RetireC struct {
	retireConfig RetireConfig
	x            float64
	dc           *gg.Context
}
