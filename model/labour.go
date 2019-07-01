package model

import (
	"time"

	"github.com/achhapolia10/anjaniv2/opdatabase"
)

//GetLabourNames gets all the unique usernames
func GetLabourNames(ch chan map[string]bool) {
	names := opdatabase.SelectLabours()
	ch <- names
}

//UpdateLabourNames updates the labour names
func UpdateLabourNames(n, d string, names map[string]bool) {
	if names[n] {
		opdatabase.UpdateLabour(n, d)
	} else {
		opdatabase.AddLabour(n, d)
	}
}

//DeleteLabours Deletes the labour names prior to 2 weeks of the current entry
func DeleteLabours() {
	now := time.Now()

	prev := now.AddDate(0, 0, -14)

	to := now.AddDate(0, 0, -21)

	for prev.Before(to) || prev.Equal(to) {

		d := ParseTime(prev)
		opdatabase.DeleteLabours(d.GetString())
		prev.AddDate(0, 0, 1)

	}

}
