package jst

import "time"

var jst = time.FixedZone("Asia/Tokyo", 9*60*60)

func Now() time.Time {
	return time.Now().In(jst)
}

func Date(year int, month time.Month, day, hour, min, sec, nsec int) time.Time {
	return time.Date(year, month, day, hour, min, sec, nsec, jst)
}

func BeginningOfDay(t time.Time) time.Time {
	year, month, day := t.In(jst).Date()
	return Date(year, month, day, 0, 0, 0, 0)
}

func BeginningOfMonth(year, month int) time.Time {
	return Date(year, time.Month(month), 1, 0, 0, 0, 0)
}

func EndOfDay(t time.Time) time.Time {
	year, month, day := t.In(jst).Date()
	return Date(year, month, day, 23, 59, 59, int(time.Second-time.Nanosecond))
}

func EndOfMonth(year, month int) time.Time {
	return Date(year, time.Month(month), 1, 0, 0, 0, 0).AddDate(0, 1, 0).Add(-time.Nanosecond)
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

func FormatHHMM(t time.Time) string {
	return Format(t, "1504")
}

func Parse(format, target string) (time.Time, error) {
	return time.ParseInLocation(format, target, jst)
}

func ParseFromYYYYMMDD(yyyymmdd string) (time.Time, error) {
	return time.ParseInLocation("20060102", yyyymmdd, jst)
}

func ParseFromYYYYMM(yyyymm string) (time.Time, error) {
	return time.ParseInLocation("200601", yyyymm, jst)
}

func ParseFromHHMMSS(hhmmss string) (time.Time, error) {
	return time.ParseInLocation("150405", hhmmss, jst)
}

func ParseFromUnix(unix int64) time.Time {
	return time.Unix(unix, 0).In(jst)
}
