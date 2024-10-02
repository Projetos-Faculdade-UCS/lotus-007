package system

import (
	"fmt"
	"goagente/internal/logging"
	"os/user"
	"strings"
)

type WindowsUserRetriever struct{}

func (WindowsUserRetriever) GetCurrentUser() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		newErr := fmt.Errorf("erro ao obter o usuário atual no Windows: %v", err)
		logging.Error(newErr)
		return "", err
	}
	username := currentUser.Username
	if strings.Contains(username, "\\") {
		parts := strings.Split(username, "\\")
		if len(parts) > 0 {
			return parts[len(parts)-1], nil
		}
	}
	return username, nil
}
