package logger

import (
	"fmt"
	"testing"
)

func TestLogger(t *testing.T) {
	l := New().WithOptions(Options{
		ParseCodes: true,
	})

	l.PrintLn()

	fmt.Println("- Log Levels -")
	l.Debug("This is a Debug Message")
	l.Note("This is an Note Message")
	l.Info("This is an Info Message")
	l.Warning("This is a Warning Message")
	l.Error("This is an Error Message")
	l.Fatal("This is a Fatal Message")
	l.PrintLn("Just a normal message")

	l.PrintLn("Just a normal message")

	l.PrintLn("~199~pink text ~39~blue text ~160~red text ~r~reset text ~15~white text")
}

func TestColor(t *testing.T) {
	l := New().WithOptions(Options{
		NoTime:     true,
		ParseCodes: true,
	})

	l.Info("~199~pink text ~39~blue text ~160~red text")
}
