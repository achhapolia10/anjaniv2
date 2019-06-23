package model

import (
	"strconv"
	"strings"
	"time"
)

//date is a structure for date
type date struct {
	day   int
	month int
	year  int
}

//getFiscal returns the date the fical year starts for the date
func getFiscal(t time.Time) time.Time {
	location := time.Now().Location()
	month := int(t.Month())
	if month < 4 {
		return time.Date(t.Year()-1, time.April, 1, 0, 0, 0, 0, location)
	}
	return time.Date(t.Year(), time.April, 1, 0, 0, 0, 0, location)
}

//parseDate parses time and returns in Format of dd,MM,YYYY in int
func parseDate(s string) date {
	d := strings.Split(s, "-")
	if len(d) != 3 {
		return date{0, 0, 0}
	}
	day, _ := strconv.Atoi(d[2])
	month, _ := strconv.Atoi(d[1])
	year, _ := strconv.Atoi(d[0])
	return date{day, month, year}
}

func parseTime(t time.Time) date {
	d := date{
		day: t.Day(), month: int(t.Month()), year: t.Year(),
	}
	return d
}

//Returns a String for a Date
func (d *date) getString() string {
	day := padDate(d.day)
	month := padDate(d.month)
	year := padDate(d.year)
	return year + "-" + month + "-" + day
}

//getMonth returns the month in type time.Month
func (d *date) getMonth() time.Month {
	switch d.month {
	case 1:
		return time.January
	case 2:
		return time.February
	case 3:
		return time.March
	case 4:
		return time.April
	case 5:
		return time.May
	case 6:
		return time.June
	case 7:
		return time.July
	case 8:
		return time.August
	case 9:
		return time.September
	case 10:
		return time.October
	case 11:
		return time.November
	case 12:
		return time.December
	default:
		return time.January

	}
}

//pad and add a 0 in number if less than 10
func padDate(i int) string {
	s := ""
	if i < 10 {
		s = "0" + strconv.Itoa(i)
		return s
	}
	a := strconv.Itoa(i)
	s = s + a
	return s
}
