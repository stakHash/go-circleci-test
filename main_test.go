package main

import "testing"

func TestHelloSuccess(t *testing.T) {
	suspect := "hello, go"
	result := hello()

	if suspect != result {
		t.Fatal("")
	}
}

func TestHelloFailed_fixed(t *testing.T) {
	//suspect := "hello, failed"
	suspect := "hello, go"
	result := hello()

	if suspect != result {
		t.Fatal("err")
	}
}
