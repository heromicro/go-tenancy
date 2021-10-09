package scope

import (
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

func SimpleScope(field string, value interface{}, params ...string) func(db *gorm.DB) *gorm.DB {
	if len(params) == 0 {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where(fmt.Sprintf("%s = ?", field), value)
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s %s ?", field, params[0]), value)
	}
}

// FilterDate 日期筛选过滤
//  - 不传默认最近7天
func FilterDate(date, field, perfix string) func(db *gorm.DB) *gorm.DB {
	if date == "" {
		date = "lately7"
	}

	dates := strings.Split(date, "-")
	if len(dates) == 2 {
		start, _ := time.Parse("2006/01/02", dates[0])
		end, _ := time.Parse("2006/01/02", dates[1])
		return FilterBetween(start, end, field, perfix)
	}

	if len(dates) == 1 {
		// { text: '今天', val: 'today' },
		// { text: '昨天', val: 'yesterday' },
		// { text: '最近7天', val: 'lately7' },
		// { text: '最近30天', val: 'lately30' },
		// { text: '本周', val: 'thisweek' },
		// { text: '上周', val: 'lateweek' },
		// { text: '本季度', val: 'quarter' },
		// { text: '本月', val: 'month' },
		// { text: '本年', val: 'year' }
		// TODO: 使用内置函数，可能造成索引失效
		switch date {
		case "today":
			return FilterToday(field, perfix)
		case "yesterday":
			return FilterYesterday(field, perfix)
		case "lately7":
			return FilterLately7(field, perfix)
		case "lately30":
			return FilterLately30(field, perfix)
		case "thisweek":
			return FilterThisWeek(field, perfix)
		case "lastweek":
			return FilterLatelyWeek(field, perfix)
		case "month":
			return FilterMonth(field, perfix)
		case "quarter":
			return FilterQuarter(field, perfix)
		case "year":
			return FilterYear(field, perfix)
		default:
			return nil
		}
	}
	return nil
}

// FilterBetween 时间段
func FilterBetween(start, end time.Time, field, perfix string) func(db *gorm.DB) *gorm.DB {
	if perfix != "" {
		field = fmt.Sprintf("%s.%s", perfix, field)
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s BETWEEN ? AND ?", field), start, end)
	}
}

// FilterToday 今天
func FilterToday(field, perfix string) func(db *gorm.DB) *gorm.DB {
	if perfix != "" {
		field = fmt.Sprintf("%s.%s", perfix, field)
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("TO_DAYS(NOW()) - TO_DAYS(%s) < 1", field))
	}
}

// FilterYesterday 昨天
func FilterYesterday(field, perfix string) func(db *gorm.DB) *gorm.DB {
	if perfix != "" {
		field = fmt.Sprintf("%s.%s", perfix, field)
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("TO_DAYS(NOW()) - TO_DAYS(%s) = 1", field))
	}
}

// FilterLately7 最近7天
func FilterLately7(field, perfix string) func(db *gorm.DB) *gorm.DB {
	if perfix != "" {
		field = fmt.Sprintf("%s.%s", perfix, field)
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("DATE_SUB(CURDATE(),INTERVAL 7 DAY) <= DATE(%s)", field))
	}
}

// FilterThisWeek 本周
func FilterThisWeek(field, perfix string) func(db *gorm.DB) *gorm.DB {
	if perfix != "" {
		field = fmt.Sprintf("%s.%s", perfix, field)
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("YEARWEEK(DATE_FORMAT(%s,'%%Y-%%m-%%d')) = YEARWEEK(NOW())", field))
	}
}

// FilterLatelyWeek 上周
func FilterLatelyWeek(field, perfix string) func(db *gorm.DB) *gorm.DB {
	if perfix != "" {
		field = fmt.Sprintf("%s.%s", perfix, field)
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("YEARWEEK(DATE_FORMAT(%s,'%%Y-%%m-%%d')) = YEARWEEK(NOW()) - 1", field))
	}
}

// FilterLately30 最近30天
func FilterLately30(field, perfix string) func(db *gorm.DB) *gorm.DB {
	if perfix != "" {
		field = fmt.Sprintf("%s.%s", perfix, field)
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("DATE_SUB(CURDATE(),INTERVAL 30 DAY) <= DATE(%s)", field))
	}
}

// FilterMonth 本月
func FilterMonth(field, perfix string) func(db *gorm.DB) *gorm.DB {
	if perfix != "" {
		field = fmt.Sprintf("%s.%s", perfix, field)
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("DATE_FORMAT( %s, '%%Y%%m' ) = DATE_FORMAT( CURDATE() , '%%Y%%m')", field))
	}
}

// FilterQuarter 本季度
func FilterQuarter(field, perfix string) func(db *gorm.DB) *gorm.DB {
	if perfix != "" {
		field = fmt.Sprintf("%s.%s", perfix, field)
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("QUARTER(%s) = QUARTER(NOW()", field))
	}
}

// FilterYear 当年
func FilterYear(field, perfix string) func(db *gorm.DB) *gorm.DB {
	if perfix != "" {
		field = fmt.Sprintf("%s.%s", perfix, field)
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("YEAR(%s)=YEAR(NOW())", field))
	}
}

// FilterBase 基础查询
// - field 字段
// - cond 条件 = ,> , >= e.g.
// - perfix 表前缀
// - value 数值
func FilterBase(field, cond, perfix string, value interface{}) func(db *gorm.DB) *gorm.DB {
	if perfix != "" {
		field = fmt.Sprintf("%s.%s", perfix, field)
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s %s ? ", field, cond), value)
	}
}
