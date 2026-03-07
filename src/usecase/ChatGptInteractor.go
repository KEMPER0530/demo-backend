package usecase

import (
	"github.com/google/uuid"
	"github.com/guregu/dynamo"
	"github.com/kemper0530/demo-backend/src/domain"
)

type ChatGptInteractor struct {
	CGR ChatGptRepository
}

var newRandomUUID = uuid.NewRandom
var tableFromDB = func(d *dynamo.DB, tableName string) dynamo.Table {
	return d.Table(tableName)
}

func (i *ChatGptInteractor) PutChatGptResult(arg domain.ChatGptResult, d *dynamo.DB) (domain.Res, error) {
	// Ensure table exists
	err := i.CGR.CreateTableIfNotExists(d, "ChatGptResult")
	if err != nil {
		return domain.Res{Response: 500, Result: "failed"}, err
	}

	// Generate a new UUID
	id, err := newRandomUUID()
	if err != nil {
		return domain.Res{Response: 500, Result: "failed to generate UUID"}, err
	}

	// Create a new instance of ChatGptResult with the new UUID
	arg.MessageID = id.String()

	table := tableFromDB(d, "ChatGptResult")
	err = i.CGR.PutResult(&table, arg)
	if err != nil {
		return domain.Res{Response: 500, Result: "failed"}, err
	}
	return domain.Res{Response: 200, Result: "success"}, nil
}
