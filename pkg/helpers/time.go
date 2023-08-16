package helpers

import "time"

func DurationSecond(duration time.Duration) time.Duration {
	return time.Second * duration
}
