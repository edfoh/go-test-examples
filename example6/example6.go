package example6

import "time"

var (
	Now = time.Now
)

func NowPlusOneHour() time.Time {
	return Now().Add(1 * time.Hour)
}
