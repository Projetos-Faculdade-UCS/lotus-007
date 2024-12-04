// Package communication lida com a comunicação externa via API.
package communication

import (
	"fmt"
	"goagente/internal/logging"
	"net/http"
)

// InfoPoster é responsável por enviar informações usando um APIClient.
type InfoPoster struct {
	client *APIClient
}

// NewInfoPoster cria uma nova instância de InfoPoster com o cliente fornecido.
func NewInfoPoster(client *APIClient) *InfoPoster {
	return &InfoPoster{
		client: client,
	}
}

// PostInfo envia informações genéricas para o servidor.
//
// Parâmetros:
// - route: rota para onde enviar as informações.
// - jsonData: dados em formato JSON a serem enviados.
// - infoType: tipo de informação sendo enviada (ex: "hardware", "core", "programa").
//
// Retorna:
// - erro: se ocorrer um erro durante o envio das informações.
func (p *InfoPoster) PostInfo(route string, jsonData string, infoType string) error {
	resp, err := p.client.GenericPost(route, jsonData)
	if err != nil {
		newErr := fmt.Errorf("erro ao enviar as informações de %s para o servidor: %s", infoType, err)
		logging.Error(newErr)
		fmt.Printf("Erro ao enviar as informações de %s para o servidor: %s\n", infoType, err)
		return err
	}

	if resp.StatusCode >= 400 && resp.StatusCode <= 599 {
		if resp.StatusCode == http.StatusBadRequest || resp.StatusCode >= http.StatusInternalServerError {
			fmt.Println("Resultado JSON:", jsonData)
			fmt.Printf("Erro ao enviar as informações de %s para o servidor.\n", infoType)
			newErr := fmt.Errorf("erro ao enviar as informações de %s para o servidor, rota: %s, status: %s", infoType, route, resp.Status)
			logging.Error(newErr)
			return newErr
		}
	}

	fmt.Println("Resposta do servidor:", resp.Status)
	fmt.Println("Resultado JSON:", jsonData)
	fmt.Printf("Informações de %s enviadas com sucesso.\n\n", infoType)
	logging.Info(fmt.Sprintf("Informações de %s enviadas com sucesso.", infoType))
	return nil
}

// PostHardwareInfo envia informações de hardware para o servidor.
//
// Parâmetros:
// - route: rota para onde enviar as informações.
// - jsonData: dados em formato JSON a serem enviados.
//
// Retorna:
// - erro: se ocorrer um erro durante o envio das informações.
func (p *InfoPoster) PostHardwareInfo(route string, jsonData string) error {
	return p.PostInfo(route, jsonData, "hardware")
}

// PostCoreInfo envia informações do core para o servidor.
//
// Parâmetros:
// - route: rota para onde enviar as informações.
// - jsonData: dados em formato JSON a serem enviados.
//
// Retorna:
// - erro: se ocorrer um erro durante o envio das informações.
func (p *InfoPoster) PostCoreInfo(route string, jsonData string) error {
	return p.PostInfo(route, jsonData, "core")
}

// PostProgramInfo envia informações de programas para o servidor.
//
// Parâmetros:
// - route: rota para onde enviar as informações.
// - jsonData: dados em formato JSON a serem enviados.
//
// Retorna:
// - erro: se ocorrer um erro durante o envio das informações.
func (p *InfoPoster) PostProgramInfo(route string, jsonData string) error {
	return p.PostInfo(route, jsonData, "programa")
}
