package model

import (
	"fmt"

	"github.com/achhapolia10/anjaniv2/opdatabase"
)

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

//GetLabourPayment returns the Labour Payment
func GetLabourPayment(f string, t string) (string, bool) {
	products, err := opdatabase.SelectProduct()
	if !err {
		return "", false
	}
	
	return "Labour Payment", true
}

/* Functions to perform Various tasks for Labour Payment */
