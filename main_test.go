package main

import (
	"fmt"
	"os"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	actual := HelloWorld("hoge")
	expected := "Hello World, hoge"
	if actual != expected {
		t.Errorf("actual %v\nwant %v", actual, expected)
	}
}

func TestMain(m *testing.M) {
	fmt.Println("before test")
	code := m.Run()
	fmt.Println("after test")
	os.Exit(code)
}
