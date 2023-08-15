// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// VirtualMachinesColumns holds the columns for the "virtual_machines" table.
	VirtualMachinesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "owner_id", Type: field.TypeString},
		{Name: "location_id", Type: field.TypeString},
		{Name: "userdata", Type: field.TypeString},
	}
	// VirtualMachinesTable holds the schema information for the "virtual_machines" table.
	VirtualMachinesTable = &schema.Table{
		Name:       "virtual_machines",
		Columns:    VirtualMachinesColumns,
		PrimaryKey: []*schema.Column{VirtualMachinesColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "virtualmachine_created_at",
				Unique:  false,
				Columns: []*schema.Column{VirtualMachinesColumns[1]},
			},
			{
				Name:    "virtualmachine_updated_at",
				Unique:  false,
				Columns: []*schema.Column{VirtualMachinesColumns[2]},
			},
			{
				Name:    "virtualmachine_owner_id",
				Unique:  false,
				Columns: []*schema.Column{VirtualMachinesColumns[4]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		VirtualMachinesTable,
	}
)

func init() {
}
