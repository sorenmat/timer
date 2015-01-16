package timer

import "time"

var (
	start   int64
	lastRun int64
)

func TimerStart() {
	start = time.Now().UnixNano()
}

func TimerStop() int64 {
	lastRun = time.Now().UnixNano()
	return lastRun
}

func TimeSpent() int64 {
	return lastRun
}
