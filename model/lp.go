package model

import (
	"sort"
	"strconv"
	"strings"

	"github.com/achhapolia10/anjaniv2/opdatabase"
)

//individualJournal stores the individual Journal For Each Day
type individualJournal struct {
	productID int
	date      date
	je        map[string]opdatabase.JournalEntry
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
	Total  float64 `json:"total"`
}

//date is a structure for date
type date struct {
	day   int
	month int
	year  int
}

//GetLabourPayment returns the Labour Payment
func GetLabourPayment(days ...string) []LPEntry {
	var dates []date
	products, err := opdatabase.SelectProductMap()
	LPEntries := make([]LPEntry, 0)
	if !err {
		return LPEntries
	}
	for _, d := range days {
		date := parseDate(d)
		dates = append(dates, date)
	}
	lpe := generateLabourPayment(dates, products)
	for _, v := range lpe {
		LPEntries = append(LPEntries, v)
	}
	return LPEntries
}

//generateLabourPayment generates a labour payment
func generateLabourPayment(d []date, p map[int]opdatabase.Product) map[string]LPEntry {
	var jes []individualJournal

	emptyJournal := opdatabase.JournalEntry{0, "", "", 0, 0, 0}
	emptyLPEntry := LPEntry{"", 0, 0, 0, 0, 0, 0, 0, 0.0}
	LPEntries := make(map[string]LPEntry)
	for _, date := range d {
		for _, product := range p {
			je, _ := opdatabase.SelectJournalEntryMap(date.getString(), product.ID)
			journal := individualJournal{
				product.ID,
				date,
				je,
			}
			jes = append(jes, journal)
		}
	}
	labours := getLabourNames(&jes)
	for _, labour := range labours {
		for _, je := range jes {
			if entry := je.je[labour]; entry != emptyJournal {
				if LPEntries[labour] == emptyLPEntry {
					tempLPEntry := LPEntry{labour, 0, 0, 0, 0, 0, 0, 0, 0.0}
					LPEntries[labour] = tempLPEntry
				}
				tempLPEntry := LPEntries[labour]
				product := p[entry.ProductID]
				switch je.date {
				case d[0]:
					glasses := ((entry.Box * product.BoxQuantity) + entry.Packet) * product.PacketQuantity
					tempLPEntry.Day1 += glasses
					tempLPEntry.Total += float64(glasses) * product.Price / 1000.0
					break
				case d[1]:
					glasses := ((entry.Box * product.BoxQuantity) + entry.Packet) * product.PacketQuantity
					tempLPEntry.Day2 += glasses
					tempLPEntry.Total += float64(glasses) * product.Price / 1000.0
					break
				case d[2]:
					glasses := ((entry.Box * product.BoxQuantity) + entry.Packet) * product.PacketQuantity
					tempLPEntry.Day3 += glasses
					tempLPEntry.Total += float64(glasses) * product.Price / 1000.0
					break
				case d[3]:
					glasses := ((entry.Box * product.BoxQuantity) + entry.Packet) * product.PacketQuantity
					tempLPEntry.Day4 += glasses
					tempLPEntry.Total += float64(glasses) * product.Price / 1000.0
					break
				case d[4]:
					glasses := ((entry.Box * product.BoxQuantity) + entry.Packet) * product.PacketQuantity
					tempLPEntry.Day5 += glasses
					tempLPEntry.Total += float64(glasses) * product.Price / 1000.0
					break
				case d[5]:
					glasses := ((entry.Box * product.BoxQuantity) + entry.Packet) * product.PacketQuantity
					tempLPEntry.Day6 += glasses
					tempLPEntry.Total += float64(glasses) * product.Price / 1000.0
					break
				case d[6]:
					glasses := ((entry.Box * product.BoxQuantity) + entry.Packet) * product.PacketQuantity
					tempLPEntry.Day7 += glasses
					tempLPEntry.Total += float64(glasses) * product.Price / 1000.0
					break
				}
				LPEntries[labour] = tempLPEntry
			}
		}
	}
	return LPEntries
}

func getLabourNames(jes *[]individualJournal) []string {
	var labours []string
	fmap := make(map[string]struct {
		present bool
	})
	for _, je := range *jes {
		for k := range je.je {
			if !fmap[k].present {
				fmap[k] = struct{ present bool }{true}
				labours = append(labours, k)
			}
		}
	}
	sort.Strings(labours)
	return labours
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
