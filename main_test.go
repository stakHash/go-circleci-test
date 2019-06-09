package main

import "testing"

func TestHelloSuccess(t *testing.T) {
	suspect := "hello, go"
	result := hello()

	if suspect != result {
		t.Fatal("")
	}
}

func TestHelloFailed(t *testing.T) {
	suspect := "hello, failed"
	result := hello()

	if suspect != result {
		t.Fatal("err")
	}
}
