package model

import (
	"github.com/achhapolia10/inventory-manager/opdatabase"
)

//NewGroup create a new Group
func NewGroup(g opdatabase.Group) {
	opdatabase.CreateGroup(g)
}

//GetGroups gets all Groups
func GetGroups() ([]opdatabase.Group, bool) {
	g, res := opdatabase.SelectGroup()
	return g, res
}

//EditGroup Edits a Group
func EditGroup(g opdatabase.Group) {
	opdatabase.EditGroup(g)
}

//DeleteGroup Deletes  a group
func DeleteGroup(g opdatabase.Group) {
	opdatabase.DeleteGroup(g)
	opdatabase.DeleteProductByGroup(g)
}
