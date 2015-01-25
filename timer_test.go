package timer

import (
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	Start()
	time.Sleep(50 * time.Millisecond)
	timeSpent := Stop()
	if timeSpent <= 50 {
		t.Error()
	}

}
func TestTimerWithTimeSpent(t *testing.T) {
	Start()
	time.Sleep(50 * time.Millisecond)
	Stop()
	if TimeSpent() <= 50 {
		t.Error()
	}

}
