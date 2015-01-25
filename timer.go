package timer

import "time"

var (
	start   int64
	lastRun int64
)

// TimerStart starts the timer, and keeps that value in a variable
func TimerStart() {
	start = time.Now().UnixNano()
}

// TimerStop stops the timer and registers the time. It returns the difference between start and stop in nano seconds.
// A call to TimeSpent can be called to retrieve this value later.
func TimerStop() int64 {
	lastRun = time.Now().UnixNano()
	return lastRun - start
}

// TimeSpent returns the difference between the call to TimerStart and TimerStop in Nano seconds
func TimeSpent() int64 {
	return lastRun - start
}
