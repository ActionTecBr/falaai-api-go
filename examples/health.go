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
	baseURL := os.Getenv("FALAAI_BASE_URL")
	if baseURL == "" {
		baseURL = falaai.DefaultBaseURL
	}
	client := falaai.NewPublicClient(baseURL)

	health, err := client.Health(context.Background())
	if err != nil {
		panic(err)
	}
	out, _ := json.MarshalIndent(health, "", "  ")
	fmt.Println(string(out))
}
