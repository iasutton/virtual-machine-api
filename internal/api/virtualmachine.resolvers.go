package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.36

import (
	"context"
	"fmt"

	"go.infratographer.com/virtual-machine-api/internal/ent/generated"
	"go.infratographer.com/x/gidx"
)

// VirtualMachine is the resolver for the virtualMachine field.
func (r *queryResolver) VirtualMachine(ctx context.Context, id gidx.PrefixedID) (*generated.VirtualMachine, error) {
	panic(fmt.Errorf("not implemented: VirtualMachine - virtualMachine"))
}
