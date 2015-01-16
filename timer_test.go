package main

import (
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	TimerStart()
	time.Sleep(50 * time.Millisecond)
	timeSpent := TimerStop()
	if timeSpent <= 50 {
		t.Error()
	}

}
func TestTimerWithTimeSpent(t *testing.T) {
	TimerStart()
	time.Sleep(50 * time.Millisecond)
	TimerStop()
	if TimeSpent() <= 50 {
		t.Error()
	}

}
