package infrastructure

import (
	"context"
	"errors"
	"github.com/kemper0530/demo-backend/src/domain"
	"os"
	"reflect"
	"testing"
)

func TestDefaultParseLocalArgs(t *testing.T) {
	origGetenv := getenvLocal
	origArgs := os.Args
	defer func() {
		getenvLocal = origGetenv
		os.Args = origArgs
	}()

	getenvLocal = func(key string) string {
		switch key {
		case "AWS_SES_TEST_FROM":
			return "from-default"
		case "AWS_SES_TEST_SUCCESS_TO":
			return "to-default"
		default:
			return ""
		}
	}

	os.Args = []string{"cmd", "-i", "1", "-f", "from", "-t", "to", "-s", "sub", "-b", "body", "-us", "user", "-in", "in", "-ou", "out", "-cr", "2023"}
	args, err := defaultParseLocalArgs()
	if err != nil {
		t.Fatalf("unexpected parse error: %v", err)
	}
	if args.mode != 1 || args.user != "user" || args.createdAt != "2023" {
		t.Fatalf("unexpected parsed args: %#v", args)
	}

	os.Args = []string{"cmd", "-i", "bad"}
	if _, err := defaultParseLocalArgs(); err == nil {
		t.Fatalf("expected parse error")
	}
}

func TestLoadEnvVars(t *testing.T) {
	origLoad := loadEnvFile
	origGetenv := getenvLocal
	defer func() {
		loadEnvFile = origLoad
		getenvLocal = origGetenv
	}()

	getenvLocal = func(key string) string { return "development" }

	calledPath := ""
	loadEnvFile = func(path string) error {
		calledPath = path
		return nil
	}
	if err := loadEnvVars(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if calledPath != "src/infrastructure/development.env" {
		t.Fatalf("unexpected path: %s", calledPath)
	}

	loadEnvFile = func(path string) error {
		return errors.New("load failed")
	}
	if err := loadEnvVars(); err == nil {
		t.Fatalf("expected load error")
	}
}

func TestLocalTestBranches(t *testing.T) {
	origLoad := loadEnvFile
	origParse := parseLocalArgs
	origRoute := routeRequestLocal
	defer func() {
		loadEnvFile = origLoad
		parseLocalArgs = origParse
		routeRequestLocal = origRoute
	}()

	loadEnvFile = func(path string) error { return errors.New("load failed") }
	if _, err := LocalTest(); err == nil {
		t.Fatalf("expected load error")
	}

	loadEnvFile = func(path string) error { return nil }

	parseLocalArgs = func() (localArgs, error) {
		return localArgs{}, errors.New("bad args")
	}
	if _, err := LocalTest(); err == nil {
		t.Fatalf("expected parse error")
	}

	parseLocalArgs = func() (localArgs, error) {
		return localArgs{mode: 9}, nil
	}
	if _, err := LocalTest(); err == nil {
		t.Fatalf("expected invalid mode error")
	}

	parseLocalArgs = func() (localArgs, error) {
		return localArgs{mode: 0, from: "f", to: "t", subject: "s", body: "b"}, nil
	}
	var gotEvent domain.AppSyncEvent
	routeRequestLocal = func(ctx context.Context, request domain.AppSyncEvent) (interface{}, error) {
		gotEvent = request
		return domain.Res{Response: 200, Result: "ok"}, nil
	}
	res, err := LocalTest()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if res.(domain.Res).Response != 200 {
		t.Fatalf("unexpected response")
	}
	expected0 := domain.AppSyncEvent{
		OperationName: "createNuxtMail",
		Arguments: map[string]interface{}{
			"from": "f", "to": "t", "subject": "s", "body": "b",
		},
	}
	if !reflect.DeepEqual(gotEvent, expected0) {
		t.Fatalf("unexpected event: %#v", gotEvent)
	}

	parseLocalArgs = func() (localArgs, error) {
		return localArgs{mode: 1, user: "u", input: "i", output: "o", createdAt: "c"}, nil
	}
	routeRequestLocal = func(ctx context.Context, request domain.AppSyncEvent) (interface{}, error) {
		gotEvent = request
		return nil, errors.New("route error")
	}
	if _, err := LocalTest(); err == nil {
		t.Fatalf("expected route error")
	}
	expected1 := domain.AppSyncEvent{
		OperationName: "createChatGptResult",
		Arguments: map[string]interface{}{
			"user": "u", "input": "i", "output": "o", "createdat": "c",
		},
	}
	if !reflect.DeepEqual(gotEvent, expected1) {
		t.Fatalf("unexpected event: %#v", gotEvent)
	}
}
