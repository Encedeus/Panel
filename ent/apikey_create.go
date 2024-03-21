// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Encedeus/panel/ent/apikey"
	"github.com/Encedeus/panel/ent/user"
	"github.com/google/uuid"
)

// ApiKeyCreate is the builder for creating a ApiKey entity.
type ApiKeyCreate struct {
	config
	mutation *ApiKeyMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (akc *ApiKeyCreate) SetCreatedAt(t time.Time) *ApiKeyCreate {
	akc.mutation.SetCreatedAt(t)
	return akc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (akc *ApiKeyCreate) SetNillableCreatedAt(t *time.Time) *ApiKeyCreate {
	if t != nil {
		akc.SetCreatedAt(*t)
	}
	return akc
}

// SetUpdatedAt sets the "updated_at" field.
func (akc *ApiKeyCreate) SetUpdatedAt(t time.Time) *ApiKeyCreate {
	akc.mutation.SetUpdatedAt(t)
	return akc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (akc *ApiKeyCreate) SetNillableUpdatedAt(t *time.Time) *ApiKeyCreate {
	if t != nil {
		akc.SetUpdatedAt(*t)
	}
	return akc
}

// SetDescription sets the "description" field.
func (akc *ApiKeyCreate) SetDescription(s string) *ApiKeyCreate {
	akc.mutation.SetDescription(s)
	return akc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (akc *ApiKeyCreate) SetNillableDescription(s *string) *ApiKeyCreate {
	if s != nil {
		akc.SetDescription(*s)
	}
	return akc
}

// SetIPAddresses sets the "ip_addresses" field.
func (akc *ApiKeyCreate) SetIPAddresses(s []string) *ApiKeyCreate {
	akc.mutation.SetIPAddresses(s)
	return akc
}

// SetKey sets the "key" field.
func (akc *ApiKeyCreate) SetKey(s string) *ApiKeyCreate {
	akc.mutation.SetKey(s)
	return akc
}

// SetUserID sets the "user_id" field.
func (akc *ApiKeyCreate) SetUserID(u uuid.UUID) *ApiKeyCreate {
	akc.mutation.SetUserID(u)
	return akc
}

// SetID sets the "id" field.
func (akc *ApiKeyCreate) SetID(u uuid.UUID) *ApiKeyCreate {
	akc.mutation.SetID(u)
	return akc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (akc *ApiKeyCreate) SetNillableID(u *uuid.UUID) *ApiKeyCreate {
	if u != nil {
		akc.SetID(*u)
	}
	return akc
}

// SetUser sets the "user" edge to the User entity.
func (akc *ApiKeyCreate) SetUser(u *User) *ApiKeyCreate {
	return akc.SetUserID(u.ID)
}

// Mutation returns the ApiKeyMutation object of the builder.
func (akc *ApiKeyCreate) Mutation() *ApiKeyMutation {
	return akc.mutation
}

// Save creates the ApiKey in the database.
func (akc *ApiKeyCreate) Save(ctx context.Context) (*ApiKey, error) {
	akc.defaults()
	return withHooks(ctx, akc.sqlSave, akc.mutation, akc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (akc *ApiKeyCreate) SaveX(ctx context.Context) *ApiKey {
	v, err := akc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (akc *ApiKeyCreate) Exec(ctx context.Context) error {
	_, err := akc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (akc *ApiKeyCreate) ExecX(ctx context.Context) {
	if err := akc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (akc *ApiKeyCreate) defaults() {
	if _, ok := akc.mutation.CreatedAt(); !ok {
		v := apikey.DefaultCreatedAt()
		akc.mutation.SetCreatedAt(v)
	}
	if _, ok := akc.mutation.UpdatedAt(); !ok {
		v := apikey.DefaultUpdatedAt()
		akc.mutation.SetUpdatedAt(v)
	}
	if _, ok := akc.mutation.ID(); !ok {
		v := apikey.DefaultID()
		akc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (akc *ApiKeyCreate) check() error {
	if _, ok := akc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ApiKey.created_at"`)}
	}
	if _, ok := akc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ApiKey.updated_at"`)}
	}
	if _, ok := akc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required field "ApiKey.key"`)}
	}
	if v, ok := akc.mutation.Key(); ok {
		if err := apikey.KeyValidator(v); err != nil {
			return &ValidationError{Name: "key", err: fmt.Errorf(`ent: validator failed for field "ApiKey.key": %w`, err)}
		}
	}
	if _, ok := akc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "ApiKey.user_id"`)}
	}
	if _, ok := akc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "ApiKey.user"`)}
	}
	return nil
}

func (akc *ApiKeyCreate) sqlSave(ctx context.Context) (*ApiKey, error) {
	if err := akc.check(); err != nil {
		return nil, err
	}
	_node, _spec := akc.createSpec()
	if err := sqlgraph.CreateNode(ctx, akc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	akc.mutation.id = &_node.ID
	akc.mutation.done = true
	return _node, nil
}

func (akc *ApiKeyCreate) createSpec() (*ApiKey, *sqlgraph.CreateSpec) {
	var (
		_node = &ApiKey{config: akc.config}
		_spec = sqlgraph.NewCreateSpec(apikey.Table, sqlgraph.NewFieldSpec(apikey.FieldID, field.TypeUUID))
	)
	if id, ok := akc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := akc.mutation.CreatedAt(); ok {
		_spec.SetField(apikey.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := akc.mutation.UpdatedAt(); ok {
		_spec.SetField(apikey.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := akc.mutation.Description(); ok {
		_spec.SetField(apikey.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := akc.mutation.IPAddresses(); ok {
		_spec.SetField(apikey.FieldIPAddresses, field.TypeJSON, value)
		_node.IPAddresses = value
	}
	if value, ok := akc.mutation.Key(); ok {
		_spec.SetField(apikey.FieldKey, field.TypeString, value)
		_node.Key = value
	}
	if nodes := akc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   apikey.UserTable,
			Columns: []string{apikey.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ApiKeyCreateBulk is the builder for creating many ApiKey entities in bulk.
type ApiKeyCreateBulk struct {
	config
	builders []*ApiKeyCreate
}

// Save creates the ApiKey entities in the database.
func (akcb *ApiKeyCreateBulk) Save(ctx context.Context) ([]*ApiKey, error) {
	specs := make([]*sqlgraph.CreateSpec, len(akcb.builders))
	nodes := make([]*ApiKey, len(akcb.builders))
	mutators := make([]Mutator, len(akcb.builders))
	for i := range akcb.builders {
		func(i int, root context.Context) {
			builder := akcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ApiKeyMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, akcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, akcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, akcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (akcb *ApiKeyCreateBulk) SaveX(ctx context.Context) []*ApiKey {
	v, err := akcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (akcb *ApiKeyCreateBulk) Exec(ctx context.Context) error {
	_, err := akcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (akcb *ApiKeyCreateBulk) ExecX(ctx context.Context) {
	if err := akcb.Exec(ctx); err != nil {
		panic(err)
	}
}
