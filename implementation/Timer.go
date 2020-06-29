package implementation

import "time"

type Timer struct{}

func (t *Timer) CurrentTimestamp() int64 {
	return time.Now().Unix()
}
