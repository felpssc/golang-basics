package main

import (
	"runtime"
	"testing"
)

func TestDependency(t *testing.T) {
	t.Log(runtime.GOARCH)

	if runtime.GOARCH == "arm64" {
		t.Skip("Não compatível com essa arquitetura.")
	}
	//...
	t.Fail()
}
