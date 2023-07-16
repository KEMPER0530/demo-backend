package usecase

import (
	"github.com/google/uuid"
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

	// Generate a new UUID
	id, err := uuid.NewRandom()
	if err != nil {
		return domain.Res{Response: 500, Result: "failed to generate UUID"}, err
	}

	// Create a new instance of ChatGptResult with the new UUID
	arg.MessageID = id.String()

	table := d.Table("ChatGptResult")
	err = i.CGR.PutResult(&table, arg)
	if err != nil {
		return domain.Res{Response: 500, Result: "failed"}, err
	}
	return domain.Res{Response: 200, Result: "success"}, nil
}
