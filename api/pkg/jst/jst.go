package jst

import "time"

var jst = time.FixedZone("JST", 9*60*60)

func Now() time.Time {
	return time.Now().In(jst)
}

func Date(year int, month time.Month, day, hour, min, sec, nsec int) time.Time {
	return time.Date(year, month, day, hour, min, sec, nsec, jst)
}

func BegginingOfDay(t time.Time) time.Time {
	year, month, day := t.In(jst).Date()
	return Date(year, month, day, 0, 0, 0, 0)
}

func BegginingOfWeek(t time.Time, week time.Weekday) time.Time {
	t = BegginingOfDay(t)
	days := int(t.Weekday() - week)
	if t.Weekday() < week {
		days += 7
	}
	return t.AddDate(0, 0, -days)
}

func BeginningOfMonth(year, month int) time.Time {
	return Date(year, time.Month(month), 1, 0, 0, 0, 0)
}

func EndOfMonth(year, month int) time.Time {
	return Date(year, time.Month(month+1), 0, 0, 0, 0, 0)
}

func Format(t time.Time, format string) string {
	return t.In(jst).Format(format)
}

func FormatYYYYMMDD(t time.Time) string {
	return Format(t, "20060102")
}

func FormatYYYYMM(t time.Time) string {
	return Format(t, "200601")
}

func FormatHHMMSS(t time.Time) string {
	return Format(t, "150405")
}

func Parse(format, target string) (time.Time, error) {
	return time.ParseInLocation(format, target, jst)
}

func ParseFromYYYYMMDD(yyyymmdd string) (time.Time, error) {
	return time.ParseInLocation("20060102", yyyymmdd, jst)
}
