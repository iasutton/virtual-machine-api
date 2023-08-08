package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"go.infratographer.com/x/entx"
	"go.infratographer.com/x/gidx"
)

type VirtualMachine struct {
	ent.Schema
}

func (VirtualMachine) Mixin() []ent.Mixin {
	return []ent.Mixin{
		entx.NewTimestampMixin(),
	}
}

func (VirtualMachine) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			GoType(gidx.PrefixedID("")).
			Unique().
			Immutable().
			Comment("The ID of the VirtualMachine.").
			Annotations(
				entgql.OrderField("ID"),
			).
			DefaultFunc(func() gidx.PrefixedID { return gidx.MustNewID(VirtualMachinePrefix) }),
		field.Text("name").
			NotEmpty().
			Comment("The name of the Virtual Machine.").
			Annotations(
				entgql.QueryField(),
				entgql.Type("String"),
				entgql.OrderField("NAME"),
			),
	}
}

func (VirtualMachine) Edges() []ent.Edge {
	return nil
}

func (VirtualMachine) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"), /* XXX think this is wrong -ians */
	}
}

func (VirtualMachine) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entx.GraphKeyDirective("id"),
		schema.Comment("Represents a virtual machine on the graph."),
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(
			entgql.MutationCreate().Description("Create a new virtual machine."),
			entgql.MutationUpdate().Description("Update an existing virtual machine."), /* XXX */
		),
	}
}
