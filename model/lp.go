package model

import (
	"math"
	"sort"

	"github.com/achhapolia10/inventory-manager/opdatabase"
)

//individualJournal stores the individual Journal For Each Day
type individualJournal struct {
	productID int
	date      Date
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

//GetLabourPayment returns the Labour Payment
func GetLabourPayment(days ...string) LPEntries {
	var dates []Date
	products, err := opdatabase.SelectProductMap()
	LPEntries := make([]LPEntry, 0)
	if !err {
		return LPEntries
	}
	for _, d := range days {
		date := ParseDate(d)
		dates = append(dates, date)
	}
	lpe := generateLabourPayment(dates, products)
	for _, v := range lpe {
		LPEntries = append(LPEntries, v)
	}
	return LPEntries
}

//generateLabourPayment generates a labour payment
func generateLabourPayment(d []Date, p map[int]opdatabase.Product) map[string]LPEntry {
	var jes []individualJournal

	emptyJournal := opdatabase.JournalEntry{ID: 0, Labour: "", Date: "", Box: 0, Packet: 0, ProductID: 0}
	emptyLPEntry := LPEntry{"", 0, 0, 0, 0, 0, 0, 0, 0, 0.0}
	LPEntries := make(map[string]LPEntry)
	for _, date := range d {
		for _, product := range p {
			je, _ := opdatabase.SelectJournalEntryMap(date.GetString(), product.ID)
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
