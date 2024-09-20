package communication

import (
	"fmt"
	"goagente/internal/logging"
	"net/http"
)

type InfoPoster struct {
	client *APIClient
}

func NewInfoPoster(client *APIClient) *InfoPoster {
	return &InfoPoster{
		client: client,
	}
}

// Método genérico para enviar informações
func (p *InfoPoster) PostInfo(route string, jsonData string, infoType string) error {
	resp, err := p.client.GenericPost(route, jsonData)
	if err != nil {
		newErr := fmt.Errorf("erro ao enviar as informações de %s para o servidor: %s", infoType, err)
		logging.Error(newErr)
		fmt.Printf("Erro ao enviar as informações de %s para o servidor: %s\n", infoType, err)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Resultado JSON:", jsonData)
		fmt.Printf("Erro ao enviar as informações de %s para o servidor.\n", infoType)
		newErr := fmt.Errorf("erro ao enviar as informações de %s para o servidor, rota: %s, status: %s", infoType, route, resp.Status)
		logging.Error(newErr)
		return newErr
	}

	fmt.Println("Resposta do servidor:", resp.Status)
	fmt.Println("Resultado JSON:", jsonData)
	fmt.Printf("Informações de %s enviadas com sucesso.\n\n", infoType)
	logging.Info(fmt.Sprintf("Informações de %s enviadas com sucesso.", infoType))
	return nil
}

// Métodos específicos que chamam PostInfo
func (p *InfoPoster) PostHardwareInfo(route string, jsonData string) error {
	return p.PostInfo(route, jsonData, "hardware")
}

func (p *InfoPoster) PostCoreInfo(route string, jsonData string) error {
	return p.PostInfo(route, jsonData, "core")
}

func (p *InfoPoster) PostProgramInfo(route string, jsonData string) error {
	return p.PostInfo(route, jsonData, "programa")
}
