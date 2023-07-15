package usecase

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/guregu/dynamo"
	"github.com/kemper0530/demo-backend/src/domain"
)

type ChatGptInteractor struct {
	CGR ChatGptRepository
}

func (i *ChatGptInteractor) PutChatGptResult(arg domain.ChatGptResult, d *dynamo.DB) (domain.Res, error) {
	// Ensure table exists
	err := i.CGR.CreateTableIfNotExists(d, "ChatGptResult")
	if err != nil {
		return domain.Res{Response: 500, Result: "failed"}, err
	}

	// Generate a new message ID
	messageID := make([]byte, 16)
	_, err = rand.Read(messageID)
	if err != nil {
		return domain.Res{Response: 500, Result: "failed to generate message ID"}, err
	}

	// Create a new instance of ChatGptResult with the new message ID
	arg.MessageID = hex.EncodeToString(messageID)

	table := d.Table("ChatGptResult")
	err = i.CGR.PutResult(&table, arg)
	if err != nil {
		return domain.Res{Response: 500, Result: "failed"}, err
	}
	return domain.Res{Response: 200, Result: "success"}, nil
}
