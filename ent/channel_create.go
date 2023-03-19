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
	"github.com/satont/twitch-notifier/ent/follow"
	"github.com/satont/twitch-notifier/ent/stream"
)

// ChannelCreate is the builder for creating a Channel entity.
type ChannelCreate struct {
	config
	mutation *ChannelMutation
	hooks    []Hook
}

// SetChannelID sets the "channel_id" field.
func (cc *ChannelCreate) SetChannelID(s string) *ChannelCreate {
	cc.mutation.SetChannelID(s)
	return cc
}

// SetService sets the "service" field.
func (cc *ChannelCreate) SetService(c channel.Service) *ChannelCreate {
	cc.mutation.SetService(c)
	return cc
}

// SetIsLive sets the "is_live" field.
func (cc *ChannelCreate) SetIsLive(b bool) *ChannelCreate {
	cc.mutation.SetIsLive(b)
	return cc
}

// SetNillableIsLive sets the "is_live" field if the given value is not nil.
func (cc *ChannelCreate) SetNillableIsLive(b *bool) *ChannelCreate {
	if b != nil {
		cc.SetIsLive(*b)
	}
	return cc
}

// SetTitle sets the "title" field.
func (cc *ChannelCreate) SetTitle(s string) *ChannelCreate {
	cc.mutation.SetTitle(s)
	return cc
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (cc *ChannelCreate) SetNillableTitle(s *string) *ChannelCreate {
	if s != nil {
		cc.SetTitle(*s)
	}
	return cc
}

// SetCategory sets the "category" field.
func (cc *ChannelCreate) SetCategory(s string) *ChannelCreate {
	cc.mutation.SetCategory(s)
	return cc
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (cc *ChannelCreate) SetNillableCategory(s *string) *ChannelCreate {
	if s != nil {
		cc.SetCategory(*s)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *ChannelCreate) SetUpdatedAt(t time.Time) *ChannelCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *ChannelCreate) SetNillableUpdatedAt(t *time.Time) *ChannelCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *ChannelCreate) SetID(u uuid.UUID) *ChannelCreate {
	cc.mutation.SetID(u)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *ChannelCreate) SetNillableID(u *uuid.UUID) *ChannelCreate {
	if u != nil {
		cc.SetID(*u)
	}
	return cc
}

// AddFollowIDs adds the "follows" edge to the Follow entity by IDs.
func (cc *ChannelCreate) AddFollowIDs(ids ...uuid.UUID) *ChannelCreate {
	cc.mutation.AddFollowIDs(ids...)
	return cc
}

// AddFollows adds the "follows" edges to the Follow entity.
func (cc *ChannelCreate) AddFollows(f ...*Follow) *ChannelCreate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return cc.AddFollowIDs(ids...)
}

// AddStreamIDs adds the "streams" edge to the Stream entity by IDs.
func (cc *ChannelCreate) AddStreamIDs(ids ...string) *ChannelCreate {
	cc.mutation.AddStreamIDs(ids...)
	return cc
}

// AddStreams adds the "streams" edges to the Stream entity.
func (cc *ChannelCreate) AddStreams(s ...*Stream) *ChannelCreate {
	ids := make([]string, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cc.AddStreamIDs(ids...)
}

// Mutation returns the ChannelMutation object of the builder.
func (cc *ChannelCreate) Mutation() *ChannelMutation {
	return cc.mutation
}

// Save creates the Channel in the database.
func (cc *ChannelCreate) Save(ctx context.Context) (*Channel, error) {
	cc.defaults()
	return withHooks[*Channel, ChannelMutation](ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *ChannelCreate) SaveX(ctx context.Context) *Channel {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ChannelCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ChannelCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ChannelCreate) defaults() {
	if _, ok := cc.mutation.IsLive(); !ok {
		v := channel.DefaultIsLive
		cc.mutation.SetIsLive(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := channel.DefaultID()
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ChannelCreate) check() error {
	if _, ok := cc.mutation.ChannelID(); !ok {
		return &ValidationError{Name: "channel_id", err: errors.New(`ent: missing required field "Channel.channel_id"`)}
	}
	if _, ok := cc.mutation.Service(); !ok {
		return &ValidationError{Name: "service", err: errors.New(`ent: missing required field "Channel.service"`)}
	}
	if v, ok := cc.mutation.Service(); ok {
		if err := channel.ServiceValidator(v); err != nil {
			return &ValidationError{Name: "service", err: fmt.Errorf(`ent: validator failed for field "Channel.service": %w`, err)}
		}
	}
	if _, ok := cc.mutation.IsLive(); !ok {
		return &ValidationError{Name: "is_live", err: errors.New(`ent: missing required field "Channel.is_live"`)}
	}
	return nil
}

func (cc *ChannelCreate) sqlSave(ctx context.Context) (*Channel, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *ChannelCreate) createSpec() (*Channel, *sqlgraph.CreateSpec) {
	var (
		_node = &Channel{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(channel.Table, sqlgraph.NewFieldSpec(channel.FieldID, field.TypeUUID))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := cc.mutation.ChannelID(); ok {
		_spec.SetField(channel.FieldChannelID, field.TypeString, value)
		_node.ChannelID = value
	}
	if value, ok := cc.mutation.Service(); ok {
		_spec.SetField(channel.FieldService, field.TypeEnum, value)
		_node.Service = value
	}
	if value, ok := cc.mutation.IsLive(); ok {
		_spec.SetField(channel.FieldIsLive, field.TypeBool, value)
		_node.IsLive = value
	}
	if value, ok := cc.mutation.Title(); ok {
		_spec.SetField(channel.FieldTitle, field.TypeString, value)
		_node.Title = &value
	}
	if value, ok := cc.mutation.Category(); ok {
		_spec.SetField(channel.FieldCategory, field.TypeString, value)
		_node.Category = &value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(channel.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = &value
	}
	if nodes := cc.mutation.FollowsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.FollowsTable,
			Columns: []string{channel.FollowsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: follow.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.StreamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   channel.StreamsTable,
			Columns: []string{channel.StreamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: stream.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ChannelCreateBulk is the builder for creating many Channel entities in bulk.
type ChannelCreateBulk struct {
	config
	builders []*ChannelCreate
}

// Save creates the Channel entities in the database.
func (ccb *ChannelCreateBulk) Save(ctx context.Context) ([]*Channel, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Channel, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChannelMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ChannelCreateBulk) SaveX(ctx context.Context) []*Channel {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ChannelCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ChannelCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
