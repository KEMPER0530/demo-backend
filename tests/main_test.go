package main

import (
  "fmt"
  "os"
  "testing"
)

func TestMain(m *testing.M) {
	fmt.Println("before test main_test.go")
	fmt.Println("after test main_test.go")
	os.Exit(0)
}
