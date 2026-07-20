package types

import (
	"time"
)

func DateString(t time.Time) string { _ = "STUB: not implemented"; return "" }

func prevalidateDate(s string, relaxed bool) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func parseTimezoneHours(s string, o byte) (int, bool) { _ = "STUB: not implemented"; return 0, false }

func parseTimezoneMinutes(s string, o byte) (int, bool) { _ = "STUB: not implemented"; return 0, false }

func timezoneSeparator(c byte) bool { _ = "STUB: not implemented"; return false }

func emptyTimeZone(t string, relaxed bool) bool { _ = "STUB: not implemented"; return false }

func parseTimezone(s string, off int, relaxed bool) (h, m int, ok bool) {
	_ = "STUB: not implemented"
	return 0, 0, false
}

func parseYear(s string) (y int, finished, ok bool) {
	_ = "STUB: not implemented"
	return 0, false, false
}

func parseMonth(s string) (m int, finished, ok bool) {
	_ = "STUB: not implemented"
	return 0, false, false
}

func parseDay(s string, y, m int) (d int, finished, ok bool) {
	_ = "STUB: not implemented"
	return 0, false, false
}

func parseHour(s string) (h int, finished, ok bool) {
	_ = "STUB: not implemented"
	return 0, false, false
}

func parseMinute(s string) (min int, finished, ok bool) {
	_ = "STUB: not implemented"
	return 0, false, false
}

func parseSecond(s string) (sec int, finished bool, off int, ok bool) {
	_ = "STUB: not implemented"
	return 0, false, 0, false
}

func digestPopularOutOfSpecDates(s string) (time.Time, bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

func DateTime(s string, relaxed bool) (time.Time, bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}
