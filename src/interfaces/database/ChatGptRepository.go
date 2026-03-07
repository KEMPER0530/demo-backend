package database

import (
	"github.com/guregu/dynamo"
	"github.com/kemper0530/demo-backend/src/domain"
	"time"
)

type ChatGptRepository struct{}

var listTables = func(d *dynamo.DB) ([]string, error) {
	return d.ListTables().All()
}

var createTable = func(d *dynamo.DB, tableName string) error {
	return d.CreateTable(tableName, domain.ChatGptResult{}).Run()
}

var sleepTableCreation = time.Sleep

var putResultToTable = func(table *dynamo.Table, arg domain.ChatGptResult) error {
	return table.Put(arg).Run()
}

func (cgr *ChatGptRepository) CreateTableIfNotExists(d *dynamo.DB, tableName string) error {
	// Check if table already exists
	tables, err := listTables(d)
	if err != nil {
		return err
	}

	for _, table := range tables {
		if table == tableName {
			return nil
		}
	}

	// If table does not exist, create table
	err = createTable(d, tableName)
	if err != nil {
		return err
	}

	// Wait for a while to allow AWS to create the table
	sleepTableCreation(20 * time.Second)

	return nil
}

func (cgr *ChatGptRepository) PutResult(table *dynamo.Table, arg domain.ChatGptResult) error {
	err := putResultToTable(table, arg)
	return err
}
