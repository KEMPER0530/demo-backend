package chat

import (
	"context"
	"github.com/PullRequestInc/go-gpt3"
	"log"
	"mailform-demo-backend/src/domain"
)

type ChatGptRepository struct{}

func (cgr *ChatGptRepository) SendPrompt(arg domain.ChatGpt, keyID string) (string, error) {
	// var res *string
	ctx := context.Background()
	client := gpt3.NewClient(keyID)

	resp, err := client.Completion(ctx, gpt3.CompletionRequest{
		Prompt:      []string{arg.Prompt},
		MaxTokens:   gpt3.IntPtr(2000),
		Stop:        []string{"."},
		Temperature: gpt3.Float32Ptr(0),
		Echo:        false,
	})
	if err != nil {
		log.Println(err)
		return "", err
	}

	return resp.Choices[0].Text, nil
}
