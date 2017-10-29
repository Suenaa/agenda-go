package model

import(
	"time"
	"fmt"
)

type Date struct {
	Time time.Time
}

//the example of st: 2017-10-20T19:00
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

//month is like January, March and so on
func (date Date) GetMonth() string {
	return date.Time.Month().String()
}


func (date Date) GetDay() int {
	return date.Time.Day()
}

func (date Date) GetHour() int {
	return date.Time.Hour()
}

func (date Date) GetMinute() int {
	return date.Time.Minute()
}


func (date Date)IsEqual(other Date) bool {
	return date.Time.Equal(other.Time)
}

func (date Date)IsAfter(other Date) bool {
	return date.Time.Before(other.Time)
}

//output is like: 2017-10-20 19:00
func (date Date)DateToString() string {
	st := date.Time.String()[0 : 16]
	return st
}

//the example of string: 2017-10-20T19:00
func StringToDate(date string) (time.Time, error) {
	return time.Parse(time.RFC3339, date + ":00Z")
}