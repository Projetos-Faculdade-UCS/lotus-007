package patrimonio

type LinuxPatRetriever struct{}

func (LinuxPatRetriever) GetCurrentPat() (string, error) {
	return "11111", nil
}
