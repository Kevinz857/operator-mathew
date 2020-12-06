package controller

import (
	"github.com/operator-mathew/pkg/controller/mathew"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, mathew.Add)
}
