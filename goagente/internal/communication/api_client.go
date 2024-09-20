package communication

import (
	"bytes"
	"encoding/json"
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
	// Serializa o payload
	jsonPayload, err := serializePayload(payload)
	if err != nil {
		logging.Error(fmt.Errorf("erro ao serializar o payload: %w", err))
		return nil, err
	}

	// Cria a requisição HTTP POST
	req, err := c.createRequest(endpoint, jsonPayload)
	if err != nil {
		logging.Error(fmt.Errorf("erro ao criar a requisição: %w", err))
		return nil, err
	}

	// Envia a requisição
	resp, err := c.sendRequest(req)
	if err != nil {
		logging.Error(fmt.Errorf("erro ao enviar a requisição: %w", err))
		return nil, err
	}

	// Verifica o status da resposta
	if err := checkResponseStatus(resp); err != nil {
		logging.Error(err)
		return resp, err
	}

	return resp, nil
}

// serializePayload serializa o payload para JSON
func serializePayload(payload interface{}) ([]byte, error) {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("erro ao serializar o payload: %s", err)
	}
	return jsonPayload, nil
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
