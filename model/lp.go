package model

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/achhapolia10/anjaniv2/opdatabase"
)

//individualJournal stores the individual Journal For Each Day
type individualJournal struct {
	product opdatabase.Product
	date    date
	je      map[string]opdatabase.JournalEntry
}

//LPEntry is a Structure for a Labour Payment Entry
type LPEntry struct {
	Labour string  `json:"labour"`
	Day1   int     `json:"day1"`
	Day2   int     `json:"day2"`
	Day3   int     `json:"day3"`
	Day4   int     `json:"day4"`
	Day5   int     `json:"day5"`
	Day6   int     `json:"day6"`
	Day7   int     `json:"day7"`
	Total  float32 `json:"total"`
}

//LPEntries is an slice type for LPEntry
type LPEntries []LPEntry

//date is a structure for date
type date struct {
	day   int
	month int
	year  int
}

//GetLabourPayment returns the Labour Payment
func GetLabourPayment(days ...string) (string, bool) {
	var dates []date
	products, err := opdatabase.SelectProduct()
	if !err {
		return "", false
	}
	for _, d := range days {
		date := parseDate(d)
		dates = append(dates, date)
	}
	generateLabourPayment(dates, products)
	return "Labour Payment", true
}

//generateLabourPayment generates a labour payment
func generateLabourPayment(d []date, p []opdatabase.Product) {
	var jes []individualJournal
	for _, date := range d {
		for _, product := range p {
			je, _ := opdatabase.SelectJournalEntryMap(date.getString(), product.ID)
			journal := individualJournal{
				product,
				date,
				je,
			}
			jes = append(jes, journal)
		}
	}
	getLabourNames(&jes)
}

func getLabourNames(jes *[]individualJournal) {
	fmt.Print((*jes)[0].je["QWE"])
}

/* Functions to perform Various tasks for Labour Payment
 * Includes Parse Date
 * Getting String for a date
 * padding a number with 0
 */

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

//Returns a String for a Date
func (d *date) getString() string {
	day := padDate(d.day)
	month := padDate(d.month)
	year := padDate(d.year)
	return year + "-" + month + "-" + day
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
