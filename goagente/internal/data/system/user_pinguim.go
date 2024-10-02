package system

import (
	"fmt"
	"goagente/internal/logging"
	"os/user"
)

type LinuxUserRetriever struct{}

func (LinuxUserRetriever) GetCurrentUser() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		newErr := fmt.Errorf("erro ao obter o usuário atual no Linux: %v", err)
		logging.Error(newErr)
		return "", err
	}
	return currentUser.Username, nil
}
