package logger

import (
	"fmt"
	"testing"
	"time"
)

type DummyAdapter struct{}

func TestLogger(t *testing.T) {
	l := New().WithOptions(Options{
		ParseCodes: true,
	}).WithBackground(123)

	l.Println()

	fmt.Println("- Log Levels -")
	l.Debug("This is a Debug Message")
	l.Note("This is an Note Message")
	l.Info("This is an Info Message")
	l.Warning("This is a Warning Message")
	l.Error("This is an Error Message")
	l.Fatal("This is a Fatal Message")
	l.Println("Just a normal message")

	l.Println("Just a normal message")

	l.Println("~39~blue text ~160~red text ~r~reset text ~15~white text")

	l.WithNoBackground()

	adp := &DummyAdapter{}

	l.LogHTTPRequest(adp)
	l.LogHTTPRequest(adp)
	l.LogHTTPRequest(adp)
	l.LogHTTPRequest(adp)
}

func (a *DummyAdapter) Method() string {
	return "GET"
}

func (a *DummyAdapter) Path() string {
	return "/test/path"
}

func (a *DummyAdapter) ClientIP() string {
	return "127.0.0.1"
}

func (a *DummyAdapter) StatusCode() int {
	return 200
}

func (a *DummyAdapter) TimeTaken() time.Duration {
	return 123 * time.Millisecond
}
