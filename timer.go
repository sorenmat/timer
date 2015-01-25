package timer

import "time"

var (
	start   int64
	lastRun int64
)

// Start starts the timer, and keeps that value in a variable
func Start() {
	start = time.Now().UnixNano()
}

// Stop stops the timer and registers the time. It returns the difference between start and stop in nano seconds.
// A call to TimeSpent can be called to retrieve this value later.
func Stop() int64 {
	lastRun = time.Now().UnixNano()
	return lastRun - start
}

// TimeSpent returns the difference between the call to Start and Stop in Nano seconds
func TimeSpent() int64 {
	return lastRun - start
}
