// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lbrictson/status/ent/operator"
)

// OperatorCreate is the builder for creating a Operator entity.
type OperatorCreate struct {
	config
	mutation *OperatorMutation
	hooks    []Hook
}

// SetEmail sets the "Email" field.
func (oc *OperatorCreate) SetEmail(s string) *OperatorCreate {
	oc.mutation.SetEmail(s)
	return oc
}

// SetHashedPassword sets the "HashedPassword" field.
func (oc *OperatorCreate) SetHashedPassword(s string) *OperatorCreate {
	oc.mutation.SetHashedPassword(s)
	return oc
}

// SetRole sets the "Role" field.
func (oc *OperatorCreate) SetRole(s string) *OperatorCreate {
	oc.mutation.SetRole(s)
	return oc
}

// Mutation returns the OperatorMutation object of the builder.
func (oc *OperatorCreate) Mutation() *OperatorMutation {
	return oc.mutation
}

// Save creates the Operator in the database.
func (oc *OperatorCreate) Save(ctx context.Context) (*Operator, error) {
	var (
		err  error
		node *Operator
	)
	if len(oc.hooks) == 0 {
		if err = oc.check(); err != nil {
			return nil, err
		}
		node, err = oc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OperatorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = oc.check(); err != nil {
				return nil, err
			}
			oc.mutation = mutation
			if node, err = oc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(oc.hooks) - 1; i >= 0; i-- {
			if oc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = oc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OperatorCreate) SaveX(ctx context.Context) *Operator {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OperatorCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OperatorCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OperatorCreate) check() error {
	if _, ok := oc.mutation.Email(); !ok {
		return &ValidationError{Name: "Email", err: errors.New(`ent: missing required field "Operator.Email"`)}
	}
	if _, ok := oc.mutation.HashedPassword(); !ok {
		return &ValidationError{Name: "HashedPassword", err: errors.New(`ent: missing required field "Operator.HashedPassword"`)}
	}
	if _, ok := oc.mutation.Role(); !ok {
		return &ValidationError{Name: "Role", err: errors.New(`ent: missing required field "Operator.Role"`)}
	}
	return nil
}

func (oc *OperatorCreate) sqlSave(ctx context.Context) (*Operator, error) {
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (oc *OperatorCreate) createSpec() (*Operator, *sqlgraph.CreateSpec) {
	var (
		_node = &Operator{config: oc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: operator.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: operator.FieldID,
			},
		}
	)
	if value, ok := oc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: operator.FieldEmail,
		})
		_node.Email = value
	}
	if value, ok := oc.mutation.HashedPassword(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: operator.FieldHashedPassword,
		})
		_node.HashedPassword = value
	}
	if value, ok := oc.mutation.Role(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: operator.FieldRole,
		})
		_node.Role = value
	}
	return _node, _spec
}

// OperatorCreateBulk is the builder for creating many Operator entities in bulk.
type OperatorCreateBulk struct {
	config
	builders []*OperatorCreate
}

// Save creates the Operator entities in the database.
func (ocb *OperatorCreateBulk) Save(ctx context.Context) ([]*Operator, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Operator, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OperatorMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OperatorCreateBulk) SaveX(ctx context.Context) []*Operator {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OperatorCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OperatorCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}