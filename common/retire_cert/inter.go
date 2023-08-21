package retire_cert

type RetireConfigInter interface {
	Create(rc *RetireCertificate) (fileName string, savePath string, errs error)
}
