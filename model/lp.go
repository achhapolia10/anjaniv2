package model

import (
	"math"
	"sort"
	"strconv"
	"strings"
	"time"

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
	TGlass int     `json:"tglass"`
	Total  float64 `json:"total"`
}

//LPEntries wrapper for LPEntry Array
type LPEntries []LPEntry

//Len length
func (l LPEntries) Len() int { return len(l) }

//Swap swap
func (l LPEntries) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

//Less less
func (l LPEntries) Less(i, j int) bool { return l[i].Labour < l[j].Labour }

//TotalAmmount total ammount
func (l LPEntries) TotalAmmount() float64 {
	var t float64
	t = 0
	for _, a := range l {
		t += math.Round(a.Total)
	}
	return t
}

//date is a structure for date
type date struct {
	day   int
	month int
	year  int
}

//GetLabourPayment returns the Labour Payment
func GetLabourPayment(days ...string) LPEntries {
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

	emptyJournal := opdatabase.JournalEntry{ID: 0, Labour: "", Date: "", Box: 0, Packet: 0, ProductID: 0}
	emptyLPEntry := LPEntry{"", 0, 0, 0, 0, 0, 0, 0, 0, 0.0}
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
					tempLPEntry := LPEntry{labour, 0, 0, 0, 0, 0, 0, 0, 0, 0.0}
					LPEntries[labour] = tempLPEntry
				}
				tempLPEntry := LPEntries[labour]
				product := p[entry.ProductID]
				switch je.date {
				case d[0]:
					glasses := ((entry.Box * product.BoxQuantity) + entry.Packet) * product.PacketQuantity
					tempLPEntry.Day1 += glasses
					tempLPEntry.TGlass += glasses
					tempLPEntry.Total += float64(glasses) * product.Price / 1000.0
					break
				case d[1]:
					glasses := ((entry.Box * product.BoxQuantity) + entry.Packet) * product.PacketQuantity
					tempLPEntry.Day2 += glasses
					tempLPEntry.TGlass += glasses
					tempLPEntry.Total += float64(glasses) * product.Price / 1000.0
					break
				case d[2]:
					glasses := ((entry.Box * product.BoxQuantity) + entry.Packet) * product.PacketQuantity
					tempLPEntry.Day3 += glasses
					tempLPEntry.TGlass += glasses
					tempLPEntry.Total += float64(glasses) * product.Price / 1000.0
					break
				case d[3]:
					glasses := ((entry.Box * product.BoxQuantity) + entry.Packet) * product.PacketQuantity
					tempLPEntry.Day4 += glasses
					tempLPEntry.TGlass += glasses
					tempLPEntry.Total += float64(glasses) * product.Price / 1000.0
					break
				case d[4]:
					glasses := ((entry.Box * product.BoxQuantity) + entry.Packet) * product.PacketQuantity
					tempLPEntry.Day5 += glasses
					tempLPEntry.TGlass += glasses
					tempLPEntry.Total += float64(glasses) * product.Price / 1000.0
					break
				case d[5]:
					glasses := ((entry.Box * product.BoxQuantity) + entry.Packet) * product.PacketQuantity
					tempLPEntry.Day6 += glasses
					tempLPEntry.TGlass += glasses
					tempLPEntry.Total += float64(glasses) * product.Price / 1000.0
					break
				case d[6]:
					glasses := ((entry.Box * product.BoxQuantity) + entry.Packet) * product.PacketQuantity
					tempLPEntry.Day7 += glasses
					tempLPEntry.TGlass += glasses
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
