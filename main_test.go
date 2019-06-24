package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
  actual := HelloWorld("hoge")
  expected := "Hello World, hoge"
  if actual != expected {
    t.Errorf("actual %v\nwant %v", actual, expected)
  }
}
