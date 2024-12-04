// Package communication fornece uma interface para comunicação com uma API HTTP.
package communication

import (
	"bytes"
	"fmt"
	"goagente/internal/logging"
	"net/http"
	"time"
)

// APIClient representa um cliente para comunicação com uma API HTTP.
// Ele inclui uma URL base e um cliente HTTP configurado com timeout.
type APIClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewAPIClient cria uma nova instância do APIClient com a URL base fornecida.
// O cliente HTTP é configurado com um timeout padrão de 10 segundos.
//
// Parâmetros:
// - baseURL: A URL base para todas as requisições.
//
// Retorna:
// - Uma instância de APIClient.
func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second, // Timeout padrão para evitar bloqueios
		},
	}
}

// GenericPost envia uma requisição POST genérica para o endpoint especificado.
// O payload pode ser fornecido como uma string ou []byte.
//
// Parâmetros:
// - endpoint: O endpoint para onde a requisição será enviada.
// - payload: Os dados a serem enviados no corpo da requisição (string ou []byte).
//
// Retorna:
// - Uma resposta HTTP, se a requisição for bem-sucedida.
// - Um erro, se ocorrer algum problema ao preparar, enviar ou processar a resposta.
func (c *APIClient) GenericPost(endpoint string, payload interface{}) (*http.Response, error) {
	// Converte o payload para []byte, se necessário
	jsonPayload, err := preparePayload(payload)
	if err != nil {
		logging.Error(fmt.Errorf("erro ao preparar o payload: %w", err))
		return nil, err
	}

	// Cria a requisição HTTP POST
	req, err := c.createRequest(endpoint, jsonPayload)
	if err != nil {
		logging.Error(fmt.Errorf("erro ao criar a requisição: %w", err))
		return nil, err
	}
	fmt.Println(string(jsonPayload)) // Exibe o payload no console para depuração

	// Envia a requisição HTTP
	resp, err := c.sendRequest(req)
	if err != nil {
		logging.Error(fmt.Errorf("erro ao enviar a requisição: %w", err))
		return nil, err
	}

	// Verifica o código de status da resposta
	if err := checkResponseStatus(resp); err != nil {
		logging.Error(err)
		return resp, err
	}

	return resp, nil
}

// preparePayload converte o payload fornecido para o tipo []byte, se necessário.
//
// Parâmetros:
// - payload: Os dados a serem enviados (pode ser []byte ou string).
//
// Retorna:
// - O payload convertido em []byte.
// - Um erro, se o tipo do payload não for suportado.
func preparePayload(payload interface{}) ([]byte, error) {
	switch p := payload.(type) {
	case []byte:
		return p, nil // Payload já está no formato []byte
	case string:
		return []byte(p), nil // Converte string para []byte
	default:
		return nil, fmt.Errorf("tipo de payload não suportado: %T", payload)
	}
}

// createRequest cria uma requisição HTTP POST com o payload fornecido.
//
// Parâmetros:
// - endpoint: O endpoint para onde a requisição será enviada.
// - jsonPayload: Os dados a serem enviados no corpo da requisição.
//
// Retorna:
// - Uma instância de http.Request, se bem-sucedida.
// - Um erro, se ocorrer algum problema ao criar a requisição.
func (c *APIClient) createRequest(endpoint string, jsonPayload []byte) (*http.Request, error) {
	url := fmt.Sprintf("%s/%s", c.BaseURL, endpoint)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("erro ao criar a requisição POST para %s: %s", url, err)
	}
	req.Header.Set("Content-Type", "application/json") // Define o Content-Type como JSON
	return req, nil
}

// sendRequest envia uma requisição HTTP e retorna a resposta.
//
// Parâmetros:
// - req: A requisição HTTP a ser enviada.
//
// Retorna:
// - Uma instância de http.Response, se bem-sucedida.
// - Um erro, se ocorrer algum problema ao enviar a requisição.
func (c *APIClient) sendRequest(req *http.Request) (*http.Response, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao enviar a requisição para %s: %s", req.URL.String(), err)
	}
	return resp, nil
}

// checkResponseStatus verifica se o código de status da resposta HTTP indica sucesso.
//
// Parâmetros:
// - resp: A resposta HTTP a ser verificada.
//
// Retorna:
// - Um erro, se o código de status for menor que 200 ou maior ou igual a 300.
// - Nil, se o código de status for válido.
func checkResponseStatus(resp *http.Response) error {
	if resp.StatusCode >= 399 {
		return fmt.Errorf("resposta HTTP com erro. Status: %d, URL: %s", resp.StatusCode, resp.Request.URL.String())
	}
	return nil
}
