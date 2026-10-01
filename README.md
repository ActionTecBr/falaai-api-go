# falaai-api (Go) — SDK for Conversation Intelligence, Speech Analytics & Compliance

[![pkg.go.dev](https://pkg.go.dev/badge/github.com/actiontecbr/falaai-api-go.svg)](https://pkg.go.dev/github.com/actiontecbr/falaai-api-go)
[![License: MIT](https://img.shields.io/badge/license-MIT-green)](LICENSE)
[![CI](https://github.com/ActionTecBr/falaai-api-go/actions/workflows/ci.yml/badge.svg)](https://github.com/ActionTecBr/falaai-api-go/actions/workflows/ci.yml)
[![Docs](https://img.shields.io/badge/docs-GitHub%20Pages-blue)](https://actiontecbr.github.io/falaai-api-go/)

Official **Go SDK** for the **FalaAI API** — transcribe audio, analyze conversations and audit compliance (COPC CX, ISO 18295-1). **Use each API independently or combine them into your own pipeline.**

> Analyze calls, contact-center recordings, voice notes, chat and email. Get speaker-separated transcripts, summaries, reasons, actions, sentiment and a **compliance risk score**.

## Use any FalaAI API independently

FalaAI is a set of **independent REST APIs**. You **do not** need FalaAI Transcription to use FalaAI analysis or compliance auditing. If your application already has a transcript, send that text straight to the analysis APIs.

| If you have... | Use |
|---|---|
| Audio but no transcript | `CreateTranscription` |
| An existing transcript | `CreateDiagnostic` |
| A transcript needing compliance analysis | `CreateAudit` |
| An existing transcript needing both | `CreateDiagnostic` + `CreateAudit` |
| Your own STT provider (Whisper, Deepgram...) | Skip `CreateTranscription` |

```text
Your STT                             ->  FalaAI Diagnostic  ->  FalaAI Risk Audit
Telegram voice -> your STT           ->  FalaAI Risk Audit
3CX / Asterisk / Genesys transcript  ->  FalaAI Diagnostic  ->  FalaAI Risk Audit
CRM conversation                     ->  FalaAI Risk Audit
```

## Use the APIs the way you want

Every FalaAI API is **independent and optional** — chain any subset, in any combination.

```mermaid
flowchart LR
  A["Audio"] -.->|optional| T["Transcribe"]
  T --> X["Text / dialog"]
  S["Your own STT / CRM / chat / existing transcript"] --> X
  X -.->|optional| D["Diagnostic"]
  X -.->|optional| R["Risk Audit"]
  D --> O["Structured intelligence + auditable report"]
  R --> O
```

> Skip **Transcribe** if you already have text. Call only **Diagnostic**, only **Risk Audit**, or both.

## Install

```bash
go get github.com/actiontecbr/falaai-api-go
```

Requires **Go 1.24+** (module `github.com/actiontecbr/falaai-api-go`).

## Quickstart

### 1. Get an API key
Create a free account and copy your `fai_` key: <https://falaai.action.tec.br/api/auth> (or the [Dashboard](https://falaai.action.tec.br/api/dashboard)).

### 2. Set environment variables

```bash
FALAAI_BASE_URL=https://api01-falaai.action.tec.br
FALAAI_API_KEY=fai_xxxxxxxx
```

### 3. Transcribe a call (audio -> text)

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	falaai "github.com/actiontecbr/falaai-api-go"
)

func main() {
	client, err := falaai.NewClientWithAPIKey(os.Getenv("FALAAI_API_KEY"))
	if err != nil {
		panic(err)
	}

	file, err := os.Open("call.mp3")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	model := "falaai-transcribe-1"
	language := "pt"
	ref := "call_202609271408"

	r, err := client.CreateTranscription(context.Background(), falaai.TranscriptionParams{
		File:              file,
		Model:             model,
		Language:          language,
		ClientReferenceID: ref,
	})
	if err != nil {
		panic(err)
	}
	out, _ := json.MarshalIndent(r, "", "  ")
	fmt.Println(string(out))
}
```

Expected response (abridged):

```json
{
  "id": "tr-...",
  "object": "transcription",
  "model": "falaai-transcribe-1",
  "language": "por",
  "duration_seconds": 25.0,
  "text": "...",
  "dialog": "Speaker 1: [...] ...",
  "usage": { "audio_seconds": 25.0, "credits_consumed": 25, "processing_ms": 951 }
}
```

> Only need analysis? **Skip step 3** and call `CreateDiagnostic` / `CreateAudit` with your own transcript.

### 4. Analyze or audit an existing transcript (no transcription needed)

```go
transcript := "Good morning, how can I help? I need to cancel my subscription."

// 5 analyses in one call: summary, reason, action, topic, sentiment
diag, err := client.CreateDiagnostic(context.Background(), falaai.DiagnosticParams{
	Text:            &transcript,
	Language:        "pt-BR",
	DurationSeconds: 81.46,
})

// Compliance risk score + violations + auditable report
audit, err := client.CreateAudit(context.Background(), falaai.AuditParams{
	Text:             &transcript,
	Language:         "pt-BR",
	ResponseLanguage: "pt-BR",
	DurationSeconds:  81.46,
})
```

## What is FalaAI API?

FalaAI API is an **AI conversation-intelligence API** for analyzing customer-service, contact-center, sales, messaging and other business conversations. It combines speech-to-text (with speaker diarization and audio-event detection), conversation analysis (summary, contact reason, action taken, topic classification, sentiment) and a **compliance/risk audit** against **COPC CX** and **ISO 18295-1**. Conversation content is processed and discarded (zero-storage).

## What can you do with FalaAI?

- **Transcribe** audio to text with speaker separation and audio events.
- **Diagnose** a conversation: summary, reason, action taken, topic and sentiment.
- **Audit** conversations: compliance risk score, detections/violations and an auditable HTML report.
- **Track usage**, **manage webhooks** and **email alerts**, and **health/version** checks.

## Use cases

- **Contact center / Quality** — audit 100% of conversations instead of a sample.
- **Compliance / Legal** — auditable evidence for audits and disputes.
- **CX / Operations** — risk score, sentiment and reason per conversation.
- **BI / Data** — typed JSON ready for your database or analytics stack.

## SDK surface (Go)

| Symbol | Package | Purpose |
|---|---|---|
| `NewClientWithAPIKey(...)` | `falaai` | create the API client |
| `Client.Health(ctx)` | `falaai` | health check |
| `Client.CreateTranscription(ctx, TranscriptionParams{...})` | `falaai` | audio to text |
| `Client.CreateDiagnostic(ctx, DiagnosticParams{...})` | `falaai` | conversation diagnostic |
| `Client.CreateAudit(ctx, AuditParams{...})` | `falaai` | compliance risk audit |
| `participant`, `DiagnosticParticipant`, `DiagnosticAudioEvent`, `Ptr[T]` | `falaai` | models / helpers |

> Model IDs: `falaai-transcribe-1`, `falaai-diagnostic-1`, `falaai-risk-audit-1`.

## Examples

Runnable examples in [`examples/`](./examples): `health.go`, `transcribe.go`, `diagnose.go`, `audit.go`.

## Authentication

Every request requires `Authorization: Bearer fai_<your_key>` — except the public endpoints (`GET/HEAD /v1/health`, `GET /api/version`). The key is set in `NewClientWithAPIKey(...)`.

## Error handling

Methods return `(result, error)`. Check the error and handle non-2xx failures.

```go
health, err := client.Health(context.Background())
if err != nil {
	log.Fatalf("FalaAI API error: %v", err)
}
fmt.Println(health.Status)
```

## Where to integrate (this SDK)

FalaAI is language-independent; this package targets **Go** backends.

| Platform / environment (this SDK's language: **Go**) | Integration |
|---|---|
| **Mattermost** | `falaai-api-go` |
| Go microservices / workers | `falaai-api-go` |
| High-throughput pipelines / batch jobs | `falaai-api-go` |
| Other stacks (3CX, Salesforce, Genesys...) | REST / cURL — [API reference](https://api01-falaai.action.tec.br/docs) (or the SDK for that backend's language) |

> These are **integration examples**, not certified native integrations. Any platform can integrate through **REST / cURL** — see the [API reference](https://api01-falaai.action.tec.br/docs). Authenticated calls use `Authorization: Bearer fai_<key>`.

## Production usage

- Store API keys in environment variables or a secret manager — never hard-code.
- Reuse a single `*falaai.Client` across requests.
- Pass a `context.Context` for timeouts/cancellation.
- Handle returned errors explicitly.

## Compatibility

| Requirement | Version |
|---|---|
| Go | 1.24+ |
| API | v1.21.49 |

## Documentation

- **SDK docs (this language):** <https://actiontecbr.github.io/falaai-api-go/>
- **API reference (Swagger UI):** <https://api01-falaai.action.tec.br/docs>
- **OpenAPI contract:** <https://api01-falaai.action.tec.br/openapi.json>
- **Sandbox:** <https://falaai.action.tec.br/api#playground>
- **Quickstart:** <https://falaai.action.tec.br/api/quickstart>
- **Product page:** <https://falaai.action.tec.br/api>

## Versioning

Semantic versioning; tagged `v<version>` (tracks the API version `1.21.49`). See [CHANGELOG.md](CHANGELOG.md) and [Releases](https://github.com/ActionTecBr/falaai-api-go/releases).

## Security

See [SECURITY.md](SECURITY.md). Never commit real keys — use environment variables.

## Contributing

See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

[MIT](LICENSE).

## Links

- Website: <https://falaai.action.tec.br>
- API base URL: <https://api01-falaai.action.tec.br>
- GitHub organization: <https://github.com/ActionTecBr>
- Other SDKs: Python, Node.js, PHP, Ruby, Java, .NET.

### Platform documentation (orientation)

- Mattermost — <https://developers.mattermost.com>