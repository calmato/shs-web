package jst

import "time"

var jst = time.FixedZone("Asia/Tokyo", 9*60*60)

// Now 現在時刻
func Now() time.Time {
	return time.Now().In(jst)
}

// Date time.Timeの生成
func Date(year int, month time.Month, day, hour, min, sec, nsec int) time.Time {
	return time.Date(year, month, day, hour, min, sec, nsec, jst)
}

// BeginningOfDay 一日の始まり
func BeginningOfDay(t time.Time) time.Time {
	year, month, day := t.In(jst).Date()
	return Date(year, month, day, 0, 0, 0, 0)
}

// BeginningOfMonth 月の初め
func BeginningOfMonth(year, month int) time.Time {
	return Date(year, time.Month(month), 1, 0, 0, 0, 0)
}

// EndOfDay 一日の終わり
func EndOfDay(t time.Time) time.Time {
	year, month, day := t.In(jst).Date()
	return Date(year, month, day, 23, 59, 59, int(time.Second-time.Nanosecond))
}

// EndOfMonth 月末
func EndOfMonth(year, month int) time.Time {
	return Date(year, time.Month(month), 1, 0, 0, 0, 0).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// Format 形式指定で時間の文字列を返す
func Format(t time.Time, format string) string {
	return t.In(jst).Format(format)
}

// FormatYYYYMMDD YYYYMMDD形式で時間の文字列を返す
func FormatYYYYMMDD(t time.Time) string {
	return Format(t, "20060102")
}

// FormatYYYYMM YYYYMM形式で時間の文字列を返す
func FormatYYYYMM(t time.Time) string {
	return Format(t, "200601")
}

// FormatHHMMSS HHMMSS形式で時間の文字列を返す
func FormatHHMMSS(t time.Time) string {
	return Format(t, "150405")
}

// FormatHHMM HHMM形式で時間の文字列を返す
func FormatHHMM(t time.Time) string {
	return Format(t, "1504")
}

// Parse 形式指定で文字列からtime.Timeを生成
func Parse(format, target string) (time.Time, error) {
	return time.ParseInLocation(format, target, jst)
}

// ParseFromYYYYMMDD YYYYMMDD形式の文字列からtime.Timeを生成
func ParseFromYYYYMMDD(yyyymmdd string) (time.Time, error) {
	return time.ParseInLocation("20060102", yyyymmdd, jst)
}

// ParseFromYYYYMM YYYYMM形式の文字列からtime.Timeを生成
func ParseFromYYYYMM(yyyymm string) (time.Time, error) {
	return time.ParseInLocation("200601", yyyymm, jst)
}

// ParseFromHHMMSS HHMMSS形式の文字列からtime.Timeを生成
func ParseFromHHMMSS(hhmmss string) (time.Time, error) {
	return time.ParseInLocation("150405", hhmmss, jst)
}

// ParseFromUnix Unixtimeからtime.Timeを生成
func ParseFromUnix(unix int64) time.Time {
	return time.Unix(unix, 0).In(jst)
}

// FiscalYear 年度を計算して返す
// - 01月01日 ~ 03月31日: 去年の年を返す
// - 04月01日 ~ 12月31日: 現在の年を返す
func FiscalYear(t time.Time) int {
	year, month, _ := t.In(jst).Date()
	if month < 4 {
		return year - 1
	}
	return year
}
