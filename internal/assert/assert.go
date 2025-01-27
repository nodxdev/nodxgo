// Package assert provides a set of functions to help with testing of this project.
package assert

import (
	"testing"
)

func NoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}

func Error(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Error("Expected error, got nil")
	}
}

func Equal(t *testing.T, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func NotEqual(t *testing.T, got, want any) {
	t.Helper()
	if got == want {
		t.Errorf("Expected %v, got %v", want, got)
	}
}

func True(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Error("Expected true, got false")
	}
}

func False(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Error("Expected false, got true")
	}
}

func Nil(t *testing.T, got any) {
	t.Helper()
	if got != nil {
		t.Errorf("Expected nil, got %v", got)
	}
}

func NotNil(t *testing.T, got any) {
	t.Helper()
	if got == nil {
		t.Error("Expected not nil, got nil")
	}
}
