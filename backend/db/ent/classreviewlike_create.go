// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yudai2929/kendai-navi/backend/db/ent/classreview"
	"github.com/yudai2929/kendai-navi/backend/db/ent/classreviewlike"
	"github.com/yudai2929/kendai-navi/backend/db/ent/user"
)

// ClassReviewLikeCreate is the builder for creating a ClassReviewLike entity.
type ClassReviewLikeCreate struct {
	config
	mutation *ClassReviewLikeMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (crlc *ClassReviewLikeCreate) SetUserID(s string) *ClassReviewLikeCreate {
	crlc.mutation.SetUserID(s)
	return crlc
}

// SetClassReviewID sets the "class_review_id" field.
func (crlc *ClassReviewLikeCreate) SetClassReviewID(s string) *ClassReviewLikeCreate {
	crlc.mutation.SetClassReviewID(s)
	return crlc
}

// SetCreatedAt sets the "created_at" field.
func (crlc *ClassReviewLikeCreate) SetCreatedAt(t time.Time) *ClassReviewLikeCreate {
	crlc.mutation.SetCreatedAt(t)
	return crlc
}

// SetUpdatedAt sets the "updated_at" field.
func (crlc *ClassReviewLikeCreate) SetUpdatedAt(t time.Time) *ClassReviewLikeCreate {
	crlc.mutation.SetUpdatedAt(t)
	return crlc
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (crlc *ClassReviewLikeCreate) SetUsersID(id string) *ClassReviewLikeCreate {
	crlc.mutation.SetUsersID(id)
	return crlc
}

// SetUsers sets the "users" edge to the User entity.
func (crlc *ClassReviewLikeCreate) SetUsers(u *User) *ClassReviewLikeCreate {
	return crlc.SetUsersID(u.ID)
}

// SetClassReviewsID sets the "class_reviews" edge to the ClassReview entity by ID.
func (crlc *ClassReviewLikeCreate) SetClassReviewsID(id string) *ClassReviewLikeCreate {
	crlc.mutation.SetClassReviewsID(id)
	return crlc
}

// SetClassReviews sets the "class_reviews" edge to the ClassReview entity.
func (crlc *ClassReviewLikeCreate) SetClassReviews(c *ClassReview) *ClassReviewLikeCreate {
	return crlc.SetClassReviewsID(c.ID)
}

// Mutation returns the ClassReviewLikeMutation object of the builder.
func (crlc *ClassReviewLikeCreate) Mutation() *ClassReviewLikeMutation {
	return crlc.mutation
}

// Save creates the ClassReviewLike in the database.
func (crlc *ClassReviewLikeCreate) Save(ctx context.Context) (*ClassReviewLike, error) {
	return withHooks(ctx, crlc.sqlSave, crlc.mutation, crlc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (crlc *ClassReviewLikeCreate) SaveX(ctx context.Context) *ClassReviewLike {
	v, err := crlc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (crlc *ClassReviewLikeCreate) Exec(ctx context.Context) error {
	_, err := crlc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (crlc *ClassReviewLikeCreate) ExecX(ctx context.Context) {
	if err := crlc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (crlc *ClassReviewLikeCreate) check() error {
	if _, ok := crlc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "ClassReviewLike.user_id"`)}
	}
	if _, ok := crlc.mutation.ClassReviewID(); !ok {
		return &ValidationError{Name: "class_review_id", err: errors.New(`ent: missing required field "ClassReviewLike.class_review_id"`)}
	}
	if _, ok := crlc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "ClassReviewLike.created_at"`)}
	}
	if _, ok := crlc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ClassReviewLike.updated_at"`)}
	}
	if _, ok := crlc.mutation.UsersID(); !ok {
		return &ValidationError{Name: "users", err: errors.New(`ent: missing required edge "ClassReviewLike.users"`)}
	}
	if _, ok := crlc.mutation.ClassReviewsID(); !ok {
		return &ValidationError{Name: "class_reviews", err: errors.New(`ent: missing required edge "ClassReviewLike.class_reviews"`)}
	}
	return nil
}

func (crlc *ClassReviewLikeCreate) sqlSave(ctx context.Context) (*ClassReviewLike, error) {
	if err := crlc.check(); err != nil {
		return nil, err
	}
	_node, _spec := crlc.createSpec()
	if err := sqlgraph.CreateNode(ctx, crlc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	crlc.mutation.id = &_node.ID
	crlc.mutation.done = true
	return _node, nil
}

func (crlc *ClassReviewLikeCreate) createSpec() (*ClassReviewLike, *sqlgraph.CreateSpec) {
	var (
		_node = &ClassReviewLike{config: crlc.config}
		_spec = sqlgraph.NewCreateSpec(classreviewlike.Table, sqlgraph.NewFieldSpec(classreviewlike.FieldID, field.TypeInt))
	)
	if value, ok := crlc.mutation.CreatedAt(); ok {
		_spec.SetField(classreviewlike.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := crlc.mutation.UpdatedAt(); ok {
		_spec.SetField(classreviewlike.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := crlc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   classreviewlike.UsersTable,
			Columns: []string{classreviewlike.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := crlc.mutation.ClassReviewsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   classreviewlike.ClassReviewsTable,
			Columns: []string{classreviewlike.ClassReviewsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(classreview.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ClassReviewID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ClassReviewLikeCreateBulk is the builder for creating many ClassReviewLike entities in bulk.
type ClassReviewLikeCreateBulk struct {
	config
	err      error
	builders []*ClassReviewLikeCreate
}

// Save creates the ClassReviewLike entities in the database.
func (crlcb *ClassReviewLikeCreateBulk) Save(ctx context.Context) ([]*ClassReviewLike, error) {
	if crlcb.err != nil {
		return nil, crlcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(crlcb.builders))
	nodes := make([]*ClassReviewLike, len(crlcb.builders))
	mutators := make([]Mutator, len(crlcb.builders))
	for i := range crlcb.builders {
		func(i int, root context.Context) {
			builder := crlcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ClassReviewLikeMutation)
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
					_, err = mutators[i+1].Mutate(root, crlcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, crlcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, crlcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (crlcb *ClassReviewLikeCreateBulk) SaveX(ctx context.Context) []*ClassReviewLike {
	v, err := crlcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (crlcb *ClassReviewLikeCreateBulk) Exec(ctx context.Context) error {
	_, err := crlcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (crlcb *ClassReviewLikeCreateBulk) ExecX(ctx context.Context) {
	if err := crlcb.Exec(ctx); err != nil {
		panic(err)
	}
}
