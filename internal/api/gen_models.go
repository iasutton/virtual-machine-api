// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package api

import (
	"go.infratographer.com/virtual-machine-api/internal/ent/generated"
	"go.infratographer.com/x/gidx"
)

type AnnotationNamespace struct {
	// The owner of the annotation namespace.
	Owner *ResourceOwner `json:"owner"`
}

type Location struct {
	ID             gidx.PrefixedID                     `json:"id"`
	VirtualMachine *generated.VirtualMachineConnection `json:"virtualMachine"`
}

func (Location) IsEntity() {}

type ResourceOwner struct {
	ID             gidx.PrefixedID                     `json:"id"`
	VirtualMachine *generated.VirtualMachineConnection `json:"virtualMachine"`
}

func (ResourceOwner) IsEntity() {}
