// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/sergot/tibiago/ent/instance"
	"github.com/sergot/tibiago/ent/instanceconfig"
	"github.com/sergot/tibiago/ent/predicate"
)

// InstanceConfigUpdate is the builder for updating InstanceConfig entities.
type InstanceConfigUpdate struct {
	config
	hooks    []Hook
	mutation *InstanceConfigMutation
}

// Where appends a list predicates to the InstanceConfigUpdate builder.
func (icu *InstanceConfigUpdate) Where(ps ...predicate.InstanceConfig) *InstanceConfigUpdate {
	icu.mutation.Where(ps...)
	return icu
}

// SetKey sets the "key" field.
func (icu *InstanceConfigUpdate) SetKey(s string) *InstanceConfigUpdate {
	icu.mutation.SetKey(s)
	return icu
}

// SetValue sets the "value" field.
func (icu *InstanceConfigUpdate) SetValue(s string) *InstanceConfigUpdate {
	icu.mutation.SetValue(s)
	return icu
}

// SetInstanceID sets the "instance" edge to the Instance entity by ID.
func (icu *InstanceConfigUpdate) SetInstanceID(id uuid.UUID) *InstanceConfigUpdate {
	icu.mutation.SetInstanceID(id)
	return icu
}

// SetNillableInstanceID sets the "instance" edge to the Instance entity by ID if the given value is not nil.
func (icu *InstanceConfigUpdate) SetNillableInstanceID(id *uuid.UUID) *InstanceConfigUpdate {
	if id != nil {
		icu = icu.SetInstanceID(*id)
	}
	return icu
}

// SetInstance sets the "instance" edge to the Instance entity.
func (icu *InstanceConfigUpdate) SetInstance(i *Instance) *InstanceConfigUpdate {
	return icu.SetInstanceID(i.ID)
}

// Mutation returns the InstanceConfigMutation object of the builder.
func (icu *InstanceConfigUpdate) Mutation() *InstanceConfigMutation {
	return icu.mutation
}

// ClearInstance clears the "instance" edge to the Instance entity.
func (icu *InstanceConfigUpdate) ClearInstance() *InstanceConfigUpdate {
	icu.mutation.ClearInstance()
	return icu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (icu *InstanceConfigUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(icu.hooks) == 0 {
		affected, err = icu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InstanceConfigMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			icu.mutation = mutation
			affected, err = icu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(icu.hooks) - 1; i >= 0; i-- {
			if icu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = icu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, icu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (icu *InstanceConfigUpdate) SaveX(ctx context.Context) int {
	affected, err := icu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (icu *InstanceConfigUpdate) Exec(ctx context.Context) error {
	_, err := icu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icu *InstanceConfigUpdate) ExecX(ctx context.Context) {
	if err := icu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (icu *InstanceConfigUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   instanceconfig.Table,
			Columns: instanceconfig.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: instanceconfig.FieldID,
			},
		},
	}
	if ps := icu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := icu.mutation.Key(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instanceconfig.FieldKey,
		})
	}
	if value, ok := icu.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instanceconfig.FieldValue,
		})
	}
	if icu.mutation.InstanceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   instanceconfig.InstanceTable,
			Columns: []string{instanceconfig.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: instance.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := icu.mutation.InstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   instanceconfig.InstanceTable,
			Columns: []string{instanceconfig.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: instance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, icu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{instanceconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// InstanceConfigUpdateOne is the builder for updating a single InstanceConfig entity.
type InstanceConfigUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *InstanceConfigMutation
}

// SetKey sets the "key" field.
func (icuo *InstanceConfigUpdateOne) SetKey(s string) *InstanceConfigUpdateOne {
	icuo.mutation.SetKey(s)
	return icuo
}

// SetValue sets the "value" field.
func (icuo *InstanceConfigUpdateOne) SetValue(s string) *InstanceConfigUpdateOne {
	icuo.mutation.SetValue(s)
	return icuo
}

// SetInstanceID sets the "instance" edge to the Instance entity by ID.
func (icuo *InstanceConfigUpdateOne) SetInstanceID(id uuid.UUID) *InstanceConfigUpdateOne {
	icuo.mutation.SetInstanceID(id)
	return icuo
}

// SetNillableInstanceID sets the "instance" edge to the Instance entity by ID if the given value is not nil.
func (icuo *InstanceConfigUpdateOne) SetNillableInstanceID(id *uuid.UUID) *InstanceConfigUpdateOne {
	if id != nil {
		icuo = icuo.SetInstanceID(*id)
	}
	return icuo
}

// SetInstance sets the "instance" edge to the Instance entity.
func (icuo *InstanceConfigUpdateOne) SetInstance(i *Instance) *InstanceConfigUpdateOne {
	return icuo.SetInstanceID(i.ID)
}

// Mutation returns the InstanceConfigMutation object of the builder.
func (icuo *InstanceConfigUpdateOne) Mutation() *InstanceConfigMutation {
	return icuo.mutation
}

// ClearInstance clears the "instance" edge to the Instance entity.
func (icuo *InstanceConfigUpdateOne) ClearInstance() *InstanceConfigUpdateOne {
	icuo.mutation.ClearInstance()
	return icuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (icuo *InstanceConfigUpdateOne) Select(field string, fields ...string) *InstanceConfigUpdateOne {
	icuo.fields = append([]string{field}, fields...)
	return icuo
}

// Save executes the query and returns the updated InstanceConfig entity.
func (icuo *InstanceConfigUpdateOne) Save(ctx context.Context) (*InstanceConfig, error) {
	var (
		err  error
		node *InstanceConfig
	)
	if len(icuo.hooks) == 0 {
		node, err = icuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InstanceConfigMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			icuo.mutation = mutation
			node, err = icuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(icuo.hooks) - 1; i >= 0; i-- {
			if icuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = icuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, icuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*InstanceConfig)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from InstanceConfigMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (icuo *InstanceConfigUpdateOne) SaveX(ctx context.Context) *InstanceConfig {
	node, err := icuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (icuo *InstanceConfigUpdateOne) Exec(ctx context.Context) error {
	_, err := icuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icuo *InstanceConfigUpdateOne) ExecX(ctx context.Context) {
	if err := icuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (icuo *InstanceConfigUpdateOne) sqlSave(ctx context.Context) (_node *InstanceConfig, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   instanceconfig.Table,
			Columns: instanceconfig.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: instanceconfig.FieldID,
			},
		},
	}
	id, ok := icuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "InstanceConfig.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := icuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, instanceconfig.FieldID)
		for _, f := range fields {
			if !instanceconfig.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != instanceconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := icuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := icuo.mutation.Key(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instanceconfig.FieldKey,
		})
	}
	if value, ok := icuo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: instanceconfig.FieldValue,
		})
	}
	if icuo.mutation.InstanceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   instanceconfig.InstanceTable,
			Columns: []string{instanceconfig.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: instance.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := icuo.mutation.InstanceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   instanceconfig.InstanceTable,
			Columns: []string{instanceconfig.InstanceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: instance.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &InstanceConfig{config: icuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, icuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{instanceconfig.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
