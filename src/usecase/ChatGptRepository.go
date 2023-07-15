package usecase

import (
	"github.com/guregu/dynamo"
	"github.com/kemper0530/demo-backend/src/domain"
)

type ChatGptRepository interface {
	PutResult(table *dynamo.Table, arg domain.ChatGptResult) error
	CreateTableIfNotExists(d *dynamo.DB, tableName string) error
}
