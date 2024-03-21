// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Encedeus/panel/ent/node"
	"github.com/Encedeus/panel/ent/predicate"
	"github.com/Encedeus/panel/ent/server"
	"github.com/google/uuid"
)

// NodeUpdate is the builder for updating Node entities.
type NodeUpdate struct {
	config
	hooks    []Hook
	mutation *NodeMutation
}

// Where appends a list predicates to the NodeUpdate builder.
func (nu *NodeUpdate) Where(ps ...predicate.Node) *NodeUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetCreatedAt sets the "created_at" field.
func (nu *NodeUpdate) SetCreatedAt(t time.Time) *NodeUpdate {
	nu.mutation.SetCreatedAt(t)
	return nu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nu *NodeUpdate) SetNillableCreatedAt(t *time.Time) *NodeUpdate {
	if t != nil {
		nu.SetCreatedAt(*t)
	}
	return nu
}

// SetUpdatedAt sets the "updated_at" field.
func (nu *NodeUpdate) SetUpdatedAt(t time.Time) *NodeUpdate {
	nu.mutation.SetUpdatedAt(t)
	return nu
}

// SetIpv4Address sets the "ipv4_address" field.
func (nu *NodeUpdate) SetIpv4Address(s string) *NodeUpdate {
	nu.mutation.SetIpv4Address(s)
	return nu
}

// SetFqdn sets the "fqdn" field.
func (nu *NodeUpdate) SetFqdn(s string) *NodeUpdate {
	nu.mutation.SetFqdn(s)
	return nu
}

// SetSkyhookVersion sets the "skyhook_version" field.
func (nu *NodeUpdate) SetSkyhookVersion(s string) *NodeUpdate {
	nu.mutation.SetSkyhookVersion(s)
	return nu
}

// SetNillableSkyhookVersion sets the "skyhook_version" field if the given value is not nil.
func (nu *NodeUpdate) SetNillableSkyhookVersion(s *string) *NodeUpdate {
	if s != nil {
		nu.SetSkyhookVersion(*s)
	}
	return nu
}

// ClearSkyhookVersion clears the value of the "skyhook_version" field.
func (nu *NodeUpdate) ClearSkyhookVersion() *NodeUpdate {
	nu.mutation.ClearSkyhookVersion()
	return nu
}

// SetSkyhookAPIKey sets the "skyhook_api_key" field.
func (nu *NodeUpdate) SetSkyhookAPIKey(s string) *NodeUpdate {
	nu.mutation.SetSkyhookAPIKey(s)
	return nu
}

// SetOs sets the "os" field.
func (nu *NodeUpdate) SetOs(s string) *NodeUpdate {
	nu.mutation.SetOs(s)
	return nu
}

// SetNillableOs sets the "os" field if the given value is not nil.
func (nu *NodeUpdate) SetNillableOs(s *string) *NodeUpdate {
	if s != nil {
		nu.SetOs(*s)
	}
	return nu
}

// ClearOs clears the value of the "os" field.
func (nu *NodeUpdate) ClearOs() *NodeUpdate {
	nu.mutation.ClearOs()
	return nu
}

// SetCPU sets the "cpu" field.
func (nu *NodeUpdate) SetCPU(s string) *NodeUpdate {
	nu.mutation.SetCPU(s)
	return nu
}

// SetNillableCPU sets the "cpu" field if the given value is not nil.
func (nu *NodeUpdate) SetNillableCPU(s *string) *NodeUpdate {
	if s != nil {
		nu.SetCPU(*s)
	}
	return nu
}

// ClearCPU clears the value of the "cpu" field.
func (nu *NodeUpdate) ClearCPU() *NodeUpdate {
	nu.mutation.ClearCPU()
	return nu
}

// SetCPUBaseClock sets the "cpu_base_clock" field.
func (nu *NodeUpdate) SetCPUBaseClock(u uint) *NodeUpdate {
	nu.mutation.ResetCPUBaseClock()
	nu.mutation.SetCPUBaseClock(u)
	return nu
}

// SetNillableCPUBaseClock sets the "cpu_base_clock" field if the given value is not nil.
func (nu *NodeUpdate) SetNillableCPUBaseClock(u *uint) *NodeUpdate {
	if u != nil {
		nu.SetCPUBaseClock(*u)
	}
	return nu
}

// AddCPUBaseClock adds u to the "cpu_base_clock" field.
func (nu *NodeUpdate) AddCPUBaseClock(u int) *NodeUpdate {
	nu.mutation.AddCPUBaseClock(u)
	return nu
}

// ClearCPUBaseClock clears the value of the "cpu_base_clock" field.
func (nu *NodeUpdate) ClearCPUBaseClock() *NodeUpdate {
	nu.mutation.ClearCPUBaseClock()
	return nu
}

// SetCores sets the "cores" field.
func (nu *NodeUpdate) SetCores(u uint) *NodeUpdate {
	nu.mutation.ResetCores()
	nu.mutation.SetCores(u)
	return nu
}

// SetNillableCores sets the "cores" field if the given value is not nil.
func (nu *NodeUpdate) SetNillableCores(u *uint) *NodeUpdate {
	if u != nil {
		nu.SetCores(*u)
	}
	return nu
}

// AddCores adds u to the "cores" field.
func (nu *NodeUpdate) AddCores(u int) *NodeUpdate {
	nu.mutation.AddCores(u)
	return nu
}

// ClearCores clears the value of the "cores" field.
func (nu *NodeUpdate) ClearCores() *NodeUpdate {
	nu.mutation.ClearCores()
	return nu
}

// SetLogicalCores sets the "logical_cores" field.
func (nu *NodeUpdate) SetLogicalCores(u uint) *NodeUpdate {
	nu.mutation.ResetLogicalCores()
	nu.mutation.SetLogicalCores(u)
	return nu
}

// SetNillableLogicalCores sets the "logical_cores" field if the given value is not nil.
func (nu *NodeUpdate) SetNillableLogicalCores(u *uint) *NodeUpdate {
	if u != nil {
		nu.SetLogicalCores(*u)
	}
	return nu
}

// AddLogicalCores adds u to the "logical_cores" field.
func (nu *NodeUpdate) AddLogicalCores(u int) *NodeUpdate {
	nu.mutation.AddLogicalCores(u)
	return nu
}

// ClearLogicalCores clears the value of the "logical_cores" field.
func (nu *NodeUpdate) ClearLogicalCores() *NodeUpdate {
	nu.mutation.ClearLogicalCores()
	return nu
}

// SetRAM sets the "ram" field.
func (nu *NodeUpdate) SetRAM(u uint64) *NodeUpdate {
	nu.mutation.ResetRAM()
	nu.mutation.SetRAM(u)
	return nu
}

// SetNillableRAM sets the "ram" field if the given value is not nil.
func (nu *NodeUpdate) SetNillableRAM(u *uint64) *NodeUpdate {
	if u != nil {
		nu.SetRAM(*u)
	}
	return nu
}

// AddRAM adds u to the "ram" field.
func (nu *NodeUpdate) AddRAM(u int64) *NodeUpdate {
	nu.mutation.AddRAM(u)
	return nu
}

// ClearRAM clears the value of the "ram" field.
func (nu *NodeUpdate) ClearRAM() *NodeUpdate {
	nu.mutation.ClearRAM()
	return nu
}

// SetStorage sets the "storage" field.
func (nu *NodeUpdate) SetStorage(u uint64) *NodeUpdate {
	nu.mutation.ResetStorage()
	nu.mutation.SetStorage(u)
	return nu
}

// SetNillableStorage sets the "storage" field if the given value is not nil.
func (nu *NodeUpdate) SetNillableStorage(u *uint64) *NodeUpdate {
	if u != nil {
		nu.SetStorage(*u)
	}
	return nu
}

// AddStorage adds u to the "storage" field.
func (nu *NodeUpdate) AddStorage(u int64) *NodeUpdate {
	nu.mutation.AddStorage(u)
	return nu
}

// ClearStorage clears the value of the "storage" field.
func (nu *NodeUpdate) ClearStorage() *NodeUpdate {
	nu.mutation.ClearStorage()
	return nu
}

// AddNodeIDs adds the "nodes" edge to the Server entity by IDs.
func (nu *NodeUpdate) AddNodeIDs(ids ...uuid.UUID) *NodeUpdate {
	nu.mutation.AddNodeIDs(ids...)
	return nu
}

// AddNodes adds the "nodes" edges to the Server entity.
func (nu *NodeUpdate) AddNodes(s ...*Server) *NodeUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return nu.AddNodeIDs(ids...)
}

// Mutation returns the NodeMutation object of the builder.
func (nu *NodeUpdate) Mutation() *NodeMutation {
	return nu.mutation
}

// ClearNodes clears all "nodes" edges to the Server entity.
func (nu *NodeUpdate) ClearNodes() *NodeUpdate {
	nu.mutation.ClearNodes()
	return nu
}

// RemoveNodeIDs removes the "nodes" edge to Server entities by IDs.
func (nu *NodeUpdate) RemoveNodeIDs(ids ...uuid.UUID) *NodeUpdate {
	nu.mutation.RemoveNodeIDs(ids...)
	return nu
}

// RemoveNodes removes "nodes" edges to Server entities.
func (nu *NodeUpdate) RemoveNodes(s ...*Server) *NodeUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return nu.RemoveNodeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NodeUpdate) Save(ctx context.Context) (int, error) {
	nu.defaults()
	return withHooks(ctx, nu.sqlSave, nu.mutation, nu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NodeUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NodeUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NodeUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nu *NodeUpdate) defaults() {
	if _, ok := nu.mutation.UpdatedAt(); !ok {
		v := node.UpdateDefaultUpdatedAt()
		nu.mutation.SetUpdatedAt(v)
	}
}

func (nu *NodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(node.Table, node.Columns, sqlgraph.NewFieldSpec(node.FieldID, field.TypeUUID))
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.CreatedAt(); ok {
		_spec.SetField(node.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := nu.mutation.UpdatedAt(); ok {
		_spec.SetField(node.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := nu.mutation.Ipv4Address(); ok {
		_spec.SetField(node.FieldIpv4Address, field.TypeString, value)
	}
	if value, ok := nu.mutation.Fqdn(); ok {
		_spec.SetField(node.FieldFqdn, field.TypeString, value)
	}
	if value, ok := nu.mutation.SkyhookVersion(); ok {
		_spec.SetField(node.FieldSkyhookVersion, field.TypeString, value)
	}
	if nu.mutation.SkyhookVersionCleared() {
		_spec.ClearField(node.FieldSkyhookVersion, field.TypeString)
	}
	if value, ok := nu.mutation.SkyhookAPIKey(); ok {
		_spec.SetField(node.FieldSkyhookAPIKey, field.TypeString, value)
	}
	if value, ok := nu.mutation.Os(); ok {
		_spec.SetField(node.FieldOs, field.TypeString, value)
	}
	if nu.mutation.OsCleared() {
		_spec.ClearField(node.FieldOs, field.TypeString)
	}
	if value, ok := nu.mutation.CPU(); ok {
		_spec.SetField(node.FieldCPU, field.TypeString, value)
	}
	if nu.mutation.CPUCleared() {
		_spec.ClearField(node.FieldCPU, field.TypeString)
	}
	if value, ok := nu.mutation.CPUBaseClock(); ok {
		_spec.SetField(node.FieldCPUBaseClock, field.TypeUint, value)
	}
	if value, ok := nu.mutation.AddedCPUBaseClock(); ok {
		_spec.AddField(node.FieldCPUBaseClock, field.TypeUint, value)
	}
	if nu.mutation.CPUBaseClockCleared() {
		_spec.ClearField(node.FieldCPUBaseClock, field.TypeUint)
	}
	if value, ok := nu.mutation.Cores(); ok {
		_spec.SetField(node.FieldCores, field.TypeUint, value)
	}
	if value, ok := nu.mutation.AddedCores(); ok {
		_spec.AddField(node.FieldCores, field.TypeUint, value)
	}
	if nu.mutation.CoresCleared() {
		_spec.ClearField(node.FieldCores, field.TypeUint)
	}
	if value, ok := nu.mutation.LogicalCores(); ok {
		_spec.SetField(node.FieldLogicalCores, field.TypeUint, value)
	}
	if value, ok := nu.mutation.AddedLogicalCores(); ok {
		_spec.AddField(node.FieldLogicalCores, field.TypeUint, value)
	}
	if nu.mutation.LogicalCoresCleared() {
		_spec.ClearField(node.FieldLogicalCores, field.TypeUint)
	}
	if value, ok := nu.mutation.RAM(); ok {
		_spec.SetField(node.FieldRAM, field.TypeUint64, value)
	}
	if value, ok := nu.mutation.AddedRAM(); ok {
		_spec.AddField(node.FieldRAM, field.TypeUint64, value)
	}
	if nu.mutation.RAMCleared() {
		_spec.ClearField(node.FieldRAM, field.TypeUint64)
	}
	if value, ok := nu.mutation.Storage(); ok {
		_spec.SetField(node.FieldStorage, field.TypeUint64, value)
	}
	if value, ok := nu.mutation.AddedStorage(); ok {
		_spec.AddField(node.FieldStorage, field.TypeUint64, value)
	}
	if nu.mutation.StorageCleared() {
		_spec.ClearField(node.FieldStorage, field.TypeUint64)
	}
	if nu.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   node.NodesTable,
			Columns: []string{node.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.RemovedNodesIDs(); len(nodes) > 0 && !nu.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   node.NodesTable,
			Columns: []string{node.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nu.mutation.NodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   node.NodesTable,
			Columns: []string{node.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{node.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nu.mutation.done = true
	return n, nil
}

// NodeUpdateOne is the builder for updating a single Node entity.
type NodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NodeMutation
}

// SetCreatedAt sets the "created_at" field.
func (nuo *NodeUpdateOne) SetCreatedAt(t time.Time) *NodeUpdateOne {
	nuo.mutation.SetCreatedAt(t)
	return nuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (nuo *NodeUpdateOne) SetNillableCreatedAt(t *time.Time) *NodeUpdateOne {
	if t != nil {
		nuo.SetCreatedAt(*t)
	}
	return nuo
}

// SetUpdatedAt sets the "updated_at" field.
func (nuo *NodeUpdateOne) SetUpdatedAt(t time.Time) *NodeUpdateOne {
	nuo.mutation.SetUpdatedAt(t)
	return nuo
}

// SetIpv4Address sets the "ipv4_address" field.
func (nuo *NodeUpdateOne) SetIpv4Address(s string) *NodeUpdateOne {
	nuo.mutation.SetIpv4Address(s)
	return nuo
}

// SetFqdn sets the "fqdn" field.
func (nuo *NodeUpdateOne) SetFqdn(s string) *NodeUpdateOne {
	nuo.mutation.SetFqdn(s)
	return nuo
}

// SetSkyhookVersion sets the "skyhook_version" field.
func (nuo *NodeUpdateOne) SetSkyhookVersion(s string) *NodeUpdateOne {
	nuo.mutation.SetSkyhookVersion(s)
	return nuo
}

// SetNillableSkyhookVersion sets the "skyhook_version" field if the given value is not nil.
func (nuo *NodeUpdateOne) SetNillableSkyhookVersion(s *string) *NodeUpdateOne {
	if s != nil {
		nuo.SetSkyhookVersion(*s)
	}
	return nuo
}

// ClearSkyhookVersion clears the value of the "skyhook_version" field.
func (nuo *NodeUpdateOne) ClearSkyhookVersion() *NodeUpdateOne {
	nuo.mutation.ClearSkyhookVersion()
	return nuo
}

// SetSkyhookAPIKey sets the "skyhook_api_key" field.
func (nuo *NodeUpdateOne) SetSkyhookAPIKey(s string) *NodeUpdateOne {
	nuo.mutation.SetSkyhookAPIKey(s)
	return nuo
}

// SetOs sets the "os" field.
func (nuo *NodeUpdateOne) SetOs(s string) *NodeUpdateOne {
	nuo.mutation.SetOs(s)
	return nuo
}

// SetNillableOs sets the "os" field if the given value is not nil.
func (nuo *NodeUpdateOne) SetNillableOs(s *string) *NodeUpdateOne {
	if s != nil {
		nuo.SetOs(*s)
	}
	return nuo
}

// ClearOs clears the value of the "os" field.
func (nuo *NodeUpdateOne) ClearOs() *NodeUpdateOne {
	nuo.mutation.ClearOs()
	return nuo
}

// SetCPU sets the "cpu" field.
func (nuo *NodeUpdateOne) SetCPU(s string) *NodeUpdateOne {
	nuo.mutation.SetCPU(s)
	return nuo
}

// SetNillableCPU sets the "cpu" field if the given value is not nil.
func (nuo *NodeUpdateOne) SetNillableCPU(s *string) *NodeUpdateOne {
	if s != nil {
		nuo.SetCPU(*s)
	}
	return nuo
}

// ClearCPU clears the value of the "cpu" field.
func (nuo *NodeUpdateOne) ClearCPU() *NodeUpdateOne {
	nuo.mutation.ClearCPU()
	return nuo
}

// SetCPUBaseClock sets the "cpu_base_clock" field.
func (nuo *NodeUpdateOne) SetCPUBaseClock(u uint) *NodeUpdateOne {
	nuo.mutation.ResetCPUBaseClock()
	nuo.mutation.SetCPUBaseClock(u)
	return nuo
}

// SetNillableCPUBaseClock sets the "cpu_base_clock" field if the given value is not nil.
func (nuo *NodeUpdateOne) SetNillableCPUBaseClock(u *uint) *NodeUpdateOne {
	if u != nil {
		nuo.SetCPUBaseClock(*u)
	}
	return nuo
}

// AddCPUBaseClock adds u to the "cpu_base_clock" field.
func (nuo *NodeUpdateOne) AddCPUBaseClock(u int) *NodeUpdateOne {
	nuo.mutation.AddCPUBaseClock(u)
	return nuo
}

// ClearCPUBaseClock clears the value of the "cpu_base_clock" field.
func (nuo *NodeUpdateOne) ClearCPUBaseClock() *NodeUpdateOne {
	nuo.mutation.ClearCPUBaseClock()
	return nuo
}

// SetCores sets the "cores" field.
func (nuo *NodeUpdateOne) SetCores(u uint) *NodeUpdateOne {
	nuo.mutation.ResetCores()
	nuo.mutation.SetCores(u)
	return nuo
}

// SetNillableCores sets the "cores" field if the given value is not nil.
func (nuo *NodeUpdateOne) SetNillableCores(u *uint) *NodeUpdateOne {
	if u != nil {
		nuo.SetCores(*u)
	}
	return nuo
}

// AddCores adds u to the "cores" field.
func (nuo *NodeUpdateOne) AddCores(u int) *NodeUpdateOne {
	nuo.mutation.AddCores(u)
	return nuo
}

// ClearCores clears the value of the "cores" field.
func (nuo *NodeUpdateOne) ClearCores() *NodeUpdateOne {
	nuo.mutation.ClearCores()
	return nuo
}

// SetLogicalCores sets the "logical_cores" field.
func (nuo *NodeUpdateOne) SetLogicalCores(u uint) *NodeUpdateOne {
	nuo.mutation.ResetLogicalCores()
	nuo.mutation.SetLogicalCores(u)
	return nuo
}

// SetNillableLogicalCores sets the "logical_cores" field if the given value is not nil.
func (nuo *NodeUpdateOne) SetNillableLogicalCores(u *uint) *NodeUpdateOne {
	if u != nil {
		nuo.SetLogicalCores(*u)
	}
	return nuo
}

// AddLogicalCores adds u to the "logical_cores" field.
func (nuo *NodeUpdateOne) AddLogicalCores(u int) *NodeUpdateOne {
	nuo.mutation.AddLogicalCores(u)
	return nuo
}

// ClearLogicalCores clears the value of the "logical_cores" field.
func (nuo *NodeUpdateOne) ClearLogicalCores() *NodeUpdateOne {
	nuo.mutation.ClearLogicalCores()
	return nuo
}

// SetRAM sets the "ram" field.
func (nuo *NodeUpdateOne) SetRAM(u uint64) *NodeUpdateOne {
	nuo.mutation.ResetRAM()
	nuo.mutation.SetRAM(u)
	return nuo
}

// SetNillableRAM sets the "ram" field if the given value is not nil.
func (nuo *NodeUpdateOne) SetNillableRAM(u *uint64) *NodeUpdateOne {
	if u != nil {
		nuo.SetRAM(*u)
	}
	return nuo
}

// AddRAM adds u to the "ram" field.
func (nuo *NodeUpdateOne) AddRAM(u int64) *NodeUpdateOne {
	nuo.mutation.AddRAM(u)
	return nuo
}

// ClearRAM clears the value of the "ram" field.
func (nuo *NodeUpdateOne) ClearRAM() *NodeUpdateOne {
	nuo.mutation.ClearRAM()
	return nuo
}

// SetStorage sets the "storage" field.
func (nuo *NodeUpdateOne) SetStorage(u uint64) *NodeUpdateOne {
	nuo.mutation.ResetStorage()
	nuo.mutation.SetStorage(u)
	return nuo
}

// SetNillableStorage sets the "storage" field if the given value is not nil.
func (nuo *NodeUpdateOne) SetNillableStorage(u *uint64) *NodeUpdateOne {
	if u != nil {
		nuo.SetStorage(*u)
	}
	return nuo
}

// AddStorage adds u to the "storage" field.
func (nuo *NodeUpdateOne) AddStorage(u int64) *NodeUpdateOne {
	nuo.mutation.AddStorage(u)
	return nuo
}

// ClearStorage clears the value of the "storage" field.
func (nuo *NodeUpdateOne) ClearStorage() *NodeUpdateOne {
	nuo.mutation.ClearStorage()
	return nuo
}

// AddNodeIDs adds the "nodes" edge to the Server entity by IDs.
func (nuo *NodeUpdateOne) AddNodeIDs(ids ...uuid.UUID) *NodeUpdateOne {
	nuo.mutation.AddNodeIDs(ids...)
	return nuo
}

// AddNodes adds the "nodes" edges to the Server entity.
func (nuo *NodeUpdateOne) AddNodes(s ...*Server) *NodeUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return nuo.AddNodeIDs(ids...)
}

// Mutation returns the NodeMutation object of the builder.
func (nuo *NodeUpdateOne) Mutation() *NodeMutation {
	return nuo.mutation
}

// ClearNodes clears all "nodes" edges to the Server entity.
func (nuo *NodeUpdateOne) ClearNodes() *NodeUpdateOne {
	nuo.mutation.ClearNodes()
	return nuo
}

// RemoveNodeIDs removes the "nodes" edge to Server entities by IDs.
func (nuo *NodeUpdateOne) RemoveNodeIDs(ids ...uuid.UUID) *NodeUpdateOne {
	nuo.mutation.RemoveNodeIDs(ids...)
	return nuo
}

// RemoveNodes removes "nodes" edges to Server entities.
func (nuo *NodeUpdateOne) RemoveNodes(s ...*Server) *NodeUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return nuo.RemoveNodeIDs(ids...)
}

// Where appends a list predicates to the NodeUpdate builder.
func (nuo *NodeUpdateOne) Where(ps ...predicate.Node) *NodeUpdateOne {
	nuo.mutation.Where(ps...)
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NodeUpdateOne) Select(field string, fields ...string) *NodeUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Node entity.
func (nuo *NodeUpdateOne) Save(ctx context.Context) (*Node, error) {
	nuo.defaults()
	return withHooks(ctx, nuo.sqlSave, nuo.mutation, nuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NodeUpdateOne) SaveX(ctx context.Context) *Node {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NodeUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NodeUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nuo *NodeUpdateOne) defaults() {
	if _, ok := nuo.mutation.UpdatedAt(); !ok {
		v := node.UpdateDefaultUpdatedAt()
		nuo.mutation.SetUpdatedAt(v)
	}
}

func (nuo *NodeUpdateOne) sqlSave(ctx context.Context) (_node *Node, err error) {
	_spec := sqlgraph.NewUpdateSpec(node.Table, node.Columns, sqlgraph.NewFieldSpec(node.FieldID, field.TypeUUID))
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Node.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, node.FieldID)
		for _, f := range fields {
			if !node.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != node.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.CreatedAt(); ok {
		_spec.SetField(node.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := nuo.mutation.UpdatedAt(); ok {
		_spec.SetField(node.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := nuo.mutation.Ipv4Address(); ok {
		_spec.SetField(node.FieldIpv4Address, field.TypeString, value)
	}
	if value, ok := nuo.mutation.Fqdn(); ok {
		_spec.SetField(node.FieldFqdn, field.TypeString, value)
	}
	if value, ok := nuo.mutation.SkyhookVersion(); ok {
		_spec.SetField(node.FieldSkyhookVersion, field.TypeString, value)
	}
	if nuo.mutation.SkyhookVersionCleared() {
		_spec.ClearField(node.FieldSkyhookVersion, field.TypeString)
	}
	if value, ok := nuo.mutation.SkyhookAPIKey(); ok {
		_spec.SetField(node.FieldSkyhookAPIKey, field.TypeString, value)
	}
	if value, ok := nuo.mutation.Os(); ok {
		_spec.SetField(node.FieldOs, field.TypeString, value)
	}
	if nuo.mutation.OsCleared() {
		_spec.ClearField(node.FieldOs, field.TypeString)
	}
	if value, ok := nuo.mutation.CPU(); ok {
		_spec.SetField(node.FieldCPU, field.TypeString, value)
	}
	if nuo.mutation.CPUCleared() {
		_spec.ClearField(node.FieldCPU, field.TypeString)
	}
	if value, ok := nuo.mutation.CPUBaseClock(); ok {
		_spec.SetField(node.FieldCPUBaseClock, field.TypeUint, value)
	}
	if value, ok := nuo.mutation.AddedCPUBaseClock(); ok {
		_spec.AddField(node.FieldCPUBaseClock, field.TypeUint, value)
	}
	if nuo.mutation.CPUBaseClockCleared() {
		_spec.ClearField(node.FieldCPUBaseClock, field.TypeUint)
	}
	if value, ok := nuo.mutation.Cores(); ok {
		_spec.SetField(node.FieldCores, field.TypeUint, value)
	}
	if value, ok := nuo.mutation.AddedCores(); ok {
		_spec.AddField(node.FieldCores, field.TypeUint, value)
	}
	if nuo.mutation.CoresCleared() {
		_spec.ClearField(node.FieldCores, field.TypeUint)
	}
	if value, ok := nuo.mutation.LogicalCores(); ok {
		_spec.SetField(node.FieldLogicalCores, field.TypeUint, value)
	}
	if value, ok := nuo.mutation.AddedLogicalCores(); ok {
		_spec.AddField(node.FieldLogicalCores, field.TypeUint, value)
	}
	if nuo.mutation.LogicalCoresCleared() {
		_spec.ClearField(node.FieldLogicalCores, field.TypeUint)
	}
	if value, ok := nuo.mutation.RAM(); ok {
		_spec.SetField(node.FieldRAM, field.TypeUint64, value)
	}
	if value, ok := nuo.mutation.AddedRAM(); ok {
		_spec.AddField(node.FieldRAM, field.TypeUint64, value)
	}
	if nuo.mutation.RAMCleared() {
		_spec.ClearField(node.FieldRAM, field.TypeUint64)
	}
	if value, ok := nuo.mutation.Storage(); ok {
		_spec.SetField(node.FieldStorage, field.TypeUint64, value)
	}
	if value, ok := nuo.mutation.AddedStorage(); ok {
		_spec.AddField(node.FieldStorage, field.TypeUint64, value)
	}
	if nuo.mutation.StorageCleared() {
		_spec.ClearField(node.FieldStorage, field.TypeUint64)
	}
	if nuo.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   node.NodesTable,
			Columns: []string{node.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.RemovedNodesIDs(); len(nodes) > 0 && !nuo.mutation.NodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   node.NodesTable,
			Columns: []string{node.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := nuo.mutation.NodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   node.NodesTable,
			Columns: []string{node.NodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(server.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Node{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{node.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nuo.mutation.done = true
	return _node, nil
}
