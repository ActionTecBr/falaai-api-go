// falaai.go - Fachada tipada (hand-written) sobre o client gerado (openapi-generator).
//
// @version 3.0.0 | 2026-09-30
//
// NAO e gerado: sobrevive a regeneracao.
// NUNCA edite os arquivos gerados (client.go/api_*.go/model_*.go) - edite ESTE arquivo.
//
// Objetivo: a SDK recebe o DOMINIO (params tipados), nunca o TRANSPORTE (contentType/io.Reader).
package falaai

import (
	"context"
	"fmt"
	"os"
)

// DefaultBaseURL - endpoint publico de producao da API FalaAI.
const DefaultBaseURL = "https://api01-falaai.action.tec.br"

// NewClientWithAPIKey - construtor de conveniencia: resolve a base URL (env FALAAI_BASE_URL,
// caindo em DefaultBaseURL), valida a api key e injeta o header Authorization.
func NewClientWithAPIKey(apiKey string) (*APIClient, error) {
	if apiKey == "" {
		return nil, fmt.Errorf("falaai: apiKey e obrigatoria")
	}
	baseURL := os.Getenv("FALAAI_BASE_URL")
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}
	return NewClient(baseURL, apiKey)
}

// NewClient - client sobre uma base URL explicita (ex.: sandbox local) + api key.
func NewClient(baseURL, apiKey string) (*APIClient, error) {
	if baseURL == "" {
		return nil, fmt.Errorf("falaai: baseURL e obrigatoria")
	}
	if apiKey == "" {
		return nil, fmt.Errorf("falaai: apiKey e obrigatoria")
	}
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: baseURL, Description: "FalaAI API"}}
	cfg.AddDefaultHeader("Authorization", "Bearer "+apiKey)
	return NewAPIClient(cfg), nil
}

// NewPublicClient - client para as rotas publicas (ex.: /v1/health), sem api key.
func NewPublicClient(baseURL string) *APIClient {
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{{URL: baseURL, Description: "FalaAI API"}}
	return NewAPIClient(cfg)
}

// Health - GET /v1/health (sem body).
func (c *APIClient) Health(ctx context.Context) (*HealthResponse, error) {
	resp, _, err := c.HealthAPI.HealthCheck(ctx).Execute()
	return resp, err
}

// HealthCheckHead - HEAD /v1/health; devolve o status HTTP.
func (c *APIClient) HealthCheckHead(ctx context.Context) (int, error) {
	_, httpResp, err := c.HealthAPI.HealthCheckHead(ctx).Execute()
	if httpResp == nil {
		return 0, err
	}
	return httpResp.StatusCode, err
}

// TranscriptionParams - parametros da transcricao (multipart/form-data).
type TranscriptionParams struct {
	File              *os.File // obrigatorio - audio (o nome do arquivo vai no formulario)
	Model             string   // obrigatorio - ex.: falaai-transcribe-1
	Language          string   // obrigatorio - ex.: pt
	ClientReferenceID string   // opcional - correlacao
}

// CreateTranscription - POST /v1/audio/transcriptions (o multipart e montado pelo gerado).
func (c *APIClient) CreateTranscription(ctx context.Context, p TranscriptionParams) (*TranscriptionResponse, error) {
	if p.File == nil {
		return nil, fmt.Errorf("falaai: TranscriptionParams.File e obrigatorio")
	}
	if p.Model == "" {
		return nil, fmt.Errorf("falaai: TranscriptionParams.Model e obrigatorio")
	}
	if p.Language == "" {
		return nil, fmt.Errorf("falaai: TranscriptionParams.Language e obrigatorio")
	}
	req := c.SpeechAPI.CreateTranscriptionV1AudioTranscriptionsPost(ctx).
		File(p.File).Model(p.Model).Language(p.Language)
	if p.ClientReferenceID != "" {
		req = req.ClientReferenceId(p.ClientReferenceID)
	}
	resp, _, err := req.Execute()
	return resp, err
}

// DiagnosticParams - parametros do diagnostico (JSON). Alias do body tipado gerado:
// obrigatorios por valor, opcionais por ponteiro.
type DiagnosticParams = DiagnosticRequest

// CreateDiagnostic - POST /v1/analyze/diagnostic.
func (c *APIClient) CreateDiagnostic(ctx context.Context, p DiagnosticParams) (*DiagnosticResponse, error) {
	resp, _, err := c.AnalysisAPI.CreateDiagnosticV1AnalyzeDiagnosticPost(ctx).DiagnosticRequest(p).Execute()
	return resp, err
}

// AuditParams - parametros da auditoria de risco (JSON). Alias do body tipado gerado.
type AuditParams = RiskAuditRequest

// CreateAudit - POST /v1/analyze/riskAudit.
func (c *APIClient) CreateAudit(ctx context.Context, p AuditParams) (*RiskAuditV2Response, error) {
	resp, _, err := c.AnalysisAPI.CreateRiskAuditV1AnalyzeRiskAuditPost(ctx).RiskAuditRequest(p).Execute()
	return resp, err
}

// Ptr - retorna o ponteiro de v (campos opcionais ponteiro): falaai.Ptr("Carla"), falaai.Ptr(float32(3.1)).
func Ptr[T any](v T) *T { return &v }
