//go:build ignore

package main

import (
	"context"
	"fmt"
	"os"

	falaai "github.com/actiontecbr/falaai-api-go"
)

func main() {
	client, err := falaai.NewClientWithResponses(os.Getenv("FALAAI_BASE_URL"))
	if err != nil {
		panic(err)
	}

	// REQUIRED: file (.zip/.txt export), start, end, timezone, date_format + Authorization
	// OPTIONAL: gap_minutes (default 720) | min_messages (default 2) | chars_per_minute (default 800) | client_reference_id
	f, err := os.Open("demo_whatsapp.zip")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	body := &falaai.ExtractConversationsMultipartRequestBody{
		File:           f,
		Start:          "2024-01-01T00:00:00",
		End:            "2024-12-31T23:59:59",
		Timezone:       "-3",
		DateFormat:     "day_first",
		GapMinutes:     720,
		MinMessages:    2,
		CharsPerMinute: 800,
	}

	resp, err := client.ExtractConversationsWithBodyWithResponse(context.Background(), body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(resp.Body))
}
