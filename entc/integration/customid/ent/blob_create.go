// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/blob"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// BlobCreate is the builder for creating a Blob entity.
type BlobCreate struct {
	config
	mutation *BlobMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetUUID sets the "uuid" field.
func (bc *BlobCreate) SetUUID(u uuid.UUID) *BlobCreate {
	bc.mutation.SetUUID(u)
	return bc
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (bc *BlobCreate) SetNillableUUID(u *uuid.UUID) *BlobCreate {
	if u != nil {
		bc.SetUUID(*u)
	}
	return bc
}

// SetCount sets the "count" field.
func (bc *BlobCreate) SetCount(i int) *BlobCreate {
	bc.mutation.SetCount(i)
	return bc
}

// SetNillableCount sets the "count" field if the given value is not nil.
func (bc *BlobCreate) SetNillableCount(i *int) *BlobCreate {
	if i != nil {
		bc.SetCount(*i)
	}
	return bc
}

// SetID sets the "id" field.
func (bc *BlobCreate) SetID(u uuid.UUID) *BlobCreate {
	bc.mutation.SetID(u)
	return bc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (bc *BlobCreate) SetNillableID(u *uuid.UUID) *BlobCreate {
	if u != nil {
		bc.SetID(*u)
	}
	return bc
}

// SetParentID sets the "parent" edge to the Blob entity by ID.
func (bc *BlobCreate) SetParentID(id uuid.UUID) *BlobCreate {
	bc.mutation.SetParentID(id)
	return bc
}

// SetNillableParentID sets the "parent" edge to the Blob entity by ID if the given value is not nil.
func (bc *BlobCreate) SetNillableParentID(id *uuid.UUID) *BlobCreate {
	if id != nil {
		bc = bc.SetParentID(*id)
	}
	return bc
}

// SetParent sets the "parent" edge to the Blob entity.
func (bc *BlobCreate) SetParent(b *Blob) *BlobCreate {
	return bc.SetParentID(b.ID)
}

// AddLinkIDs adds the "links" edge to the Blob entity by IDs.
func (bc *BlobCreate) AddLinkIDs(ids ...uuid.UUID) *BlobCreate {
	bc.mutation.AddLinkIDs(ids...)
	return bc
}

// AddLinks adds the "links" edges to the Blob entity.
func (bc *BlobCreate) AddLinks(b ...*Blob) *BlobCreate {
	ids := make([]uuid.UUID, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return bc.AddLinkIDs(ids...)
}

// Mutation returns the BlobMutation object of the builder.
func (bc *BlobCreate) Mutation() *BlobMutation {
	return bc.mutation
}

// Save creates the Blob in the database.
func (bc *BlobCreate) Save(ctx context.Context) (*Blob, error) {
	var (
		err  error
		node *Blob
	)
	bc.defaults()
	if len(bc.hooks) == 0 {
		if err = bc.check(); err != nil {
			return nil, err
		}
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BlobMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bc.check(); err != nil {
				return nil, err
			}
			bc.mutation = mutation
			if node, err = bc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			if bc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, bc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Blob)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from BlobMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BlobCreate) SaveX(ctx context.Context) *Blob {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BlobCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BlobCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BlobCreate) defaults() {
	if _, ok := bc.mutation.UUID(); !ok {
		v := blob.DefaultUUID()
		bc.mutation.SetUUID(v)
	}
	if _, ok := bc.mutation.Count(); !ok {
		v := blob.DefaultCount
		bc.mutation.SetCount(v)
	}
	if _, ok := bc.mutation.ID(); !ok {
		v := blob.DefaultID()
		bc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BlobCreate) check() error {
	if _, ok := bc.mutation.UUID(); !ok {
		return &ValidationError{Name: "uuid", err: errors.New(`ent: missing required field "Blob.uuid"`)}
	}
	if _, ok := bc.mutation.Count(); !ok {
		return &ValidationError{Name: "count", err: errors.New(`ent: missing required field "Blob.count"`)}
	}
	return nil
}

func (bc *BlobCreate) sqlSave(ctx context.Context) (*Blob, error) {
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
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
	return _node, nil
}

func (bc *BlobCreate) createSpec() (*Blob, *sqlgraph.CreateSpec) {
	var (
		_node = &Blob{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: blob.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: blob.FieldID,
			},
		}
	)
	_spec.OnConflict = bc.conflict
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := bc.mutation.UUID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: blob.FieldUUID,
		})
		_node.UUID = value
	}
	if value, ok := bc.mutation.Count(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: blob.FieldCount,
		})
		_node.Count = value
	}
	if nodes := bc.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   blob.ParentTable,
			Columns: []string{blob.ParentColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: blob.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.blob_parent = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.LinksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blob.LinksTable,
			Columns: blob.LinksPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: blob.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &BlobLinkCreate{config: bc.config, mutation: newBlobLinkMutation(bc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Blob.Create().
//		SetUUID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BlobUpsert) {
//			SetUUID(v+v).
//		}).
//		Exec(ctx)
//
func (bc *BlobCreate) OnConflict(opts ...sql.ConflictOption) *BlobUpsertOne {
	bc.conflict = opts
	return &BlobUpsertOne{
		create: bc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Blob.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (bc *BlobCreate) OnConflictColumns(columns ...string) *BlobUpsertOne {
	bc.conflict = append(bc.conflict, sql.ConflictColumns(columns...))
	return &BlobUpsertOne{
		create: bc,
	}
}

type (
	// BlobUpsertOne is the builder for "upsert"-ing
	//  one Blob node.
	BlobUpsertOne struct {
		create *BlobCreate
	}

	// BlobUpsert is the "OnConflict" setter.
	BlobUpsert struct {
		*sql.UpdateSet
	}
)

// SetUUID sets the "uuid" field.
func (u *BlobUpsert) SetUUID(v uuid.UUID) *BlobUpsert {
	u.Set(blob.FieldUUID, v)
	return u
}

// UpdateUUID sets the "uuid" field to the value that was provided on create.
func (u *BlobUpsert) UpdateUUID() *BlobUpsert {
	u.SetExcluded(blob.FieldUUID)
	return u
}

// SetCount sets the "count" field.
func (u *BlobUpsert) SetCount(v int) *BlobUpsert {
	u.Set(blob.FieldCount, v)
	return u
}

// UpdateCount sets the "count" field to the value that was provided on create.
func (u *BlobUpsert) UpdateCount() *BlobUpsert {
	u.SetExcluded(blob.FieldCount)
	return u
}

// AddCount adds v to the "count" field.
func (u *BlobUpsert) AddCount(v int) *BlobUpsert {
	u.Add(blob.FieldCount, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Blob.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(blob.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *BlobUpsertOne) UpdateNewValues() *BlobUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(blob.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Blob.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *BlobUpsertOne) Ignore() *BlobUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BlobUpsertOne) DoNothing() *BlobUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BlobCreate.OnConflict
// documentation for more info.
func (u *BlobUpsertOne) Update(set func(*BlobUpsert)) *BlobUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BlobUpsert{UpdateSet: update})
	}))
	return u
}

// SetUUID sets the "uuid" field.
func (u *BlobUpsertOne) SetUUID(v uuid.UUID) *BlobUpsertOne {
	return u.Update(func(s *BlobUpsert) {
		s.SetUUID(v)
	})
}

// UpdateUUID sets the "uuid" field to the value that was provided on create.
func (u *BlobUpsertOne) UpdateUUID() *BlobUpsertOne {
	return u.Update(func(s *BlobUpsert) {
		s.UpdateUUID()
	})
}

// SetCount sets the "count" field.
func (u *BlobUpsertOne) SetCount(v int) *BlobUpsertOne {
	return u.Update(func(s *BlobUpsert) {
		s.SetCount(v)
	})
}

// AddCount adds v to the "count" field.
func (u *BlobUpsertOne) AddCount(v int) *BlobUpsertOne {
	return u.Update(func(s *BlobUpsert) {
		s.AddCount(v)
	})
}

// UpdateCount sets the "count" field to the value that was provided on create.
func (u *BlobUpsertOne) UpdateCount() *BlobUpsertOne {
	return u.Update(func(s *BlobUpsert) {
		s.UpdateCount()
	})
}

// Exec executes the query.
func (u *BlobUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for BlobCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BlobUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *BlobUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: BlobUpsertOne.ID is not supported by MySQL driver. Use BlobUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *BlobUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// BlobCreateBulk is the builder for creating many Blob entities in bulk.
type BlobCreateBulk struct {
	config
	builders []*BlobCreate
	conflict []sql.ConflictOption
}

// Save creates the Blob entities in the database.
func (bcb *BlobCreateBulk) Save(ctx context.Context) ([]*Blob, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Blob, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BlobMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = bcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BlobCreateBulk) SaveX(ctx context.Context) []*Blob {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BlobCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BlobCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Blob.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BlobUpsert) {
//			SetUUID(v+v).
//		}).
//		Exec(ctx)
//
func (bcb *BlobCreateBulk) OnConflict(opts ...sql.ConflictOption) *BlobUpsertBulk {
	bcb.conflict = opts
	return &BlobUpsertBulk{
		create: bcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Blob.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (bcb *BlobCreateBulk) OnConflictColumns(columns ...string) *BlobUpsertBulk {
	bcb.conflict = append(bcb.conflict, sql.ConflictColumns(columns...))
	return &BlobUpsertBulk{
		create: bcb,
	}
}

// BlobUpsertBulk is the builder for "upsert"-ing
// a bulk of Blob nodes.
type BlobUpsertBulk struct {
	create *BlobCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Blob.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(blob.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *BlobUpsertBulk) UpdateNewValues() *BlobUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(blob.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Blob.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *BlobUpsertBulk) Ignore() *BlobUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BlobUpsertBulk) DoNothing() *BlobUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BlobCreateBulk.OnConflict
// documentation for more info.
func (u *BlobUpsertBulk) Update(set func(*BlobUpsert)) *BlobUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BlobUpsert{UpdateSet: update})
	}))
	return u
}

// SetUUID sets the "uuid" field.
func (u *BlobUpsertBulk) SetUUID(v uuid.UUID) *BlobUpsertBulk {
	return u.Update(func(s *BlobUpsert) {
		s.SetUUID(v)
	})
}

// UpdateUUID sets the "uuid" field to the value that was provided on create.
func (u *BlobUpsertBulk) UpdateUUID() *BlobUpsertBulk {
	return u.Update(func(s *BlobUpsert) {
		s.UpdateUUID()
	})
}

// SetCount sets the "count" field.
func (u *BlobUpsertBulk) SetCount(v int) *BlobUpsertBulk {
	return u.Update(func(s *BlobUpsert) {
		s.SetCount(v)
	})
}

// AddCount adds v to the "count" field.
func (u *BlobUpsertBulk) AddCount(v int) *BlobUpsertBulk {
	return u.Update(func(s *BlobUpsert) {
		s.AddCount(v)
	})
}

// UpdateCount sets the "count" field to the value that was provided on create.
func (u *BlobUpsertBulk) UpdateCount() *BlobUpsertBulk {
	return u.Update(func(s *BlobUpsert) {
		s.UpdateCount()
	})
}

// Exec executes the query.
func (u *BlobUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the BlobCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for BlobCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BlobUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
