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

	file, err := os.Open("demo_callcenter.mp3")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	model := "falaai-transcribe-1"
	language := "pt"
	clientReferenceId := "call_202609271408"

// REQUIRED: file (audio) + Authorization (fai_ key)
// OPTIONAL (server defaults): model -> falaai-transcribe-1 | language -> pt | client_reference_id -> (empty)
	r, err := client.CreateTranscription(context.Background(), falaai.TranscriptionParams{
		File:              file,
		Model:             model,
		Language:          language,
		ClientReferenceID: clientReferenceId,
	})
	if err != nil {
		panic(err)
	}
	out, _ := json.MarshalIndent(r, "", "  ")
	fmt.Println(string(out))
}
