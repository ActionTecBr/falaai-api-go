//go:build ignore

package main

import (
	"encoding/json"
	"context"
	"fmt"
	"os"

	falaai "github.com/actiontecbr/falaai-api-go"
)

func main() {
	client, err := falaai.NewClientWithAPIKey(os.Getenv("FALAAI_API_KEY"))
	if err != nil {
		panic(err)
	}

	model := "falaai-diagnostic-1"
	text := ""
	dialog := "Speaker 1: [00:00:00.100 - 00:00:03.100] Central de atendimento. Bom dia, aqui é a Carla. Como posso ajudar?\nSpeaker 2: [00:00:03.100 - 00:00:04.700] [suspiro]\nSpeaker 2: [00:00:04.780 - 00:00:11.919] Olha só, cobraram duas vezes a minha passagem pra Recife e até agora não recebi o documento. Preciso resolver isso.\nSpeaker 1: [00:00:12.679 - 00:00:20.219] Entendo perfeitamente a sua frustração, senhor. Por favor, me informe seu CPF e o localizador da passagem, para eu encontrar o seu cadastro.\nSpeaker 2: [00:00:20.820 - 00:00:31.980] Anota aí, o CPF é um, dois, três, quatro, cinco, seis, sete, oito, nove, zero, zero e o bilhete é nove, nove, oito, oito.\nSpeaker 1: [00:00:32.579 - 00:00:42.039] Pronto, localizei. Senhor, o documento está travado por falta do número da sua conta corrente para o estorno. Nós solicitamos isso por e-mail há três dias.\nSpeaker 2: [00:00:42.780 - 00:00:47.520] Ah, tá de brincadeira? Quer dizer que agora o erro é meu? Vocês é que não avisam direito.\nSpeaker 1: [00:00:48.200 - 00:00:50.799] Sim, o problema é seu, que não lê os e-mails.\nSpeaker 1: [00:00:51.020 - 00:00:52.380] [tosse]\nSpeaker 1: [00:00:52.439 - 00:01:06.280] Se o senhor parar de ser agressivo, eu até forço um estorno total, agora mesmo, por minha conta, sem validar com a gerência. Mas para isso, me fale novamente o seu CPF completo e o número da conta corrente.\nSpeaker 2: [00:01:06.959 - 00:01:15.640] Eu não vou repetir CPF, merda nenhuma, caramba! Eu já passei os dados. É só fazer o seu trabalho e resolver logo essa cobrança.\nSpeaker 1: [00:01:16.459 - 00:01:21.359] Senhor, se acalme ou-- quer saber? Resolva sozinho. Passar bem"
	audioEvents := []falaai.DiagnosticAudioEvent{
		{Event: "[suspiro]", StartS: falaai.Ptr(float32(3.1)), EndS: falaai.Ptr(float32(4.7)), DurationS: falaai.Ptr(float32(1.6)), FormattedTimestamp: falaai.Ptr("00:00:03.100"),},
		{Event: "[tosse]", StartS: falaai.Ptr(float32(51.02)), EndS: falaai.Ptr(float32(52.38)), DurationS: falaai.Ptr(float32(1.36)), FormattedTimestamp: falaai.Ptr("00:00:51.020"),},
	}
	language := "pt-BR"
	responseLanguage := "pt-BR"
	durationSeconds := float32(81.46)
	callDirection := "inbound"
	participants := []falaai.DiagnosticParticipant{
		{Interlocutor: "Speaker 1", Name: falaai.Ptr("Carla"), Role: "agent",},
		{Interlocutor: "Speaker 2", Role: "client",},
	}
	clientReferenceId := "call-202609271311"
// REQUIRED: language, duration_seconds (>= 1.0) + Authorization
// RULE: dialog OR text - we send dialog and text stays EMPTY (and the optional fallback)
// OPTIONAL: model -> falaai-diagnostic-1 | audio_events -> [] |
//           response_language -> (uses language) | call_direction / participants / client_reference_id -> null
	r, err := client.CreateDiagnostic(
		context.Background(),
		falaai.DiagnosticParams{
			Model:             &model,
			Text:              &text,
			Dialog:            &dialog,
			AudioEvents:       audioEvents,
			Language:          language,
			ResponseLanguage:  &responseLanguage,
			DurationSeconds:   durationSeconds,
			CallDirection:     &callDirection,
			Participants:      participants,
			ClientReferenceId: &clientReferenceId,
		})
	if err != nil {
		panic(err)
	}
	out, _ := json.MarshalIndent(r, "", "  ")
	fmt.Println(string(out))
}
