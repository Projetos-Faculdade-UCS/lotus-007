package communication

import (
	"bytes"
	"fmt"
	"goagente/internal/logging"
	"net/http"
	"time"
)

type APIClient struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewAPIClient cria uma nova instância do cliente API com suporte a HTTPS e timeout configurado
func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second, // Timeout padrão para evitar bloqueios
		},
	}
}

// GenericPost envia uma requisição POST genérica usando HTTPS
func (c *APIClient) GenericPost(endpoint string, payload interface{}) (*http.Response, error) {
	// Verifica o tipo do payload e converte para []byte, se necessário
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
	fmt.Println(string(jsonPayload)) // Para verificar o payload no console

	// Envia a requisição
	resp, err := c.sendRequest(req)
	if err != nil {
		logging.Error(fmt.Errorf("erro ao enviar a requisição: %w", err))
		return nil, err
	}

	if err := checkResponseStatus(resp); err != nil {
		logging.Error(err)
		return resp, err
	}

	return resp, nil
}

// preparePayload converte o payload para []byte, se necessário
func preparePayload(payload interface{}) ([]byte, error) {
	switch p := payload.(type) {
	case []byte:
		return p, nil // Payload já é []byte
	case string:
		return []byte(p), nil // Converte string para []byte
	default:
		return nil, fmt.Errorf("tipo de payload não suportado: %T", payload)
	}
}

// createRequest cria a requisição HTTP POST
func (c *APIClient) createRequest(endpoint string, jsonPayload []byte) (*http.Request, error) {
	url := fmt.Sprintf("%s/%s", c.BaseURL, endpoint)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return nil, fmt.Errorf("erro ao criar a requisição POST para %s: %s", url, err)
	}
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

// sendRequest envia a requisição HTTP
func (c *APIClient) sendRequest(req *http.Request) (*http.Response, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao enviar a requisição para %s: %s", req.URL.String(), err)
	}
	return resp, nil
}

// checkResponseStatus verifica se o código de status da resposta HTTP é válido
func checkResponseStatus(resp *http.Response) error {
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("resposta HTTP com erro. Status: %d, URL: %s", resp.StatusCode, resp.Request.URL.String())
	}
	return nil
}
