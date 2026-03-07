package usecase

import (
	"errors"
	"github.com/google/uuid"
	"github.com/guregu/dynamo"
	"github.com/kemper0530/demo-backend/src/domain"
	"testing"
)

type fakeChatRepo struct {
	createErr    error
	putErr       error
	createCalled int
	putCalled    int
	createTable  string
	gotArg       domain.ChatGptResult
}

func (f *fakeChatRepo) CreateTableIfNotExists(d *dynamo.DB, tableName string) error {
	f.createCalled++
	f.createTable = tableName
	return f.createErr
}

func (f *fakeChatRepo) PutResult(table *dynamo.Table, arg domain.ChatGptResult) error {
	f.putCalled++
	f.gotArg = arg
	return f.putErr
}

func TestChatGptInteractorPutChatGptResult(t *testing.T) {
	origUUID := newRandomUUID
	origTableFromDB := tableFromDB
	defer func() {
		newRandomUUID = origUUID
		tableFromDB = origTableFromDB
	}()

	repo := &fakeChatRepo{}
	interactor := &ChatGptInteractor{CGR: repo}
	tableFromDB = func(d *dynamo.DB, tableName string) dynamo.Table {
		return dynamo.Table{}
	}

	repo.createErr = errors.New("create failed")
	res, err := interactor.PutChatGptResult(domain.ChatGptResult{User: "u"}, nil)
	if err == nil || res.Response != 500 || res.Result != "failed" {
		t.Fatalf("expected create table failure")
	}

	repo.createErr = nil
	newRandomUUID = func() (uuid.UUID, error) {
		return uuid.Nil, errors.New("uuid failed")
	}
	res, err = interactor.PutChatGptResult(domain.ChatGptResult{User: "u"}, nil)
	if err == nil || res.Result != "failed to generate UUID" {
		t.Fatalf("expected uuid failure")
	}

	newRandomUUID = func() (uuid.UUID, error) {
		return uuid.MustParse("11111111-1111-1111-1111-111111111111"), nil
	}
	repo.putErr = errors.New("put failed")
	res, err = interactor.PutChatGptResult(domain.ChatGptResult{User: "u"}, nil)
	if err == nil || res.Result != "failed" {
		t.Fatalf("expected put failure")
	}

	repo.putErr = nil
	res, err = interactor.PutChatGptResult(domain.ChatGptResult{User: "u"}, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.Response != 200 || res.Result != "success" {
		t.Fatalf("unexpected success response")
	}
	if repo.createTable != "ChatGptResult" || repo.putCalled == 0 {
		t.Fatalf("repository methods were not called")
	}
	if repo.gotArg.MessageID == "" || repo.gotArg.User != "u" {
		t.Fatalf("unexpected stored arg: %#v", repo.gotArg)
	}
}
