package main

import (
	"testing"
)

func TestFoo(t *testing.T) {
	if false {
		t.Error("Intentional error 1")
	}
}

func TestBar(t *testing.T) {
	result := Bar()
	if result != "bar" {
		t.Errorf("Expecting bar, got %s", result)
	}

}

func TestQuz(t *testing.T) {
	t.Error("Intentional error 2")
}
