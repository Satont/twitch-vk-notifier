// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/satont/twitch-notifier/ent/channel"
	"github.com/satont/twitch-notifier/ent/stream"
)

// StreamCreate is the builder for creating a Stream entity.
type StreamCreate struct {
	config
	mutation *StreamMutation
	hooks    []Hook
}

// SetChannelID sets the "channel_id" field.
func (sc *StreamCreate) SetChannelID(u uuid.UUID) *StreamCreate {
	sc.mutation.SetChannelID(u)
	return sc
}

// SetTitles sets the "titles" field.
func (sc *StreamCreate) SetTitles(s []string) *StreamCreate {
	sc.mutation.SetTitles(s)
	return sc
}

// SetCategories sets the "categories" field.
func (sc *StreamCreate) SetCategories(s []string) *StreamCreate {
	sc.mutation.SetCategories(s)
	return sc
}

// SetStartedAt sets the "started_at" field.
func (sc *StreamCreate) SetStartedAt(t time.Time) *StreamCreate {
	sc.mutation.SetStartedAt(t)
	return sc
}

// SetNillableStartedAt sets the "started_at" field if the given value is not nil.
func (sc *StreamCreate) SetNillableStartedAt(t *time.Time) *StreamCreate {
	if t != nil {
		sc.SetStartedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *StreamCreate) SetUpdatedAt(t time.Time) *StreamCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *StreamCreate) SetNillableUpdatedAt(t *time.Time) *StreamCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetEndedAt sets the "ended_at" field.
func (sc *StreamCreate) SetEndedAt(t time.Time) *StreamCreate {
	sc.mutation.SetEndedAt(t)
	return sc
}

// SetNillableEndedAt sets the "ended_at" field if the given value is not nil.
func (sc *StreamCreate) SetNillableEndedAt(t *time.Time) *StreamCreate {
	if t != nil {
		sc.SetEndedAt(*t)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *StreamCreate) SetID(s string) *StreamCreate {
	sc.mutation.SetID(s)
	return sc
}

// SetChannel sets the "channel" edge to the Channel entity.
func (sc *StreamCreate) SetChannel(c *Channel) *StreamCreate {
	return sc.SetChannelID(c.ID)
}

// Mutation returns the StreamMutation object of the builder.
func (sc *StreamCreate) Mutation() *StreamMutation {
	return sc.mutation
}

// Save creates the Stream in the database.
func (sc *StreamCreate) Save(ctx context.Context) (*Stream, error) {
	sc.defaults()
	return withHooks[*Stream, StreamMutation](ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StreamCreate) SaveX(ctx context.Context) *Stream {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *StreamCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *StreamCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *StreamCreate) defaults() {
	if _, ok := sc.mutation.Titles(); !ok {
		v := stream.DefaultTitles
		sc.mutation.SetTitles(v)
	}
	if _, ok := sc.mutation.Categories(); !ok {
		v := stream.DefaultCategories
		sc.mutation.SetCategories(v)
	}
	if _, ok := sc.mutation.StartedAt(); !ok {
		v := stream.DefaultStartedAt()
		sc.mutation.SetStartedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *StreamCreate) check() error {
	if _, ok := sc.mutation.ChannelID(); !ok {
		return &ValidationError{Name: "channel_id", err: errors.New(`ent: missing required field "Stream.channel_id"`)}
	}
	if _, ok := sc.mutation.ChannelID(); !ok {
		return &ValidationError{Name: "channel", err: errors.New(`ent: missing required edge "Stream.channel"`)}
	}
	return nil
}

func (sc *StreamCreate) sqlSave(ctx context.Context) (*Stream, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Stream.ID type: %T", _spec.ID.Value)
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *StreamCreate) createSpec() (*Stream, *sqlgraph.CreateSpec) {
	var (
		_node = &Stream{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(stream.Table, sqlgraph.NewFieldSpec(stream.FieldID, field.TypeString))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.Titles(); ok {
		_spec.SetField(stream.FieldTitles, field.TypeJSON, value)
		_node.Titles = value
	}
	if value, ok := sc.mutation.Categories(); ok {
		_spec.SetField(stream.FieldCategories, field.TypeJSON, value)
		_node.Categories = value
	}
	if value, ok := sc.mutation.StartedAt(); ok {
		_spec.SetField(stream.FieldStartedAt, field.TypeTime, value)
		_node.StartedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(stream.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = &value
	}
	if value, ok := sc.mutation.EndedAt(); ok {
		_spec.SetField(stream.FieldEndedAt, field.TypeTime, value)
		_node.EndedAt = &value
	}
	if nodes := sc.mutation.ChannelIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   stream.ChannelTable,
			Columns: []string{stream.ChannelColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: channel.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ChannelID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// StreamCreateBulk is the builder for creating many Stream entities in bulk.
type StreamCreateBulk struct {
	config
	builders []*StreamCreate
}

// Save creates the Stream entities in the database.
func (scb *StreamCreateBulk) Save(ctx context.Context) ([]*Stream, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Stream, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*StreamMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *StreamCreateBulk) SaveX(ctx context.Context) []*Stream {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *StreamCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *StreamCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
