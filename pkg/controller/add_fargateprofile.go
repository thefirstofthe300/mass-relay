package controller

import (
	"github.com/thefirstofthe300/mass-relay/pkg/controller/fargateprofile"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, fargateprofile.Add)
}
