//go:build ignore

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

	// REQUIRED: file (.zip/.txt export), start, end, timezone, date_format + Authorization
	// OPTIONAL: gap_minutes (default 720) | min_messages (default 2) | client_reference_id
	f, err := os.Open("demo_whatsapp.zip")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	resp, _, err := client.WhatsappAPI.ExtractConversationsV1WhatsappExtractConversationsPost(context.Background()).
		File(f).
		Start("2024-01-01 00:00").
		End("2024-12-31 23:59").
		Timezone("-3").
		DateFormat("day_first").
		GapMinutes(720).
		MinMessages(2).
		VoiceNote(true).
		Expression(true).
		Execute()
	if err != nil {
		panic(err)
	}

	out, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(out))
}
