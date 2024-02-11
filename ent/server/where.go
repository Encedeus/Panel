// Code generated by ent, DO NOT EDIT.

package server

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Encedeus/panel/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Server {
	return predicate.Server(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Server {
	return predicate.Server(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Server {
	return predicate.Server(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Server {
	return predicate.Server(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Server {
	return predicate.Server(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Server {
	return predicate.Server(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Server {
	return predicate.Server(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Server {
	return predicate.Server(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Server {
	return predicate.Server(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Server {
	return predicate.Server(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Server {
	return predicate.Server(sql.FieldEQ(FieldUpdatedAt, v))
}

// RAM applies equality check predicate on the "ram" field. It's identical to RAMEQ.
func RAM(v uint) predicate.Server {
	return predicate.Server(sql.FieldEQ(FieldRAM, v))
}

// Storage applies equality check predicate on the "storage" field. It's identical to StorageEQ.
func Storage(v uint) predicate.Server {
	return predicate.Server(sql.FieldEQ(FieldStorage, v))
}

// LogicalCores applies equality check predicate on the "logical_cores" field. It's identical to LogicalCoresEQ.
func LogicalCores(v uint) predicate.Server {
	return predicate.Server(sql.FieldEQ(FieldLogicalCores, v))
}

// Port applies equality check predicate on the "port" field. It's identical to PortEQ.
func Port(v uint16) predicate.Server {
	return predicate.Server(sql.FieldEQ(FieldPort, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Server {
	return predicate.Server(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Server {
	return predicate.Server(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Server {
	return predicate.Server(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Server {
	return predicate.Server(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Server {
	return predicate.Server(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Server {
	return predicate.Server(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Server {
	return predicate.Server(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Server {
	return predicate.Server(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Server {
	return predicate.Server(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Server {
	return predicate.Server(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Server {
	return predicate.Server(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Server {
	return predicate.Server(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Server {
	return predicate.Server(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Server {
	return predicate.Server(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Server {
	return predicate.Server(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Server {
	return predicate.Server(sql.FieldLTE(FieldUpdatedAt, v))
}

// RAMEQ applies the EQ predicate on the "ram" field.
func RAMEQ(v uint) predicate.Server {
	return predicate.Server(sql.FieldEQ(FieldRAM, v))
}

// RAMNEQ applies the NEQ predicate on the "ram" field.
func RAMNEQ(v uint) predicate.Server {
	return predicate.Server(sql.FieldNEQ(FieldRAM, v))
}

// RAMIn applies the In predicate on the "ram" field.
func RAMIn(vs ...uint) predicate.Server {
	return predicate.Server(sql.FieldIn(FieldRAM, vs...))
}

// RAMNotIn applies the NotIn predicate on the "ram" field.
func RAMNotIn(vs ...uint) predicate.Server {
	return predicate.Server(sql.FieldNotIn(FieldRAM, vs...))
}

// RAMGT applies the GT predicate on the "ram" field.
func RAMGT(v uint) predicate.Server {
	return predicate.Server(sql.FieldGT(FieldRAM, v))
}

// RAMGTE applies the GTE predicate on the "ram" field.
func RAMGTE(v uint) predicate.Server {
	return predicate.Server(sql.FieldGTE(FieldRAM, v))
}

// RAMLT applies the LT predicate on the "ram" field.
func RAMLT(v uint) predicate.Server {
	return predicate.Server(sql.FieldLT(FieldRAM, v))
}

// RAMLTE applies the LTE predicate on the "ram" field.
func RAMLTE(v uint) predicate.Server {
	return predicate.Server(sql.FieldLTE(FieldRAM, v))
}

// StorageEQ applies the EQ predicate on the "storage" field.
func StorageEQ(v uint) predicate.Server {
	return predicate.Server(sql.FieldEQ(FieldStorage, v))
}

// StorageNEQ applies the NEQ predicate on the "storage" field.
func StorageNEQ(v uint) predicate.Server {
	return predicate.Server(sql.FieldNEQ(FieldStorage, v))
}

// StorageIn applies the In predicate on the "storage" field.
func StorageIn(vs ...uint) predicate.Server {
	return predicate.Server(sql.FieldIn(FieldStorage, vs...))
}

// StorageNotIn applies the NotIn predicate on the "storage" field.
func StorageNotIn(vs ...uint) predicate.Server {
	return predicate.Server(sql.FieldNotIn(FieldStorage, vs...))
}

// StorageGT applies the GT predicate on the "storage" field.
func StorageGT(v uint) predicate.Server {
	return predicate.Server(sql.FieldGT(FieldStorage, v))
}

// StorageGTE applies the GTE predicate on the "storage" field.
func StorageGTE(v uint) predicate.Server {
	return predicate.Server(sql.FieldGTE(FieldStorage, v))
}

// StorageLT applies the LT predicate on the "storage" field.
func StorageLT(v uint) predicate.Server {
	return predicate.Server(sql.FieldLT(FieldStorage, v))
}

// StorageLTE applies the LTE predicate on the "storage" field.
func StorageLTE(v uint) predicate.Server {
	return predicate.Server(sql.FieldLTE(FieldStorage, v))
}

// LogicalCoresEQ applies the EQ predicate on the "logical_cores" field.
func LogicalCoresEQ(v uint) predicate.Server {
	return predicate.Server(sql.FieldEQ(FieldLogicalCores, v))
}

// LogicalCoresNEQ applies the NEQ predicate on the "logical_cores" field.
func LogicalCoresNEQ(v uint) predicate.Server {
	return predicate.Server(sql.FieldNEQ(FieldLogicalCores, v))
}

// LogicalCoresIn applies the In predicate on the "logical_cores" field.
func LogicalCoresIn(vs ...uint) predicate.Server {
	return predicate.Server(sql.FieldIn(FieldLogicalCores, vs...))
}

// LogicalCoresNotIn applies the NotIn predicate on the "logical_cores" field.
func LogicalCoresNotIn(vs ...uint) predicate.Server {
	return predicate.Server(sql.FieldNotIn(FieldLogicalCores, vs...))
}

// LogicalCoresGT applies the GT predicate on the "logical_cores" field.
func LogicalCoresGT(v uint) predicate.Server {
	return predicate.Server(sql.FieldGT(FieldLogicalCores, v))
}

// LogicalCoresGTE applies the GTE predicate on the "logical_cores" field.
func LogicalCoresGTE(v uint) predicate.Server {
	return predicate.Server(sql.FieldGTE(FieldLogicalCores, v))
}

// LogicalCoresLT applies the LT predicate on the "logical_cores" field.
func LogicalCoresLT(v uint) predicate.Server {
	return predicate.Server(sql.FieldLT(FieldLogicalCores, v))
}

// LogicalCoresLTE applies the LTE predicate on the "logical_cores" field.
func LogicalCoresLTE(v uint) predicate.Server {
	return predicate.Server(sql.FieldLTE(FieldLogicalCores, v))
}

// PortEQ applies the EQ predicate on the "port" field.
func PortEQ(v uint16) predicate.Server {
	return predicate.Server(sql.FieldEQ(FieldPort, v))
}

// PortNEQ applies the NEQ predicate on the "port" field.
func PortNEQ(v uint16) predicate.Server {
	return predicate.Server(sql.FieldNEQ(FieldPort, v))
}

// PortIn applies the In predicate on the "port" field.
func PortIn(vs ...uint16) predicate.Server {
	return predicate.Server(sql.FieldIn(FieldPort, vs...))
}

// PortNotIn applies the NotIn predicate on the "port" field.
func PortNotIn(vs ...uint16) predicate.Server {
	return predicate.Server(sql.FieldNotIn(FieldPort, vs...))
}

// PortGT applies the GT predicate on the "port" field.
func PortGT(v uint16) predicate.Server {
	return predicate.Server(sql.FieldGT(FieldPort, v))
}

// PortGTE applies the GTE predicate on the "port" field.
func PortGTE(v uint16) predicate.Server {
	return predicate.Server(sql.FieldGTE(FieldPort, v))
}

// PortLT applies the LT predicate on the "port" field.
func PortLT(v uint16) predicate.Server {
	return predicate.Server(sql.FieldLT(FieldPort, v))
}

// PortLTE applies the LTE predicate on the "port" field.
func PortLTE(v uint16) predicate.Server {
	return predicate.Server(sql.FieldLTE(FieldPort, v))
}

// HasNode applies the HasEdge predicate on the "node" edge.
func HasNode() predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, NodeTable, NodeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNodeWith applies the HasEdge predicate on the "node" edge with a given conditions (other predicates).
func HasNodeWith(preds ...predicate.Node) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		step := newNodeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGame applies the HasEdge predicate on the "game" edge.
func HasGame() predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GameTable, GameColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGameWith applies the HasEdge predicate on the "game" edge with a given conditions (other predicates).
func HasGameWith(preds ...predicate.Game) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		step := newGameStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Server) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Server) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
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
func Not(p predicate.Server) predicate.Server {
	return predicate.Server(func(s *sql.Selector) {
		p(s.Not())
	})
}
