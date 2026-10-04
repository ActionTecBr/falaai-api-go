# sdks/go/examples - exemplos go (canonicos)

@version 1.4.0 | criado: 28/09/2026 17:22 | atualizado: 04/10/2026 02:35

## O que e
Os 4 exemplos go (go) dos endpoints da API, EXPORTADOS DA LANDING PAGE (fonte unica
de verdade). Nunca editados a mao.

## Fonte unica de verdade
- FalaAI_landing/lib/sandbox-examples.ts -> SANDBOX_EXAMPLES[<endpoint>].go
- que INTERPOLA FalaAI_landing/lib/sandbox-json-example.ts (SANDBOX_JSON_EXAMPLE = dados reais).

## Script (ferramenta)
_generate_go_examples.mjs (prefixo _ = ferramenta, nao exemplo)

| Comando                                      | Acao                                                      |
|----------------------------------------------|-----------------------------------------------------------|
| node _generate_go_examples.mjs             | Exporta (escreve os go a partir da landing + README)     |
| node _generate_go_examples.mjs --check     | Verifica (compara go x landing; exit 1 se divergir)      |

- Roda de QUALQUER pasta (resolve a landing pelo proprio caminho).
- Pre-requisito: Node (v22+).
- NUNCA edite os go nem este README a mao: edite a landing e reexecute o script.

## Ordem obrigatoria (regra)
1. Altere a FONTE (landing / SANDBOX_JSON_EXAMPLE).
2. Valide o cURL primeiro:  node _generate_go_examples.mjs --check   (exit 0 = ok)
3. So depois espelhe nas linguagens.
4. Cadeia integrada no pipeline: scripts/regenerate_all.py (T11) exporta via node.

## Relatorio da ultima execucao
| Data | Modo | Resultado |
|------|------|-----------|
| 04/10/2026 02:35 | verificacao (--check) | OK - 5 exemplos go 100% conforme a landing. |

| Arquivo | Endpoint | Status | Gerado em |
|---------|----------|--------|-----------|
| health.go | GET  /v1/health | ok | 04/10/2026 02:16 |
| transcribe.go | POST /v1/audio/transcriptions | ok | 04/10/2026 02:16 |
| diagnose.go | POST /v1/analyze/diagnostic | ok | 04/10/2026 02:16 |
| audit.go | POST /v1/analyze/riskAudit | ok | 04/10/2026 02:16 |
| whatsapp.go | whatsapp.go | ok | 04/10/2026 02:16 |

## Datas
- Criacao:     28/09/2026 17:22
- Atualizacao: 04/10/2026 02:35

Gerado automaticamente por _generate_go_examples.mjs - NAO edite a mao.
