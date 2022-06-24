// Package tools
// Description:
package tools

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

// Struct2Map Struct2Map
func Struct2Map(obj interface{}) map[interface{}]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[interface{}]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func MapToSlice(input map[interface{}]interface{}) []interface{} {
	output := make([]interface{}, 0)
	for i := 0; i < len(input); i++ {
		output = append(output, input[i])
	}
	return output
}

// GetMonthDay 获得当前月的初始和结束日期
// Date 16:29 2020/8/6
// Param // param null
// return
func GetMonthDay(now time.Time) (string, string) {
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()
	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	f := firstOfMonth.Unix()
	l := lastOfMonth.Unix()
	return time.Unix(f, 0).Format("2006-01-02") + " 00:00:00", time.Unix(l, 0).Format("2006-01-02") + " 23:59:59"
}

// GetWeekDay 获得当前周的初始和结束日期
// Date 16:32 2020/8/6
// Param // param null
// return
func GetWeekDay(now time.Time) (string, string) {
	offset := int(time.Monday - now.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if offset > 0 {
		offset = -6
	}

	lastoffset := int(time.Saturday - now.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if lastoffset == 6 {
		lastoffset = -1
	}

	firstOfWeek := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	lastOfWeeK := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, lastoffset+1)
	f := firstOfWeek.Unix()
	l := lastOfWeeK.Unix()
	return time.Unix(f, 0).Format("2006-01-02") + " 00:00:00", time.Unix(l, 0).Format("2006-01-02") + " 23:59:59"
}

// GetQuarterDay //获得当前季度的初始和结束日期
// Date 16:33 2020/8/6
// Param // param null
// return
func GetQuarterDay(now time.Time) (string, string) {
	year := now.Format("2006")
	month := int(time.Now().Month())
	var firstOfQuarter string
	var lastOfQuarter string
	if month >= 1 && month <= 3 {
		//1月1号
		firstOfQuarter = year + "-01-01 00:00:00"
		lastOfQuarter = year + "-03-31 23:59:59"
	} else if month >= 4 && month <= 6 {
		firstOfQuarter = year + "-04-01 00:00:00"
		lastOfQuarter = year + "-06-30 23:59:59"
	} else if month >= 7 && month <= 9 {
		firstOfQuarter = year + "-07-01 00:00:00"
		lastOfQuarter = year + "-09-30 23:59:59"
	} else {
		firstOfQuarter = year + "-10-01 00:00:00"
		lastOfQuarter = year + "-12-31 23:59:59"
	}
	return firstOfQuarter, lastOfQuarter
}

// GetBetweenDates 根据开始日期和结束日期计算出时间段内所有日期
// 参数为日期格式，如：2020-01-01
func GetBetweenDates(sdate, edate string) []string {
	d := []string{}
	timeFormatTpl := "2006-01-02 15:04:05"
	if len(timeFormatTpl) != len(sdate) {
		timeFormatTpl = timeFormatTpl[0:len(sdate)]
	}
	date, err := time.Parse(timeFormatTpl, sdate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	date2, err := time.Parse(timeFormatTpl, edate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return d
	}
	// 输出日期格式固定
	timeFormatTpl = "2006-01-02"
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}

//StringStrip 字符串去除空格
func StringStrip(input string) string {
	if input == "" {
		return ""
	}
	return strings.Join(strings.Fields(input), "")
}

//Pivot 行列转换
func Pivot(A [][]float64) [][]float64 {
	// 交换行和列索引
	result := make([][]float64, len(A[0]))
	for i, _ := range result {
		result[i] = make([]float64, len(A))
	}

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			result[j][i] = A[i][j]
		}
	}
	return result
}

type ArrayExist interface {
	Exist() bool
}

type ArrayInt struct {
	Arr []int
	Obj int
}

func (a ArrayInt) Exist() bool {
	for _, v := range a.Arr {
		if a.Obj == v {
			return true
		}
	}
	return false
}

//ExistsObj 判断元素是否在数组内
func ExistsObj(arr []string, e string) (ok bool) {
	for _, v := range arr {
		if e == v {
			return true
		}
	}
	return
}

//ExistsInt 判断元素是否在数组内
func ExistsInt(arr []int, e int) (ok bool) {
	for _, v := range arr {
		if e == v {
			return true
		}
	}
	return
}

// UtcHour utc小时转换为本地时间
func UtcHour(hour string) (h string) {
	if len(hour) < 2 {
		hour = "0" + hour
	}
	t := fmt.Sprintf("2006-01-02 %s:00", hour)
	//date, err := time.Parse(timeFormatTpl, hour)
	date, _ := time.ParseInLocation("2006-01-02 15:04:05", t, time.UTC)
	// 2019-06-26 22:06:00 +0800 CST
	h = fmt.Sprintf("%d:00", date.Local().Hour())
	return
}
