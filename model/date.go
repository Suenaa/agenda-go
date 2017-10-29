package model

import(
	"time"
	"fmt"
)

type Date struct {
	Time time.Time
}

func (date *Date) Init(st string) {
	t, err := StringToDate(st)
	if err == nil {
		date.Time = t
	} else {
		fmt.Println(err);
	}
}

func (date Date) GetYear() int {
	return date.Time.Year()
}

func (date *Date) SetYear(year int) {
	//date.Time.Year = year
}

//待修改
func (date Date) GetMonth() time.Month {
	return date.Time.Month()
}

func (date *Date) SetMonth(month int) {
	//date.Time.Month = month
}

func (date Date) GetDay() int {
	return date.Time.Day()
}

func (date *Date) SetDay(day int) {
	//date.Time.Day = day
}

func (date Date) GetHour() int {
	return date.Time.Hour()
}

func (date *Date) SetHour(hour int) {
	//date.Time.Hour = hour
}

func (date Date) GetMinute() int {
	return date.Time.Minute()
}

func (date *Date) SetMinute(minute int) {
	//date.Time.Minute = minute
}

func (date Date)IsEqual(other Date) bool {
	return date.Time.Equal(other.Time)
}

func (date Date)IsAfter(other Date) bool {
	return date.Time.Before(other.Time)
}

//待修改
func (date Date)IsValid() bool {
	return false
}

//
func (date Date)DateToString() string {
	return date.Time.String()
}

//
func StringToDate(date string) (time.Time, error) {
	//time.Parse("2017-10-20 19:00", date)
	//if (err != nil) {
	//	return nil
	//}
	return time.Parse("2017-10-20 19:00", date)
}