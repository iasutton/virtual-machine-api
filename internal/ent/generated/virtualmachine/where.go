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

package virtualmachine

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"go.infratographer.com/virtual-machine-api/internal/ent/generated/predicate"
	"go.infratographer.com/x/gidx"
)

// ID filters vertices based on their ID field.
func ID(id gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldUpdatedAt, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldName, v))
}

// OwnerID applies equality check predicate on the "owner_id" field. It's identical to OwnerIDEQ.
func OwnerID(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldOwnerID, v))
}

// LocationID applies equality check predicate on the "location_id" field. It's identical to LocationIDEQ.
func LocationID(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldLocationID, v))
}

// Userdata applies equality check predicate on the "userdata" field. It's identical to UserdataEQ.
func Userdata(v []uint8) predicate.VirtualMachine {
	vc := []byte(v)
	return predicate.VirtualMachine(sql.FieldEQ(FieldUserdata, vc))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLTE(FieldUpdatedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldContainsFold(FieldName, v))
}

// OwnerIDEQ applies the EQ predicate on the "owner_id" field.
func OwnerIDEQ(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldOwnerID, v))
}

// OwnerIDNEQ applies the NEQ predicate on the "owner_id" field.
func OwnerIDNEQ(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNEQ(FieldOwnerID, v))
}

// OwnerIDIn applies the In predicate on the "owner_id" field.
func OwnerIDIn(vs ...gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldIn(FieldOwnerID, vs...))
}

// OwnerIDNotIn applies the NotIn predicate on the "owner_id" field.
func OwnerIDNotIn(vs ...gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNotIn(FieldOwnerID, vs...))
}

// OwnerIDGT applies the GT predicate on the "owner_id" field.
func OwnerIDGT(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGT(FieldOwnerID, v))
}

// OwnerIDGTE applies the GTE predicate on the "owner_id" field.
func OwnerIDGTE(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGTE(FieldOwnerID, v))
}

// OwnerIDLT applies the LT predicate on the "owner_id" field.
func OwnerIDLT(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLT(FieldOwnerID, v))
}

// OwnerIDLTE applies the LTE predicate on the "owner_id" field.
func OwnerIDLTE(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLTE(FieldOwnerID, v))
}

// OwnerIDContains applies the Contains predicate on the "owner_id" field.
func OwnerIDContains(v gidx.PrefixedID) predicate.VirtualMachine {
	vc := string(v)
	return predicate.VirtualMachine(sql.FieldContains(FieldOwnerID, vc))
}

// OwnerIDHasPrefix applies the HasPrefix predicate on the "owner_id" field.
func OwnerIDHasPrefix(v gidx.PrefixedID) predicate.VirtualMachine {
	vc := string(v)
	return predicate.VirtualMachine(sql.FieldHasPrefix(FieldOwnerID, vc))
}

// OwnerIDHasSuffix applies the HasSuffix predicate on the "owner_id" field.
func OwnerIDHasSuffix(v gidx.PrefixedID) predicate.VirtualMachine {
	vc := string(v)
	return predicate.VirtualMachine(sql.FieldHasSuffix(FieldOwnerID, vc))
}

// OwnerIDEqualFold applies the EqualFold predicate on the "owner_id" field.
func OwnerIDEqualFold(v gidx.PrefixedID) predicate.VirtualMachine {
	vc := string(v)
	return predicate.VirtualMachine(sql.FieldEqualFold(FieldOwnerID, vc))
}

// OwnerIDContainsFold applies the ContainsFold predicate on the "owner_id" field.
func OwnerIDContainsFold(v gidx.PrefixedID) predicate.VirtualMachine {
	vc := string(v)
	return predicate.VirtualMachine(sql.FieldContainsFold(FieldOwnerID, vc))
}

// LocationIDEQ applies the EQ predicate on the "location_id" field.
func LocationIDEQ(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldEQ(FieldLocationID, v))
}

// LocationIDNEQ applies the NEQ predicate on the "location_id" field.
func LocationIDNEQ(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNEQ(FieldLocationID, v))
}

// LocationIDIn applies the In predicate on the "location_id" field.
func LocationIDIn(vs ...gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldIn(FieldLocationID, vs...))
}

// LocationIDNotIn applies the NotIn predicate on the "location_id" field.
func LocationIDNotIn(vs ...gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldNotIn(FieldLocationID, vs...))
}

// LocationIDGT applies the GT predicate on the "location_id" field.
func LocationIDGT(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGT(FieldLocationID, v))
}

// LocationIDGTE applies the GTE predicate on the "location_id" field.
func LocationIDGTE(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldGTE(FieldLocationID, v))
}

// LocationIDLT applies the LT predicate on the "location_id" field.
func LocationIDLT(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLT(FieldLocationID, v))
}

// LocationIDLTE applies the LTE predicate on the "location_id" field.
func LocationIDLTE(v gidx.PrefixedID) predicate.VirtualMachine {
	return predicate.VirtualMachine(sql.FieldLTE(FieldLocationID, v))
}

// LocationIDContains applies the Contains predicate on the "location_id" field.
func LocationIDContains(v gidx.PrefixedID) predicate.VirtualMachine {
	vc := string(v)
	return predicate.VirtualMachine(sql.FieldContains(FieldLocationID, vc))
}

// LocationIDHasPrefix applies the HasPrefix predicate on the "location_id" field.
func LocationIDHasPrefix(v gidx.PrefixedID) predicate.VirtualMachine {
	vc := string(v)
	return predicate.VirtualMachine(sql.FieldHasPrefix(FieldLocationID, vc))
}

// LocationIDHasSuffix applies the HasSuffix predicate on the "location_id" field.
func LocationIDHasSuffix(v gidx.PrefixedID) predicate.VirtualMachine {
	vc := string(v)
	return predicate.VirtualMachine(sql.FieldHasSuffix(FieldLocationID, vc))
}

// LocationIDEqualFold applies the EqualFold predicate on the "location_id" field.
func LocationIDEqualFold(v gidx.PrefixedID) predicate.VirtualMachine {
	vc := string(v)
	return predicate.VirtualMachine(sql.FieldEqualFold(FieldLocationID, vc))
}

// LocationIDContainsFold applies the ContainsFold predicate on the "location_id" field.
func LocationIDContainsFold(v gidx.PrefixedID) predicate.VirtualMachine {
	vc := string(v)
	return predicate.VirtualMachine(sql.FieldContainsFold(FieldLocationID, vc))
}

// UserdataEQ applies the EQ predicate on the "userdata" field.
func UserdataEQ(v []uint8) predicate.VirtualMachine {
	vc := []byte(v)
	return predicate.VirtualMachine(sql.FieldEQ(FieldUserdata, vc))
}

// UserdataNEQ applies the NEQ predicate on the "userdata" field.
func UserdataNEQ(v []uint8) predicate.VirtualMachine {
	vc := []byte(v)
	return predicate.VirtualMachine(sql.FieldNEQ(FieldUserdata, vc))
}

// UserdataIn applies the In predicate on the "userdata" field.
func UserdataIn(vs ...[]uint8) predicate.VirtualMachine {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = []byte(vs[i])
	}
	return predicate.VirtualMachine(sql.FieldIn(FieldUserdata, v...))
}

// UserdataNotIn applies the NotIn predicate on the "userdata" field.
func UserdataNotIn(vs ...[]uint8) predicate.VirtualMachine {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = []byte(vs[i])
	}
	return predicate.VirtualMachine(sql.FieldNotIn(FieldUserdata, v...))
}

// UserdataGT applies the GT predicate on the "userdata" field.
func UserdataGT(v []uint8) predicate.VirtualMachine {
	vc := []byte(v)
	return predicate.VirtualMachine(sql.FieldGT(FieldUserdata, vc))
}

// UserdataGTE applies the GTE predicate on the "userdata" field.
func UserdataGTE(v []uint8) predicate.VirtualMachine {
	vc := []byte(v)
	return predicate.VirtualMachine(sql.FieldGTE(FieldUserdata, vc))
}

// UserdataLT applies the LT predicate on the "userdata" field.
func UserdataLT(v []uint8) predicate.VirtualMachine {
	vc := []byte(v)
	return predicate.VirtualMachine(sql.FieldLT(FieldUserdata, vc))
}

// UserdataLTE applies the LTE predicate on the "userdata" field.
func UserdataLTE(v []uint8) predicate.VirtualMachine {
	vc := []byte(v)
	return predicate.VirtualMachine(sql.FieldLTE(FieldUserdata, vc))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.VirtualMachine) predicate.VirtualMachine {
	return predicate.VirtualMachine(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.VirtualMachine) predicate.VirtualMachine {
	return predicate.VirtualMachine(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.VirtualMachine) predicate.VirtualMachine {
	return predicate.VirtualMachine(func(s *sql.Selector) {
		p(s.Not())
	})
}
