package model

import (
	"github.com/achhapolia10/anjaniv2/opdatabase"
)

//NewGroup create a new Group
func NewGroup(g opdatabase.Group) {
	opdatabase.CreateGroup(g)
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
