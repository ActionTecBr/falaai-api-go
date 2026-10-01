package e2e

import (
	"context"
	"fmt"
	"os"
	"testing"

	falaai "github.com/actiontecbr/falaai-api-go"
)

func TestAiChain(t *testing.T) {
	c, _ := NewClient(ProdURL(), TestKey())
	ctx := context.Background()

	f, err := os.Open(AudioPath())
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	tr, httpResp, err := c.SpeechAPI.CreateTranscriptionV1AudioTranscriptionsPost(ctx).
		File(f).Model("falaai-transcribe-1").Language("pt").
		ClientReferenceId("e2e-call-2026-09-22-001").Execute()
	if err != nil {
		t.Fatal(err)
	}
	status := httpResp.StatusCode
	Log("transcriptions", "POST", "/v1/audio/transcriptions",
		map[string]string{"file": "analise_25s.mp3", "model": "falaai-transcribe-1", "language": "pt", "client_reference_id": "e2e-call-2026-09-22-001"},
		tr, fmt.Sprintf("HTTP %d", status), status)
	if status != 200 {
		t.Fatalf("transcribe status %d", status)
	}
	tp := tr
	mustStr(t, "id", tp.Id)
	mustStr(t, "object", tp.Object)
	mustStr(t, "model", tp.Model)
	mustStr(t, "filename", tp.Filename)
	mustStr(t, "processed_at", tp.ProcessedAt)
	mustGt0(t, "usage.audio_seconds", tp.Usage.AudioSeconds)
	_ = tp.Usage.CreditsConsumed
	_ = tp.Usage.ProcessingMs
	mustStr(t, "language", tp.Language)
	mustGt0(t, "duration_seconds", tp.DurationSeconds)
	mustStr(t, "text", tp.Text)
	mustStr(t, "dialog", tp.Dialog)
	for _, ev := range tp.AudioEvents {
		mustStr(t, "event", ev.Event)
		_ = ev.StartS
		_ = ev.EndS
		_ = ev.DurationS
		mustStr(t, "formatted_timestamp", ev.FormattedTimestamp)
	}
	_ = tp.EventTypes
	mustGt0Int(t, "word_count", tp.WordCount)
	_ = tp.Input.DurationS
	mustStr(t, "input.original_format", tp.Input.OriginalFormat)
	mustStr(t, "input.codec", tp.Input.Codec)
	_ = tp.Input.SampleRate
	_ = tp.Input.Channels

	dialog := tp.Dialog
	dur := tp.DurationSeconds
	txt := tp.Text

	dbody := falaai.DiagnosticRequest{Dialog: &dialog, DurationSeconds: dur, Language: "pt-BR", Text: &txt, ClientReferenceId: sptr("e2e-diag-2026-09-22-001")}
	d, httpResp, err := c.AnalysisAPI.CreateDiagnosticV1AnalyzeDiagnosticPost(ctx).DiagnosticRequest(dbody).Execute()
	if err != nil {
		t.Fatal(err)
	}
	status = httpResp.StatusCode
	Log("diagnostic", "POST", "/v1/analyze/diagnostic", dbody, d, fmt.Sprintf("HTTP %d", status), status)
	if status != 200 {
		t.Fatalf("diagnostic status %d", status)
	}
	dp := d
	mustStr(t, "id", dp.Id)
	mustStr(t, "response_language", dp.ResponseLanguage)
	if dp.Object != "analysis" {
		t.Fatalf("object %q != analysis", dp.Object)
	}
	if dp.Analysis.DialogueSummary.Explanation == nil {
		t.Fatalf("analysis.dialogue_summary vazio")
	}
	if dp.Analysis.ContactReason.Explanation == nil {
		t.Fatalf("analysis.contact_reason vazio")
	}
	_ = dp.Analysis.IdentifiedAction
	_ = dp.Analysis.IdentifiedLabel
	_ = dp.Analysis.Sentiment
	_ = dp.Usage.Characters
	_ = dp.Usage.CreditsConsumed
	_ = dp.Usage.ProcessingMs

	abody := falaai.RiskAuditRequest{
		Dialog:           &dialog,
		DurationSeconds:  dur,
		Language:         "pt-BR",
		ResponseLanguage: "pt-BR",
		Text:             &txt,
		CallDirection:    sptr("inbound"),
		Participants: []falaai.Participant{
			{Interlocutor: "Speaker 1", Name: sptr("Mateus"), Role: "agent"},
			{Interlocutor: "Speaker 2", Name: sptr("Cliente"), Role: "client"},
		},
		ClientReferenceId: sptr("e2e-aud-2026-09-22-001"),
	}
	a, httpResp, err := c.AnalysisAPI.CreateRiskAuditV1AnalyzeRiskAuditPost(ctx).RiskAuditRequest(abody).Execute()
	if err != nil {
		t.Fatal(err)
	}
	status = httpResp.StatusCode
	Log("riskAudit", "POST", "/v1/analyze/riskAudit", abody, a, fmt.Sprintf("HTTP %d", status), status)
	if status != 200 {
		t.Fatalf("riskAudit status %d", status)
	}
	pub := a.Response
	mustStr(t, "meta.id", pub.Meta.Id)
	_ = pub.Meta.Usage.Characters
	_ = pub.Meta.Usage.CreditsConsumed
	_ = pub.Meta.Usage.ProcessingMs
	_ = pub.Participants
	_ = pub.Verdict
	_ = pub.Scores
	_ = pub.Detections
	_ = pub.Analysis
	_ = pub.Timeline
	_ = pub.AudioEventModel
	if pub.CategoriesSummary == nil {
		t.Fatalf("categories_summary ausente")
	}
	_ = pub.Indexer
	_ = pub.Summary
	if pub.ActionsI18n == nil {
		t.Fatalf("actions_i18n ausente")
	}
	_ = pub.AuditDecisions
	_ = pub.ScoringExplanation
	_ = pub.HtmlReport
}