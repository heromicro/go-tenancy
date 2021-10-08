package utils

import (
	"time"
)

// GetStartTime 获取开始时间
func GetStartTime(date string) time.Time {
	switch date {
	case "today":
		return time.Now()
	case "yesterday":
		return time.Now().AddDate(0, 0, -1)
	case "week":
		return GetMondayOfWeek()
	case "month":
		return time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.Now().Location())
	case "year":
		return time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.Now().Location())
	case "quarter":
		return GetQuarterDay()
	case "lately7":
		return time.Now().AddDate(0, 0, -7)
	case "lately30":
		return time.Now().AddDate(0, 0, -30)
	default:
		return time.Now().AddDate(0, 0, -7)
	}

}

// GetMondayOfWeek 获取本周周一的日期
func GetMondayOfWeek() time.Time {
	if time.Now().Weekday() == time.Monday {
		return time.Now()
	} else {
		offset := int(time.Monday - time.Now().Weekday())
		if offset > 0 {
			offset = -6
		}
		return time.Now().AddDate(0, 0, offset)
	}
}

// GetQuarterDay 本季度第一天
func GetQuarterDay() time.Time {
	month := int(time.Now().Month())
	if month >= 1 && month <= 3 {
		return time.Date(time.Now().Year(), 1, 1, 0, 0, 0, 0, time.Now().Location())
	} else if month >= 4 && month <= 6 {
		return time.Date(time.Now().Year(), 4, 1, 0, 0, 0, 0, time.Now().Location())
	} else if month >= 7 && month <= 9 {
		return time.Date(time.Now().Year(), 7, 1, 0, 0, 0, 0, time.Now().Location())
	} else {
		return time.Date(time.Now().Year(), 10, 1, 0, 0, 0, 0, time.Now().Location())
	}
}

// GetDatesBetweenTwoDays 获取两时间之间的日期
func GetDatesBetweenTwoDays(start, end time.Time) []string {
	var dates []string
	sub := end.Sub(start).Hours() / 24

	if sub < 0 {
		return dates
	}
	if sub >= 0 && sub <= 1 {
		dates = append(dates, start.Format("01-02"))
		return dates
	}

	dates = append(dates, start.Format("01-02"))
	i := 1
	for sub > 1 {
		dates = append(dates, start.AddDate(0, 0, i).Format("01-02"))
		i++
		sub--
	}

	return dates
}
