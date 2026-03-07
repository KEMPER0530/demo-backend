package database

import (
	"errors"
	"github.com/guregu/dynamo"
	"github.com/kemper0530/demo-backend/src/domain"
	"testing"
	"time"
)

func TestCreateTableIfNotExists(t *testing.T) {
	origList := listTables
	origCreate := createTable
	origSleep := sleepTableCreation
	defer func() {
		listTables = origList
		createTable = origCreate
		sleepTableCreation = origSleep
	}()

	repo := &ChatGptRepository{}

	listTables = func(d *dynamo.DB) ([]string, error) {
		return nil, errors.New("list failed")
	}
	if err := repo.CreateTableIfNotExists(nil, "ChatGptResult"); err == nil {
		t.Fatalf("expected list error")
	}

	listTables = func(d *dynamo.DB) ([]string, error) {
		return []string{"ChatGptResult"}, nil
	}
	createCalled := false
	createTable = func(d *dynamo.DB, tableName string) error {
		createCalled = true
		return nil
	}
	if err := repo.CreateTableIfNotExists(nil, "ChatGptResult"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if createCalled {
		t.Fatalf("create table should not run when table exists")
	}

	listTables = func(d *dynamo.DB) ([]string, error) {
		return []string{}, nil
	}
	createTable = func(d *dynamo.DB, tableName string) error {
		return errors.New("create failed")
	}
	if err := repo.CreateTableIfNotExists(nil, "ChatGptResult"); err == nil {
		t.Fatalf("expected create error")
	}

	sleepDuration := time.Duration(0)
	createTable = func(d *dynamo.DB, tableName string) error {
		return nil
	}
	sleepTableCreation = func(d time.Duration) {
		sleepDuration = d
	}
	if err := repo.CreateTableIfNotExists(nil, "ChatGptResult"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if sleepDuration != 20*time.Second {
		t.Fatalf("unexpected sleep duration: %v", sleepDuration)
	}
}

func TestPutResult(t *testing.T) {
	origPut := putResultToTable
	defer func() { putResultToTable = origPut }()

	repo := &ChatGptRepository{}
	putResultToTable = func(table *dynamo.Table, arg domain.ChatGptResult) error {
		return errors.New("put failed")
	}
	if err := repo.PutResult(nil, domain.ChatGptResult{}); err == nil {
		t.Fatalf("expected put error")
	}

	putResultToTable = func(table *dynamo.Table, arg domain.ChatGptResult) error {
		return nil
	}
	if err := repo.PutResult(nil, domain.ChatGptResult{}); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
