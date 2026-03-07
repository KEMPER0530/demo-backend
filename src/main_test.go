package main

import (
	"context"
	"errors"
	"github.com/kemper0530/demo-backend/src/domain"
	"testing"
)

func TestMainProductionBranch(t *testing.T) {
	origGetenv := getenvMain
	origPrint := printMain
	origLambdaStart := lambdaStartMain
	origLocalTest := localTestMain
	origLog := logMain
	defer func() {
		getenvMain = origGetenv
		printMain = origPrint
		lambdaStartMain = origLambdaStart
		localTestMain = origLocalTest
		logMain = origLog
	}()

	printCalled := false
	lambdaCalled := false
	localCalled := false

	getenvMain = func(key string) string {
		if key == "GO_ENV" {
			return "production"
		}
		return ""
	}
	printMain = func(v ...interface{}) (int, error) {
		printCalled = true
		return 0, nil
	}
	lambdaStartMain = func(handler interface{}) {
		lambdaCalled = true
	}
	localTestMain = func() (interface{}, error) {
		localCalled = true
		return nil, nil
	}
	logMain = func(v ...interface{}) {}

	main()
	if !printCalled || !lambdaCalled {
		t.Fatalf("production branch did not execute expected functions")
	}
	if localCalled {
		t.Fatalf("local test should not be called in production mode")
	}
}

func TestMainDevelopmentSuccessAndError(t *testing.T) {
	origGetenv := getenvMain
	origPrint := printMain
	origLambdaStart := lambdaStartMain
	origLocalTest := localTestMain
	origLog := logMain
	defer func() {
		getenvMain = origGetenv
		printMain = origPrint
		lambdaStartMain = origLambdaStart
		localTestMain = origLocalTest
		logMain = origLog
	}()

	getenvMain = func(key string) string { return "development" }
	lambdaStartMain = func(handler interface{}) {}
	printMain = func(v ...interface{}) (int, error) { return 0, nil }

	logCount := 0
	logMain = func(v ...interface{}) {
		logCount++
	}

	localTestMain = func() (interface{}, error) {
		return "ok", nil
	}
	main()
	if logCount != 1 {
		t.Fatalf("expected 1 log call on success, got %d", logCount)
	}

	localTestMain = func() (interface{}, error) {
		return nil, errors.New("failed")
	}
	main()
	if logCount != 2 {
		t.Fatalf("expected error log call, got %d", logCount)
	}
}

func TestHandleRequestsDelegatesRouteRequest(t *testing.T) {
	origRoute := routeRequestMain
	defer func() { routeRequestMain = origRoute }()

	expected := domain.Res{Response: 200, Result: "ok"}
	routeRequestMain = func(ctx context.Context, request domain.AppSyncEvent) (interface{}, error) {
		if request.OperationName != "createNuxtMail" {
			t.Fatalf("unexpected request passed")
		}
		return expected, nil
	}

	got, err := HandleRequests(context.Background(), domain.AppSyncEvent{
		OperationName: "createNuxtMail",
		Arguments:     map[string]interface{}{},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != expected {
		t.Fatalf("unexpected response: %#v", got)
	}
}
