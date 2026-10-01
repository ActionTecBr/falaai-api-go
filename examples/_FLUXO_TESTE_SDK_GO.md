# FLUXO DE TESTES DA SDK Go (módulo `github.com/actiontecbr/falaai-api-go`)
@version 1.0.0 | 30/09/2026 | MANUAL — nao e regenerado pelo exportador de exemplos
SDK: `D:\ProjetoFalaAI\FalaAI\FalaAI_api\sdks\go\` (modulo Go `github.com/actiontecbr/falaai-api-go`) · Docker: imagem `golang:1.25`
Regra de fundo: `.opencode/rules/macro/modulos/falaai-api/sdk-fonte-unica.md` (tests/e2e = MANUAL)

## Scripts e arquivos usados (nomes e paths exatos — LINGUAGEM: Go)
| # | Script / Arquivo | Path completo | Papel no fluxo |
|---|---|---|---|
| 1 | `run_docker.ps1` v1.0.0 | `D:\ProjetoFalaAI\FalaAI\FalaAI_api\sdks\go\tests\e2e\run_docker.ps1` | WRAPPER (PowerShell): valida go.mod + mp3 → **PASSO 0a** (sync exemplos) → Docker `golang:1.25` → **PASSO FINAL** (sync sandbox-responses) |
| 2 | `run_examples.sh` v1.0.0 | `D:\ProjetoFalaAI\FalaAI\FalaAI_api\sdks\go\tests\e2e\run_examples.sh` | RUNNER (bash, roda DENTRO do container): **limpa logs antigos**, `go run` dos 4 exemplos, grava os logs JSON + `.html` |
| 3 | `_generate_go_examples.mjs` v1.4.0 | `D:\ProjetoFalaAI\FalaAI\FalaAI_api\sdks\go\examples\_generate_go_examples.mjs` | EXPORTADOR (**PASSO 0a**, roda no HOST): re-exporta os 4 `.go` da FONTE UNICA + **gera o README** (versao/data) |
| 4 | `sandbox-examples.ts` | `D:\ProjetoFalaAI\FalaAI\FalaAI_landing\lib\sandbox-examples.ts` | FONTE UNICA dos exemplos (8 linguagens x 4 endpoints, tokens `{{...}}`) |
| 5 | `health.go` · `transcribe.go` · `diagnose.go` · `audit.go` | `D:\ProjetoFalaAI\FalaAI\FalaAI_api\sdks\go\examples\*.go` | OS 4 EXEMPLOS GERADOS (usam o SDK Go: `NewClientWithAPIKey`/`NewPublicClient` + `json.MarshalIndent`) |
| 6 | `README.md` | `D:\ProjetoFalaAI\FalaAI\FalaAI_api\sdks\go\examples\README.md` | manual dos exemplos (**GERADO** pelo exportador v1.4.0 — versao/data) |
| 7 | `demo_callcenter.mp3` | `D:\ProjetoFalaAI\FalaAI\FalaAI_api\sdks\go\tests\e2e\demo_callcenter.mp3` | audio do teste (1.3 MB) |
| 8 | `logs\go_<endpoint>_<ts>_tst.json` (+ `.html`) | `D:\ProjetoFalaAI\FalaAI\FalaAI_api\sdks\go\tests\e2e\logs\` | LOGS gerados pelo runner (**limpa os antigos a cada rodada**) |
| 9 | `VERSION.txt` | `D:\ProjetoFalaAI\FalaAI\FalaAI_api\VERSION.txt` | versao da API que o health devolve (`api_v1.21.49`) — SEM BOM |
| 10 | `_FLUXO_TESTE_SDK_GO.md` (este doc) | `D:\ProjetoFalaAI\FalaAI\FalaAI_api\sdks\go\examples\_FLUXO_TESTE_SDK_GO.md` | este doc |

## Requisitos do dev (inviolaveis)
- R-A: roda em Docker DA LINGUAGEM (`golang:1.25`)
- R-B: roda os 4 exemplos GERADOS (`health/transcribe/diagnose/audit.go`) — nunca reescreve
- R-C: LOG = request (codigo do exemplo) + response (o objeto do SDK Go) + `.html` da auditoria
- R-D: chave `fai_668e6474b83dd24540860d9ab9df0b738f7171a0f5bd8179` hardcoded no runner (local-only)
- R-E: **PASSO 0 antes de tudo** — exemplos sincronizados com a FONTE UNICA

## Fluxo (com os nomes reais dos scripts)
```mermaid
flowchart TD
    A["1. run_docker.ps1<br/>sdks/go/tests/e2e/run_docker.ps1"] --> B{"valida go.mod + demo_callcenter.mp3"}
    B --> P0["PASSO 0a (HOST): node _generate_go_examples.mjs<br/>(exemplos + README = sandbox-examples.ts)"]
    P0 --> C["2. docker run golang:1.25<br/>-e FALAAI_API_KEY -e TZ -e FALAAI_BASE_URL<br/>-v sdks/go:/go-sdk"]
    C --> D["3. run_examples.sh LIMPA logs antigos<br/>+ cp mp3 p/ o dir do modulo"]
    D --> E{"para cada endpoint<br/>health / transcribe / diagnostic / auditoria"}
    E --> F["go run ./examples/<f>.go<br/>(usa o SDK Go: NewClientWithAPIKey + json.MarshalIndent)"]
    F --> G{"stdout = JSON valido?"}
    G -- "sim" --> H["log: logs/go_<endpoint>_<ts>_tst.json<br/>{language, endpoint, generated_at,<br/>request = codigo Go, response = objeto do SDK}"]
    G -- "nao" --> I["log {response: null, error: stdout}"]
    H --> J{"endpoint = auditoria?"}
    I --> J
    J -- "sim" --> K["unzip html_report (base64+gzip)<br/>grava go_auditoria_<ts>_tst.html"]
    J -- "nao" --> E
    K --> E
    E -- "fim" --> L["TODOS OK (exit 0) / HOUVE FALHA (exit 1)"]
    L --> M["PASSO FINAL: node _sync_sandbox_responses.mjs<br/>(so em sucesso) -> lib/sandbox-responses.ts"]
```

## Comando (API no micro do dev)
```powershell
powershell -ExecutionPolicy Bypass -File "D:\ProjetoFalaAI\FalaAI\FalaAI_api\sdks\go\tests\e2e\run_docker.ps1" -BaseUrl http://host.docker.internal:8002 -Only all
# -Only: all | health | transcribe | diagnostic | auditoria
# health = publico (sem creditos) · transcribe/diagnostic/auditoria = consomem creditos da chave FAI
```

## Log (formato — secoes separadas por linha em branco, JSON valido)
```json
{
  "language": "go",
  "endpoint": "health",
  "generated_at": "2026-09-30T17:22:41Z",

  "request": "<codigo do exemplo .go, linhas escapadas como \n>",

  "response": { "o objeto do SDK Go serializado (json.MarshalIndent) — snake_case" }
}
```
Falha: `"response": null, "error": "<stdout/stderr>"` · Auditoria: + `go_auditoria_<ts>_tst.html` (unzip de `html_report`)

## O que garantir
1. O exemplo usa o SDK Go (`NewClientWithAPIKey`/`NewPublicClient` + `json.MarshalIndent`) — NUNCA chamada HTTP/curl direta.
2. O request sai do SDK (Bearer fai_...) — a API responde — o SDK devolve — o runner grava o log.
3. PASSO 0 roda ANTES de qualquer teste (exemplos na ultima versao da fonte unica).
4. O runner LIMPA os logs antigos a cada rodada (so os da execucao atual ficam).
5. PASSO FINAL: em sucesso, o wrapper sincroniza `lib/sandbox-responses.ts` (automatico).
6. Refaco a qualquer momento: mesmo comando = mesmo comportamento (deterministico).

## Historico
- **v1.0.0 (30/09 14:23):** doc inicial (padrao 3.1) — wrapper v1.0.0 (PASSO 0a + sync), runner v1.0.0 (log JSON canonico + secoes separadas + limpa logs + `.html`), exportador v1.4.0 (README GERADO); exemplos com `json.MarshalIndent`; patch do mime no `client.go` (multipart).