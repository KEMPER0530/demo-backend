package database

import (
	"github.com/guregu/dynamo"
	"github.com/kemper0530/demo-backend/src/domain"
	"time"
)

type ChatGptRepository struct{}

func (cgr *ChatGptRepository) CreateTableIfNotExists(d *dynamo.DB, tableName string) error {
	// Check if table already exists
	tables, err := d.ListTables().All()
	if err != nil {
		return err
	}

	for _, table := range tables {
		if table == tableName {
			return nil
		}
	}

	// If table does not exist, create table
	err = d.CreateTable(tableName, domain.ChatGptResult{}).Run()
	if err != nil {
		return err
	}

	// Wait for a while to allow AWS to create the table
	time.Sleep(20 * time.Second)

	return nil
}

func (cgr *ChatGptRepository) PutResult(table *dynamo.Table, arg domain.ChatGptResult) error {
	err := table.Put(arg).Run()
	return err
}
