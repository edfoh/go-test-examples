package example5

import "time"

var (
	Now = func() time.Time { return time.Now() }
)

func NowPlusOneHour() time.Time {
	return Now().Add(1 * time.Hour)
}
