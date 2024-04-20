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
	"github.com/yudai2929/kendai-navi/backend/db/ent/classreview"
	"github.com/yudai2929/kendai-navi/backend/db/ent/classreviewlike"
	"github.com/yudai2929/kendai-navi/backend/db/ent/predicate"
	"github.com/yudai2929/kendai-navi/backend/db/ent/user"
)

// ClassReviewUpdate is the builder for updating ClassReview entities.
type ClassReviewUpdate struct {
	config
	hooks    []Hook
	mutation *ClassReviewMutation
}

// Where appends a list predicates to the ClassReviewUpdate builder.
func (cru *ClassReviewUpdate) Where(ps ...predicate.ClassReview) *ClassReviewUpdate {
	cru.mutation.Where(ps...)
	return cru
}

// SetUserID sets the "user_id" field.
func (cru *ClassReviewUpdate) SetUserID(s string) *ClassReviewUpdate {
	cru.mutation.SetUserID(s)
	return cru
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cru *ClassReviewUpdate) SetNillableUserID(s *string) *ClassReviewUpdate {
	if s != nil {
		cru.SetUserID(*s)
	}
	return cru
}

// SetClassID sets the "class_id" field.
func (cru *ClassReviewUpdate) SetClassID(i int) *ClassReviewUpdate {
	cru.mutation.ResetClassID()
	cru.mutation.SetClassID(i)
	return cru
}

// SetNillableClassID sets the "class_id" field if the given value is not nil.
func (cru *ClassReviewUpdate) SetNillableClassID(i *int) *ClassReviewUpdate {
	if i != nil {
		cru.SetClassID(*i)
	}
	return cru
}

// AddClassID adds i to the "class_id" field.
func (cru *ClassReviewUpdate) AddClassID(i int) *ClassReviewUpdate {
	cru.mutation.AddClassID(i)
	return cru
}

// SetTeacherID sets the "teacher_id" field.
func (cru *ClassReviewUpdate) SetTeacherID(i int) *ClassReviewUpdate {
	cru.mutation.ResetTeacherID()
	cru.mutation.SetTeacherID(i)
	return cru
}

// SetNillableTeacherID sets the "teacher_id" field if the given value is not nil.
func (cru *ClassReviewUpdate) SetNillableTeacherID(i *int) *ClassReviewUpdate {
	if i != nil {
		cru.SetTeacherID(*i)
	}
	return cru
}

// AddTeacherID adds i to the "teacher_id" field.
func (cru *ClassReviewUpdate) AddTeacherID(i int) *ClassReviewUpdate {
	cru.mutation.AddTeacherID(i)
	return cru
}

// SetComment sets the "comment" field.
func (cru *ClassReviewUpdate) SetComment(s string) *ClassReviewUpdate {
	cru.mutation.SetComment(s)
	return cru
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (cru *ClassReviewUpdate) SetNillableComment(s *string) *ClassReviewUpdate {
	if s != nil {
		cru.SetComment(*s)
	}
	return cru
}

// SetClassYear sets the "class_year" field.
func (cru *ClassReviewUpdate) SetClassYear(i int) *ClassReviewUpdate {
	cru.mutation.ResetClassYear()
	cru.mutation.SetClassYear(i)
	return cru
}

// SetNillableClassYear sets the "class_year" field if the given value is not nil.
func (cru *ClassReviewUpdate) SetNillableClassYear(i *int) *ClassReviewUpdate {
	if i != nil {
		cru.SetClassYear(*i)
	}
	return cru
}

// AddClassYear adds i to the "class_year" field.
func (cru *ClassReviewUpdate) AddClassYear(i int) *ClassReviewUpdate {
	cru.mutation.AddClassYear(i)
	return cru
}

// SetTerm sets the "term" field.
func (cru *ClassReviewUpdate) SetTerm(i int) *ClassReviewUpdate {
	cru.mutation.ResetTerm()
	cru.mutation.SetTerm(i)
	return cru
}

// SetNillableTerm sets the "term" field if the given value is not nil.
func (cru *ClassReviewUpdate) SetNillableTerm(i *int) *ClassReviewUpdate {
	if i != nil {
		cru.SetTerm(*i)
	}
	return cru
}

// AddTerm adds i to the "term" field.
func (cru *ClassReviewUpdate) AddTerm(i int) *ClassReviewUpdate {
	cru.mutation.AddTerm(i)
	return cru
}

// SetSatisfactionLevel sets the "satisfaction_level" field.
func (cru *ClassReviewUpdate) SetSatisfactionLevel(i int) *ClassReviewUpdate {
	cru.mutation.ResetSatisfactionLevel()
	cru.mutation.SetSatisfactionLevel(i)
	return cru
}

// SetNillableSatisfactionLevel sets the "satisfaction_level" field if the given value is not nil.
func (cru *ClassReviewUpdate) SetNillableSatisfactionLevel(i *int) *ClassReviewUpdate {
	if i != nil {
		cru.SetSatisfactionLevel(*i)
	}
	return cru
}

// AddSatisfactionLevel adds i to the "satisfaction_level" field.
func (cru *ClassReviewUpdate) AddSatisfactionLevel(i int) *ClassReviewUpdate {
	cru.mutation.AddSatisfactionLevel(i)
	return cru
}

// SetEasyLevel sets the "easy_level" field.
func (cru *ClassReviewUpdate) SetEasyLevel(i int) *ClassReviewUpdate {
	cru.mutation.ResetEasyLevel()
	cru.mutation.SetEasyLevel(i)
	return cru
}

// SetNillableEasyLevel sets the "easy_level" field if the given value is not nil.
func (cru *ClassReviewUpdate) SetNillableEasyLevel(i *int) *ClassReviewUpdate {
	if i != nil {
		cru.SetEasyLevel(*i)
	}
	return cru
}

// AddEasyLevel adds i to the "easy_level" field.
func (cru *ClassReviewUpdate) AddEasyLevel(i int) *ClassReviewUpdate {
	cru.mutation.AddEasyLevel(i)
	return cru
}

// SetAttendanceMethod sets the "attendance_method" field.
func (cru *ClassReviewUpdate) SetAttendanceMethod(i int) *ClassReviewUpdate {
	cru.mutation.ResetAttendanceMethod()
	cru.mutation.SetAttendanceMethod(i)
	return cru
}

// SetNillableAttendanceMethod sets the "attendance_method" field if the given value is not nil.
func (cru *ClassReviewUpdate) SetNillableAttendanceMethod(i *int) *ClassReviewUpdate {
	if i != nil {
		cru.SetAttendanceMethod(*i)
	}
	return cru
}

// AddAttendanceMethod adds i to the "attendance_method" field.
func (cru *ClassReviewUpdate) AddAttendanceMethod(i int) *ClassReviewUpdate {
	cru.mutation.AddAttendanceMethod(i)
	return cru
}

// SetEvaluationMethod sets the "evaluation_method" field.
func (cru *ClassReviewUpdate) SetEvaluationMethod(i int) *ClassReviewUpdate {
	cru.mutation.ResetEvaluationMethod()
	cru.mutation.SetEvaluationMethod(i)
	return cru
}

// SetNillableEvaluationMethod sets the "evaluation_method" field if the given value is not nil.
func (cru *ClassReviewUpdate) SetNillableEvaluationMethod(i *int) *ClassReviewUpdate {
	if i != nil {
		cru.SetEvaluationMethod(*i)
	}
	return cru
}

// AddEvaluationMethod adds i to the "evaluation_method" field.
func (cru *ClassReviewUpdate) AddEvaluationMethod(i int) *ClassReviewUpdate {
	cru.mutation.AddEvaluationMethod(i)
	return cru
}

// SetCreatedAt sets the "created_at" field.
func (cru *ClassReviewUpdate) SetCreatedAt(t time.Time) *ClassReviewUpdate {
	cru.mutation.SetCreatedAt(t)
	return cru
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cru *ClassReviewUpdate) SetNillableCreatedAt(t *time.Time) *ClassReviewUpdate {
	if t != nil {
		cru.SetCreatedAt(*t)
	}
	return cru
}

// SetUpdatedAt sets the "updated_at" field.
func (cru *ClassReviewUpdate) SetUpdatedAt(t time.Time) *ClassReviewUpdate {
	cru.mutation.SetUpdatedAt(t)
	return cru
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cru *ClassReviewUpdate) SetNillableUpdatedAt(t *time.Time) *ClassReviewUpdate {
	if t != nil {
		cru.SetUpdatedAt(*t)
	}
	return cru
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (cru *ClassReviewUpdate) SetUsersID(id string) *ClassReviewUpdate {
	cru.mutation.SetUsersID(id)
	return cru
}

// SetUsers sets the "users" edge to the User entity.
func (cru *ClassReviewUpdate) SetUsers(u *User) *ClassReviewUpdate {
	return cru.SetUsersID(u.ID)
}

// AddClassReviewLikeIDs adds the "class_review_likes" edge to the ClassReviewLike entity by IDs.
func (cru *ClassReviewUpdate) AddClassReviewLikeIDs(ids ...int) *ClassReviewUpdate {
	cru.mutation.AddClassReviewLikeIDs(ids...)
	return cru
}

// AddClassReviewLikes adds the "class_review_likes" edges to the ClassReviewLike entity.
func (cru *ClassReviewUpdate) AddClassReviewLikes(c ...*ClassReviewLike) *ClassReviewUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cru.AddClassReviewLikeIDs(ids...)
}

// Mutation returns the ClassReviewMutation object of the builder.
func (cru *ClassReviewUpdate) Mutation() *ClassReviewMutation {
	return cru.mutation
}

// ClearUsers clears the "users" edge to the User entity.
func (cru *ClassReviewUpdate) ClearUsers() *ClassReviewUpdate {
	cru.mutation.ClearUsers()
	return cru
}

// ClearClassReviewLikes clears all "class_review_likes" edges to the ClassReviewLike entity.
func (cru *ClassReviewUpdate) ClearClassReviewLikes() *ClassReviewUpdate {
	cru.mutation.ClearClassReviewLikes()
	return cru
}

// RemoveClassReviewLikeIDs removes the "class_review_likes" edge to ClassReviewLike entities by IDs.
func (cru *ClassReviewUpdate) RemoveClassReviewLikeIDs(ids ...int) *ClassReviewUpdate {
	cru.mutation.RemoveClassReviewLikeIDs(ids...)
	return cru
}

// RemoveClassReviewLikes removes "class_review_likes" edges to ClassReviewLike entities.
func (cru *ClassReviewUpdate) RemoveClassReviewLikes(c ...*ClassReviewLike) *ClassReviewUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cru.RemoveClassReviewLikeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cru *ClassReviewUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cru.sqlSave, cru.mutation, cru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cru *ClassReviewUpdate) SaveX(ctx context.Context) int {
	affected, err := cru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cru *ClassReviewUpdate) Exec(ctx context.Context) error {
	_, err := cru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cru *ClassReviewUpdate) ExecX(ctx context.Context) {
	if err := cru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cru *ClassReviewUpdate) check() error {
	if _, ok := cru.mutation.UsersID(); cru.mutation.UsersCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ClassReview.users"`)
	}
	return nil
}

func (cru *ClassReviewUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(classreview.Table, classreview.Columns, sqlgraph.NewFieldSpec(classreview.FieldID, field.TypeString))
	if ps := cru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cru.mutation.ClassID(); ok {
		_spec.SetField(classreview.FieldClassID, field.TypeInt, value)
	}
	if value, ok := cru.mutation.AddedClassID(); ok {
		_spec.AddField(classreview.FieldClassID, field.TypeInt, value)
	}
	if value, ok := cru.mutation.TeacherID(); ok {
		_spec.SetField(classreview.FieldTeacherID, field.TypeInt, value)
	}
	if value, ok := cru.mutation.AddedTeacherID(); ok {
		_spec.AddField(classreview.FieldTeacherID, field.TypeInt, value)
	}
	if value, ok := cru.mutation.Comment(); ok {
		_spec.SetField(classreview.FieldComment, field.TypeString, value)
	}
	if value, ok := cru.mutation.ClassYear(); ok {
		_spec.SetField(classreview.FieldClassYear, field.TypeInt, value)
	}
	if value, ok := cru.mutation.AddedClassYear(); ok {
		_spec.AddField(classreview.FieldClassYear, field.TypeInt, value)
	}
	if value, ok := cru.mutation.Term(); ok {
		_spec.SetField(classreview.FieldTerm, field.TypeInt, value)
	}
	if value, ok := cru.mutation.AddedTerm(); ok {
		_spec.AddField(classreview.FieldTerm, field.TypeInt, value)
	}
	if value, ok := cru.mutation.SatisfactionLevel(); ok {
		_spec.SetField(classreview.FieldSatisfactionLevel, field.TypeInt, value)
	}
	if value, ok := cru.mutation.AddedSatisfactionLevel(); ok {
		_spec.AddField(classreview.FieldSatisfactionLevel, field.TypeInt, value)
	}
	if value, ok := cru.mutation.EasyLevel(); ok {
		_spec.SetField(classreview.FieldEasyLevel, field.TypeInt, value)
	}
	if value, ok := cru.mutation.AddedEasyLevel(); ok {
		_spec.AddField(classreview.FieldEasyLevel, field.TypeInt, value)
	}
	if value, ok := cru.mutation.AttendanceMethod(); ok {
		_spec.SetField(classreview.FieldAttendanceMethod, field.TypeInt, value)
	}
	if value, ok := cru.mutation.AddedAttendanceMethod(); ok {
		_spec.AddField(classreview.FieldAttendanceMethod, field.TypeInt, value)
	}
	if value, ok := cru.mutation.EvaluationMethod(); ok {
		_spec.SetField(classreview.FieldEvaluationMethod, field.TypeInt, value)
	}
	if value, ok := cru.mutation.AddedEvaluationMethod(); ok {
		_spec.AddField(classreview.FieldEvaluationMethod, field.TypeInt, value)
	}
	if value, ok := cru.mutation.CreatedAt(); ok {
		_spec.SetField(classreview.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cru.mutation.UpdatedAt(); ok {
		_spec.SetField(classreview.FieldUpdatedAt, field.TypeTime, value)
	}
	if cru.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   classreview.UsersTable,
			Columns: []string{classreview.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cru.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   classreview.UsersTable,
			Columns: []string{classreview.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cru.mutation.ClassReviewLikesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   classreview.ClassReviewLikesTable,
			Columns: []string{classreview.ClassReviewLikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(classreviewlike.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cru.mutation.RemovedClassReviewLikesIDs(); len(nodes) > 0 && !cru.mutation.ClassReviewLikesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   classreview.ClassReviewLikesTable,
			Columns: []string{classreview.ClassReviewLikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(classreviewlike.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cru.mutation.ClassReviewLikesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   classreview.ClassReviewLikesTable,
			Columns: []string{classreview.ClassReviewLikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(classreviewlike.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{classreview.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cru.mutation.done = true
	return n, nil
}

// ClassReviewUpdateOne is the builder for updating a single ClassReview entity.
type ClassReviewUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ClassReviewMutation
}

// SetUserID sets the "user_id" field.
func (cruo *ClassReviewUpdateOne) SetUserID(s string) *ClassReviewUpdateOne {
	cruo.mutation.SetUserID(s)
	return cruo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cruo *ClassReviewUpdateOne) SetNillableUserID(s *string) *ClassReviewUpdateOne {
	if s != nil {
		cruo.SetUserID(*s)
	}
	return cruo
}

// SetClassID sets the "class_id" field.
func (cruo *ClassReviewUpdateOne) SetClassID(i int) *ClassReviewUpdateOne {
	cruo.mutation.ResetClassID()
	cruo.mutation.SetClassID(i)
	return cruo
}

// SetNillableClassID sets the "class_id" field if the given value is not nil.
func (cruo *ClassReviewUpdateOne) SetNillableClassID(i *int) *ClassReviewUpdateOne {
	if i != nil {
		cruo.SetClassID(*i)
	}
	return cruo
}

// AddClassID adds i to the "class_id" field.
func (cruo *ClassReviewUpdateOne) AddClassID(i int) *ClassReviewUpdateOne {
	cruo.mutation.AddClassID(i)
	return cruo
}

// SetTeacherID sets the "teacher_id" field.
func (cruo *ClassReviewUpdateOne) SetTeacherID(i int) *ClassReviewUpdateOne {
	cruo.mutation.ResetTeacherID()
	cruo.mutation.SetTeacherID(i)
	return cruo
}

// SetNillableTeacherID sets the "teacher_id" field if the given value is not nil.
func (cruo *ClassReviewUpdateOne) SetNillableTeacherID(i *int) *ClassReviewUpdateOne {
	if i != nil {
		cruo.SetTeacherID(*i)
	}
	return cruo
}

// AddTeacherID adds i to the "teacher_id" field.
func (cruo *ClassReviewUpdateOne) AddTeacherID(i int) *ClassReviewUpdateOne {
	cruo.mutation.AddTeacherID(i)
	return cruo
}

// SetComment sets the "comment" field.
func (cruo *ClassReviewUpdateOne) SetComment(s string) *ClassReviewUpdateOne {
	cruo.mutation.SetComment(s)
	return cruo
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (cruo *ClassReviewUpdateOne) SetNillableComment(s *string) *ClassReviewUpdateOne {
	if s != nil {
		cruo.SetComment(*s)
	}
	return cruo
}

// SetClassYear sets the "class_year" field.
func (cruo *ClassReviewUpdateOne) SetClassYear(i int) *ClassReviewUpdateOne {
	cruo.mutation.ResetClassYear()
	cruo.mutation.SetClassYear(i)
	return cruo
}

// SetNillableClassYear sets the "class_year" field if the given value is not nil.
func (cruo *ClassReviewUpdateOne) SetNillableClassYear(i *int) *ClassReviewUpdateOne {
	if i != nil {
		cruo.SetClassYear(*i)
	}
	return cruo
}

// AddClassYear adds i to the "class_year" field.
func (cruo *ClassReviewUpdateOne) AddClassYear(i int) *ClassReviewUpdateOne {
	cruo.mutation.AddClassYear(i)
	return cruo
}

// SetTerm sets the "term" field.
func (cruo *ClassReviewUpdateOne) SetTerm(i int) *ClassReviewUpdateOne {
	cruo.mutation.ResetTerm()
	cruo.mutation.SetTerm(i)
	return cruo
}

// SetNillableTerm sets the "term" field if the given value is not nil.
func (cruo *ClassReviewUpdateOne) SetNillableTerm(i *int) *ClassReviewUpdateOne {
	if i != nil {
		cruo.SetTerm(*i)
	}
	return cruo
}

// AddTerm adds i to the "term" field.
func (cruo *ClassReviewUpdateOne) AddTerm(i int) *ClassReviewUpdateOne {
	cruo.mutation.AddTerm(i)
	return cruo
}

// SetSatisfactionLevel sets the "satisfaction_level" field.
func (cruo *ClassReviewUpdateOne) SetSatisfactionLevel(i int) *ClassReviewUpdateOne {
	cruo.mutation.ResetSatisfactionLevel()
	cruo.mutation.SetSatisfactionLevel(i)
	return cruo
}

// SetNillableSatisfactionLevel sets the "satisfaction_level" field if the given value is not nil.
func (cruo *ClassReviewUpdateOne) SetNillableSatisfactionLevel(i *int) *ClassReviewUpdateOne {
	if i != nil {
		cruo.SetSatisfactionLevel(*i)
	}
	return cruo
}

// AddSatisfactionLevel adds i to the "satisfaction_level" field.
func (cruo *ClassReviewUpdateOne) AddSatisfactionLevel(i int) *ClassReviewUpdateOne {
	cruo.mutation.AddSatisfactionLevel(i)
	return cruo
}

// SetEasyLevel sets the "easy_level" field.
func (cruo *ClassReviewUpdateOne) SetEasyLevel(i int) *ClassReviewUpdateOne {
	cruo.mutation.ResetEasyLevel()
	cruo.mutation.SetEasyLevel(i)
	return cruo
}

// SetNillableEasyLevel sets the "easy_level" field if the given value is not nil.
func (cruo *ClassReviewUpdateOne) SetNillableEasyLevel(i *int) *ClassReviewUpdateOne {
	if i != nil {
		cruo.SetEasyLevel(*i)
	}
	return cruo
}

// AddEasyLevel adds i to the "easy_level" field.
func (cruo *ClassReviewUpdateOne) AddEasyLevel(i int) *ClassReviewUpdateOne {
	cruo.mutation.AddEasyLevel(i)
	return cruo
}

// SetAttendanceMethod sets the "attendance_method" field.
func (cruo *ClassReviewUpdateOne) SetAttendanceMethod(i int) *ClassReviewUpdateOne {
	cruo.mutation.ResetAttendanceMethod()
	cruo.mutation.SetAttendanceMethod(i)
	return cruo
}

// SetNillableAttendanceMethod sets the "attendance_method" field if the given value is not nil.
func (cruo *ClassReviewUpdateOne) SetNillableAttendanceMethod(i *int) *ClassReviewUpdateOne {
	if i != nil {
		cruo.SetAttendanceMethod(*i)
	}
	return cruo
}

// AddAttendanceMethod adds i to the "attendance_method" field.
func (cruo *ClassReviewUpdateOne) AddAttendanceMethod(i int) *ClassReviewUpdateOne {
	cruo.mutation.AddAttendanceMethod(i)
	return cruo
}

// SetEvaluationMethod sets the "evaluation_method" field.
func (cruo *ClassReviewUpdateOne) SetEvaluationMethod(i int) *ClassReviewUpdateOne {
	cruo.mutation.ResetEvaluationMethod()
	cruo.mutation.SetEvaluationMethod(i)
	return cruo
}

// SetNillableEvaluationMethod sets the "evaluation_method" field if the given value is not nil.
func (cruo *ClassReviewUpdateOne) SetNillableEvaluationMethod(i *int) *ClassReviewUpdateOne {
	if i != nil {
		cruo.SetEvaluationMethod(*i)
	}
	return cruo
}

// AddEvaluationMethod adds i to the "evaluation_method" field.
func (cruo *ClassReviewUpdateOne) AddEvaluationMethod(i int) *ClassReviewUpdateOne {
	cruo.mutation.AddEvaluationMethod(i)
	return cruo
}

// SetCreatedAt sets the "created_at" field.
func (cruo *ClassReviewUpdateOne) SetCreatedAt(t time.Time) *ClassReviewUpdateOne {
	cruo.mutation.SetCreatedAt(t)
	return cruo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cruo *ClassReviewUpdateOne) SetNillableCreatedAt(t *time.Time) *ClassReviewUpdateOne {
	if t != nil {
		cruo.SetCreatedAt(*t)
	}
	return cruo
}

// SetUpdatedAt sets the "updated_at" field.
func (cruo *ClassReviewUpdateOne) SetUpdatedAt(t time.Time) *ClassReviewUpdateOne {
	cruo.mutation.SetUpdatedAt(t)
	return cruo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cruo *ClassReviewUpdateOne) SetNillableUpdatedAt(t *time.Time) *ClassReviewUpdateOne {
	if t != nil {
		cruo.SetUpdatedAt(*t)
	}
	return cruo
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (cruo *ClassReviewUpdateOne) SetUsersID(id string) *ClassReviewUpdateOne {
	cruo.mutation.SetUsersID(id)
	return cruo
}

// SetUsers sets the "users" edge to the User entity.
func (cruo *ClassReviewUpdateOne) SetUsers(u *User) *ClassReviewUpdateOne {
	return cruo.SetUsersID(u.ID)
}

// AddClassReviewLikeIDs adds the "class_review_likes" edge to the ClassReviewLike entity by IDs.
func (cruo *ClassReviewUpdateOne) AddClassReviewLikeIDs(ids ...int) *ClassReviewUpdateOne {
	cruo.mutation.AddClassReviewLikeIDs(ids...)
	return cruo
}

// AddClassReviewLikes adds the "class_review_likes" edges to the ClassReviewLike entity.
func (cruo *ClassReviewUpdateOne) AddClassReviewLikes(c ...*ClassReviewLike) *ClassReviewUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cruo.AddClassReviewLikeIDs(ids...)
}

// Mutation returns the ClassReviewMutation object of the builder.
func (cruo *ClassReviewUpdateOne) Mutation() *ClassReviewMutation {
	return cruo.mutation
}

// ClearUsers clears the "users" edge to the User entity.
func (cruo *ClassReviewUpdateOne) ClearUsers() *ClassReviewUpdateOne {
	cruo.mutation.ClearUsers()
	return cruo
}

// ClearClassReviewLikes clears all "class_review_likes" edges to the ClassReviewLike entity.
func (cruo *ClassReviewUpdateOne) ClearClassReviewLikes() *ClassReviewUpdateOne {
	cruo.mutation.ClearClassReviewLikes()
	return cruo
}

// RemoveClassReviewLikeIDs removes the "class_review_likes" edge to ClassReviewLike entities by IDs.
func (cruo *ClassReviewUpdateOne) RemoveClassReviewLikeIDs(ids ...int) *ClassReviewUpdateOne {
	cruo.mutation.RemoveClassReviewLikeIDs(ids...)
	return cruo
}

// RemoveClassReviewLikes removes "class_review_likes" edges to ClassReviewLike entities.
func (cruo *ClassReviewUpdateOne) RemoveClassReviewLikes(c ...*ClassReviewLike) *ClassReviewUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cruo.RemoveClassReviewLikeIDs(ids...)
}

// Where appends a list predicates to the ClassReviewUpdate builder.
func (cruo *ClassReviewUpdateOne) Where(ps ...predicate.ClassReview) *ClassReviewUpdateOne {
	cruo.mutation.Where(ps...)
	return cruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cruo *ClassReviewUpdateOne) Select(field string, fields ...string) *ClassReviewUpdateOne {
	cruo.fields = append([]string{field}, fields...)
	return cruo
}

// Save executes the query and returns the updated ClassReview entity.
func (cruo *ClassReviewUpdateOne) Save(ctx context.Context) (*ClassReview, error) {
	return withHooks(ctx, cruo.sqlSave, cruo.mutation, cruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cruo *ClassReviewUpdateOne) SaveX(ctx context.Context) *ClassReview {
	node, err := cruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cruo *ClassReviewUpdateOne) Exec(ctx context.Context) error {
	_, err := cruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cruo *ClassReviewUpdateOne) ExecX(ctx context.Context) {
	if err := cruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cruo *ClassReviewUpdateOne) check() error {
	if _, ok := cruo.mutation.UsersID(); cruo.mutation.UsersCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ClassReview.users"`)
	}
	return nil
}

func (cruo *ClassReviewUpdateOne) sqlSave(ctx context.Context) (_node *ClassReview, err error) {
	if err := cruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(classreview.Table, classreview.Columns, sqlgraph.NewFieldSpec(classreview.FieldID, field.TypeString))
	id, ok := cruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ClassReview.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, classreview.FieldID)
		for _, f := range fields {
			if !classreview.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != classreview.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cruo.mutation.ClassID(); ok {
		_spec.SetField(classreview.FieldClassID, field.TypeInt, value)
	}
	if value, ok := cruo.mutation.AddedClassID(); ok {
		_spec.AddField(classreview.FieldClassID, field.TypeInt, value)
	}
	if value, ok := cruo.mutation.TeacherID(); ok {
		_spec.SetField(classreview.FieldTeacherID, field.TypeInt, value)
	}
	if value, ok := cruo.mutation.AddedTeacherID(); ok {
		_spec.AddField(classreview.FieldTeacherID, field.TypeInt, value)
	}
	if value, ok := cruo.mutation.Comment(); ok {
		_spec.SetField(classreview.FieldComment, field.TypeString, value)
	}
	if value, ok := cruo.mutation.ClassYear(); ok {
		_spec.SetField(classreview.FieldClassYear, field.TypeInt, value)
	}
	if value, ok := cruo.mutation.AddedClassYear(); ok {
		_spec.AddField(classreview.FieldClassYear, field.TypeInt, value)
	}
	if value, ok := cruo.mutation.Term(); ok {
		_spec.SetField(classreview.FieldTerm, field.TypeInt, value)
	}
	if value, ok := cruo.mutation.AddedTerm(); ok {
		_spec.AddField(classreview.FieldTerm, field.TypeInt, value)
	}
	if value, ok := cruo.mutation.SatisfactionLevel(); ok {
		_spec.SetField(classreview.FieldSatisfactionLevel, field.TypeInt, value)
	}
	if value, ok := cruo.mutation.AddedSatisfactionLevel(); ok {
		_spec.AddField(classreview.FieldSatisfactionLevel, field.TypeInt, value)
	}
	if value, ok := cruo.mutation.EasyLevel(); ok {
		_spec.SetField(classreview.FieldEasyLevel, field.TypeInt, value)
	}
	if value, ok := cruo.mutation.AddedEasyLevel(); ok {
		_spec.AddField(classreview.FieldEasyLevel, field.TypeInt, value)
	}
	if value, ok := cruo.mutation.AttendanceMethod(); ok {
		_spec.SetField(classreview.FieldAttendanceMethod, field.TypeInt, value)
	}
	if value, ok := cruo.mutation.AddedAttendanceMethod(); ok {
		_spec.AddField(classreview.FieldAttendanceMethod, field.TypeInt, value)
	}
	if value, ok := cruo.mutation.EvaluationMethod(); ok {
		_spec.SetField(classreview.FieldEvaluationMethod, field.TypeInt, value)
	}
	if value, ok := cruo.mutation.AddedEvaluationMethod(); ok {
		_spec.AddField(classreview.FieldEvaluationMethod, field.TypeInt, value)
	}
	if value, ok := cruo.mutation.CreatedAt(); ok {
		_spec.SetField(classreview.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := cruo.mutation.UpdatedAt(); ok {
		_spec.SetField(classreview.FieldUpdatedAt, field.TypeTime, value)
	}
	if cruo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   classreview.UsersTable,
			Columns: []string{classreview.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cruo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   classreview.UsersTable,
			Columns: []string{classreview.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cruo.mutation.ClassReviewLikesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   classreview.ClassReviewLikesTable,
			Columns: []string{classreview.ClassReviewLikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(classreviewlike.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cruo.mutation.RemovedClassReviewLikesIDs(); len(nodes) > 0 && !cruo.mutation.ClassReviewLikesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   classreview.ClassReviewLikesTable,
			Columns: []string{classreview.ClassReviewLikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(classreviewlike.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cruo.mutation.ClassReviewLikesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   classreview.ClassReviewLikesTable,
			Columns: []string{classreview.ClassReviewLikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(classreviewlike.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ClassReview{config: cruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{classreview.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cruo.mutation.done = true
	return _node, nil
}
