package system

type LinuxOSRetriever struct{}

func (LinuxOSRetriever) GetCurrentOS() (string, error) {
	return "Linux", nil
}
