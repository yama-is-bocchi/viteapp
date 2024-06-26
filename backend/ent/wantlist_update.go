// Code generated by ent, DO NOT EDIT.

package ent

import (
	"Ebook/ent/predicate"
	"Ebook/ent/user"
	"Ebook/ent/wantlist"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// WantlistUpdate is the builder for updating Wantlist entities.
type WantlistUpdate struct {
	config
	hooks    []Hook
	mutation *WantlistMutation
}

// Where appends a list predicates to the WantlistUpdate builder.
func (wu *WantlistUpdate) Where(ps ...predicate.Wantlist) *WantlistUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetTitle sets the "title" field.
func (wu *WantlistUpdate) SetTitle(s string) *WantlistUpdate {
	wu.mutation.SetTitle(s)
	return wu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (wu *WantlistUpdate) SetNillableTitle(s *string) *WantlistUpdate {
	if s != nil {
		wu.SetTitle(*s)
	}
	return wu
}

// SetUserID sets the "user_id" field.
func (wu *WantlistUpdate) SetUserID(i int) *WantlistUpdate {
	wu.mutation.SetUserID(i)
	return wu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (wu *WantlistUpdate) SetNillableUserID(i *int) *WantlistUpdate {
	if i != nil {
		wu.SetUserID(*i)
	}
	return wu
}

// SetUser sets the "user" edge to the User entity.
func (wu *WantlistUpdate) SetUser(u *User) *WantlistUpdate {
	return wu.SetUserID(u.ID)
}

// Mutation returns the WantlistMutation object of the builder.
func (wu *WantlistUpdate) Mutation() *WantlistMutation {
	return wu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (wu *WantlistUpdate) ClearUser() *WantlistUpdate {
	wu.mutation.ClearUser()
	return wu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WantlistUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, wu.sqlSave, wu.mutation, wu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WantlistUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WantlistUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WantlistUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wu *WantlistUpdate) check() error {
	if v, ok := wu.mutation.Title(); ok {
		if err := wantlist.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Wantlist.title": %w`, err)}
		}
	}
	if _, ok := wu.mutation.UserID(); wu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Wantlist.user"`)
	}
	return nil
}

func (wu *WantlistUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := wu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(wantlist.Table, wantlist.Columns, sqlgraph.NewFieldSpec(wantlist.FieldID, field.TypeInt))
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.Title(); ok {
		_spec.SetField(wantlist.FieldTitle, field.TypeString, value)
	}
	if wu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   wantlist.UserTable,
			Columns: []string{wantlist.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   wantlist.UserTable,
			Columns: []string{wantlist.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wantlist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	wu.mutation.done = true
	return n, nil
}

// WantlistUpdateOne is the builder for updating a single Wantlist entity.
type WantlistUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WantlistMutation
}

// SetTitle sets the "title" field.
func (wuo *WantlistUpdateOne) SetTitle(s string) *WantlistUpdateOne {
	wuo.mutation.SetTitle(s)
	return wuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (wuo *WantlistUpdateOne) SetNillableTitle(s *string) *WantlistUpdateOne {
	if s != nil {
		wuo.SetTitle(*s)
	}
	return wuo
}

// SetUserID sets the "user_id" field.
func (wuo *WantlistUpdateOne) SetUserID(i int) *WantlistUpdateOne {
	wuo.mutation.SetUserID(i)
	return wuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (wuo *WantlistUpdateOne) SetNillableUserID(i *int) *WantlistUpdateOne {
	if i != nil {
		wuo.SetUserID(*i)
	}
	return wuo
}

// SetUser sets the "user" edge to the User entity.
func (wuo *WantlistUpdateOne) SetUser(u *User) *WantlistUpdateOne {
	return wuo.SetUserID(u.ID)
}

// Mutation returns the WantlistMutation object of the builder.
func (wuo *WantlistUpdateOne) Mutation() *WantlistMutation {
	return wuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (wuo *WantlistUpdateOne) ClearUser() *WantlistUpdateOne {
	wuo.mutation.ClearUser()
	return wuo
}

// Where appends a list predicates to the WantlistUpdate builder.
func (wuo *WantlistUpdateOne) Where(ps ...predicate.Wantlist) *WantlistUpdateOne {
	wuo.mutation.Where(ps...)
	return wuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WantlistUpdateOne) Select(field string, fields ...string) *WantlistUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Wantlist entity.
func (wuo *WantlistUpdateOne) Save(ctx context.Context) (*Wantlist, error) {
	return withHooks(ctx, wuo.sqlSave, wuo.mutation, wuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WantlistUpdateOne) SaveX(ctx context.Context) *Wantlist {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WantlistUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WantlistUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuo *WantlistUpdateOne) check() error {
	if v, ok := wuo.mutation.Title(); ok {
		if err := wantlist.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Wantlist.title": %w`, err)}
		}
	}
	if _, ok := wuo.mutation.UserID(); wuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Wantlist.user"`)
	}
	return nil
}

func (wuo *WantlistUpdateOne) sqlSave(ctx context.Context) (_node *Wantlist, err error) {
	if err := wuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(wantlist.Table, wantlist.Columns, sqlgraph.NewFieldSpec(wantlist.FieldID, field.TypeInt))
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Wantlist.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, wantlist.FieldID)
		for _, f := range fields {
			if !wantlist.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != wantlist.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.Title(); ok {
		_spec.SetField(wantlist.FieldTitle, field.TypeString, value)
	}
	if wuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   wantlist.UserTable,
			Columns: []string{wantlist.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   wantlist.UserTable,
			Columns: []string{wantlist.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Wantlist{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wantlist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	wuo.mutation.done = true
	return _node, nil
}
